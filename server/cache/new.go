package cache

import "log"

type STORAGE_TYPE int8

const (
	STORAGE_INMEMORY STORAGE_TYPE = 0
	STORAGE_BOLTDB   STORAGE_TYPE = 1
)

func New(typ STORAGE_TYPE) Cache {
	var c Cache
	switch typ {
	case STORAGE_INMEMORY:
		c = newInMemoryCache()
	case STORAGE_BOLTDB:
		c = newBoltCache()
	}
	if c == nil {
		panic("unknown cache type ")
	}
	log.Println(typ, "ready to serve")
	return c
}
