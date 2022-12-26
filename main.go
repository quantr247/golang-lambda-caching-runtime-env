package main

import (
	"context"
	"fmt"
	"memcache-example/memcache"
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

var memCache memcache.MemCache

func main() {
	memCache = memcache.NewMemCache()
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) (ExampleData, error) {
	data, err := getData(memCache)
	if err != nil {
		fmt.Println("get data err: ", err.Error())
	}
	return data, nil
}

func getData(memCache memcache.MemCache) (ExampleData, error) {
	var exampleData ExampleData

	memCacheKey := "example-key"
	// Get example data from cache
	cacheData, err := memCache.Get(memCacheKey)
	if err == nil && cacheData != nil {
		exampleData = cacheData.(ExampleData)

		if time.Since(exampleData.ExpireAt).Seconds() < 0 {
			return exampleData, nil
		}

		// Delete data in mem cache if expired time exceed. This is to get latest data from DB.
		_ = memCache.Delete(memCacheKey)
	}

	// If don't have example data in mem cache. Get example data from DB.
	exampleData, err = queryDataFromDB("name_test")
	if err != nil {
		return exampleData, err
	}

	// Set expired time for example data and set into mem cache
	exampleData.ExpireAt = time.Now().Add(time.Duration(TimeExpired) * time.Second)
	_ = memCache.Set(memCacheKey, exampleData)

	return exampleData, nil
}

func queryDataFromDB(name string) (ExampleData, error) {
	dataDB := ExampleData{
		Name:  name,
		Value: "value_test",
	}

	// Sleep 100 millisecond for testing
	time.Sleep(100 * time.Millisecond)
	return dataDB, nil
}
