package core

import (
	"github.com/patrickmn/go-cache"
)

type rtdbMethod struct {
	cache *cache.Cache
}

func NewRtdbMethod(cache *cache.Cache) *rtdbMethod {
	return &rtdbMethod{
		cache,
	}
}

func (m *rtdbMethod) Write(pid string, value interface{}) bool {
	m.cache.Set(pid, value, cache.NoExpiration)
	return true
}

func (m *rtdbMethod) Read(pid string) interface{} {
	val, found := m.cache.Get(pid)
	if found {
		return val
	}
	return ""
}
