package core

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type rtdb struct {
	db *cache.Cache
}

func New() *rtdb {
	db := cache.New(5*time.Minute, 10*time.Minute)
	return &rtdb{
		db,
	}
}

func (r *rtdb) Write(pid string, val string) bool {
	r.db.Set(pid, val, cache.NoExpiration)
	return true
}

func (r *rtdb) Read(pid string) string {
	val, found := r.db.Get(pid)
	if found {
		return val.(string)
	}
	return ""
}
