package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	cache map[string]interface{}
	mu    sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (c *Cache) Insert(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *Cache) Fetch(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.Unlock()
	value, ok := c.cache[key]
	return value, ok
}

func (c *Cache) Delete(key String) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, key)
}

func main() {
	c := NewCache()
	c.Insert("key1", "value1")
	c.Insert("key2", "value2")
	if val, ok := c.Fetch("key1"); ok {
		fmt.Println(val)
	} else {
		fmt.Println("not found")
	}
	c.Delete("key2")
	if _, ok := c.Fetch("Key2"); !ok {
		fmt.Println("Key Deleted")
	} else {
		fmt.Println("Key Not Deleted")
	}
}
