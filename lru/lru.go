package lru

import "container/list"

type Cache struct {
	MaxBytes  int64
	usedByte  int64
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len()
}

func (c *Cache) New(m int, evicted func(key string, value Value)) *Cache {
	return &Cache{
		MaxBytes:  m,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: evicted,
	}
}

func (c *Cache) Add(key string, value Value) {}

func (c *Cache) RemoveOld() {}

func (c *Cache) Get(key string) {}
