package main

import (
	"context"
	"fmt"
	"inmemcache-example/inmemcache"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

const (
	// TimeExpired represent for the time example data will expire in mem cache
	TimeExpired int32 = 60 // seconds
)

type ExampleData struct {
	Name     string
	Value    string
	ExpireAt time.Time
}

var inMemoryCache inmemcache.InMemCache

func main() {
	inMemoryCache = inmemcache.NewInMemCache()
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) (ExampleData, error) {
	data, err := getData(inMemoryCache)
	if err != nil {
		fmt.Println("get data err: ", err.Error())
	}
	return data, nil
}

func getData(inMemoryCache inmemcache.InMemCache) (ExampleData, error) {
	var exampleData ExampleData

	cacheKey := "example-key"
	// Get example data from in memory cache
	cacheData, err := inMemoryCache.Get(cacheKey)
	if err == nil && cacheData != nil {
		exampleData = cacheData.(ExampleData)

		if time.Since(exampleData.ExpireAt).Seconds() < 0 {
			return exampleData, nil
		}

		// Delete data in mem cache if expired time exceed. This is to get latest data from DB.
		_ = inMemoryCache.Delete(cacheKey)
	}

	// If don't have example data in mem cache. Get example data from DB.
	exampleData, err = queryDataFromDB("name_test")
	if err != nil {
		return exampleData, err
	}

	// Set expired time for example data and set into mem cache
	exampleData.ExpireAt = time.Now().Add(time.Duration(TimeExpired) * time.Second)
	_ = inMemoryCache.Set(cacheKey, exampleData)

	return exampleData, nil
}

func queryDataFromDB(name string) (ExampleData, error) {
	dataDB := ExampleData{
		Name:  name,
		Value: "value_test",
	}

	// Sleep 100 millisecond for testing
	time.Sleep(50 * time.Millisecond)
	return dataDB, nil
}
