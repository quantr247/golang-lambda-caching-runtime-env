package memcache

import "errors"

const (
	msgMemCacheDataNil = "MemCacheData is nil"
)

// memCache represents for global variables to caching data in Lambda
type memCache struct {
	// MemCacheData represents for variable to save mem cache data
	MemCacheData map[string]interface{}
}

type MemCache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Delete(key string) error
}

// NewMemCache represents for create new MemCache instances
func NewMemCache() MemCache {
	return &memCache{
		MemCacheData: make(map[string]interface{}),
	}
}

func (h *memCache) Get(key string) (value interface{}, err error) {
	if h.MemCacheData == nil {
		return nil, errors.New(msgMemCacheDataNil)
	}
	return h.MemCacheData[key], nil
}

func (h *memCache) Set(key string, value interface{}) (err error) {
	if h.MemCacheData == nil {
		return errors.New(msgMemCacheDataNil)
	}
	h.MemCacheData[key] = value
	return nil
}

func (h *memCache) Delete(key string) (err error) {
	if h.MemCacheData == nil {
		return errors.New(msgMemCacheDataNil)
	}
	delete(h.MemCacheData, key)
	return nil
}
