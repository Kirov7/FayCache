package cache

import "log"

type STORAGE_TYPE int8

const (
	STORAGE_INMEMORY STORAGE_TYPE = 0
)

func New(typ STORAGE_TYPE) Cache {
	var c Cache
	if typ == STORAGE_INMEMORY {
		c = newInMemoryCache()
	}
	if c == nil {
		panic("unknown cache type ")
	}
	log.Println(typ, "ready to serve")
	return c
}
