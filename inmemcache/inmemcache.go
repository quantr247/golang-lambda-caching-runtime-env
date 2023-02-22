package inmemcache

import "errors"

const (
	msgInMemCacheDataNil = "InMemCacheData is nil"
)

// inMemCache represents for global variables to caching data in Lambda
type inMemCache struct {
	// InMemCacheData represents for variable to save in memory cache data
	InMemCacheData map[string]interface{}
}

type InMemCache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Delete(key string) error
}

// NewInMemCache represents for create new InMemCache instances
func NewInMemCache() InMemCache {
	return &inMemCache{
		InMemCacheData: make(map[string]interface{}),
	}
}

// Get represent get value of key from in memory cache
func (h *inMemCache) Get(key string) (value interface{}, err error) {
	if h.InMemCacheData == nil {
		return nil, errors.New(msgInMemCacheDataNil)
	}
	return h.InMemCacheData[key], nil
}

// Set represent set value to key in memory cache
func (h *inMemCache) Set(key string, value interface{}) (err error) {
	if h.InMemCacheData == nil {
		return errors.New(msgInMemCacheDataNil)
	}
	h.InMemCacheData[key] = value
	return nil
}

// Delete represent delete key in memory cache
func (h *inMemCache) Delete(key string) (err error) {
	if h.InMemCacheData == nil {
		return errors.New(msgInMemCacheDataNil)
	}
	delete(h.InMemCacheData, key)
	return nil
}
