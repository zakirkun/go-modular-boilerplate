package simplecache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

type SimpleCache struct {
	Cache     *cache.Cache
	ExpiredAt int
	PurgeTime int
}

type ICache interface {
	Open() *cache.Cache
	Set(key string, data interface{})
	Get(key string) *interface{}
	Delete(key string)
}

func NewSimpleCache(s SimpleCache) ICache {

	return &SimpleCache{
		ExpiredAt: s.ExpiredAt,
		PurgeTime: s.PurgeTime,
	}
}

func (s *SimpleCache) Open() *cache.Cache {
	cacheInstance := cache.New(time.Minute*time.Duration(s.ExpiredAt), time.Minute*time.Duration(s.PurgeTime))
	s.Cache = cacheInstance

	return cacheInstance
}

func (s *SimpleCache) Set(key string, data interface{}) {
	s.Cache.Set(key, data, time.Minute*time.Duration(s.ExpiredAt))
}

func (s *SimpleCache) Get(key string) *interface{} {
	data, found := s.Cache.Get(key)

	if found {
		return &data
	}

	return nil
}

func (s *SimpleCache) Delete(key string) {
	s.Cache.Delete(key)
}
