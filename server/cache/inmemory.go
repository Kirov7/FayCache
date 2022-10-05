package cache

import "sync"

type inMemoryCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
	Stat
}

func (i *inMemoryCache) Set(k string, v []byte) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	tmp, exist := i.c[k]
	if exist {
		i.del(k, tmp)
	}
	i.c[k] = v
	i.add(k, v)
	return nil
}

func (i *inMemoryCache) Get(k string) ([]byte, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	return i.c[k], nil
}

func (i *inMemoryCache) Del(k string) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	v, exist := i.c[k]
	if exist {
		delete(i.c, k)
		i.del(k, v)
	}
	return nil
}

func (i *inMemoryCache) GetStat() Stat {
	return i.Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{
		c:     make(map[string][]byte),
		mutex: sync.RWMutex{},
	}
}
