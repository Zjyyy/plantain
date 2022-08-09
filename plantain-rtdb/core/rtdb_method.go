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

func (m *rtdbMethod) Write(pid string, val string) bool {
	m.db.Set(pid, val, cache.NoExpiration)
	return true
}

func (m *rtdbMethod) Read(pid string) string {
	val, found := m.db.Get(pid)
	if found {
		return val.(string)
	}
	return ""
}
