package lru

import "container/list"

type Cache struct{
	maxBytes  int64
	nBytes    int64
	ll 		  *list.List
	cache     map[string]*list.Element
	//某条记录被删除时的回调函数 可以为nil
	OnEvicted func(key string , value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64 , onEvicted func(string , Value)) *Cache{
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Len() int{
	return c.ll.Len()
}

func (c *Cache) Get(key string) (value Value , ok bool){
	if element , ok := c.cache[key];ok{
		c.ll.MoveToFront(element)
		kv := element.Value.(*entry)
		return kv.value , true
	}
	return
}

func (c *Cache) RemoveOldest() {
	//要删除的元素
	element := c.ll.Back()
	if element != nil{
		c.ll.Remove(element)
		kv := element.Value.(*entry)
		delete(c.cache , kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil{
			c.OnEvicted(kv.key , kv.value)
		}
	}
}

func (c *Cache) Add(key string , value Value){
	if element , ok := c.cache[key];ok{
		c.ll.MoveToFront(element)
		kv := element.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	}else {
		element := c.ll.PushFront(&entry{key , value})
		c.cache[key] = element
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nBytes{
		c.RemoveOldest()
	}
}