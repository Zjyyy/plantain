package core

import (
	"github.com/patrickmn/go-cache"
)

type rtdbMethod struct {
	db *cache.Cache
}

func NewRtdbMethod(db *cache.Cache) *rtdbMethod {
	return &rtdbMethod{
		db,
	}
}

func (m *rtdbMethod) Write(pid string, value interface{}) bool {
	m.db.Set(pid, value, cache.NoExpiration)
	return true
}

func (m *rtdbMethod) Read(pid string) interface{} {
	val, found := m.db.Get(pid)
	if found {
		return val
	}
	return ""
}
