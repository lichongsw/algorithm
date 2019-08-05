package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity         int
	doubleLinkedList *list.List
	hashMap          map[interface{}]*list.Element
}

type cacheEntry struct {
	key   interface{}
	value interface{}
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		capacity:         capacity,
		doubleLinkedList: list.New(),
		hashMap:          make(map[interface{}]*list.Element),
	}
}

// Add adds a value to the cache.
func (c *LRUCache) Add(key interface{}, value interface{}) {
	// check nil
	if c.hashMap == nil {
		c.hashMap = make(map[interface{}]*list.Element)
		c.doubleLinkedList = list.New()
	}

	// update value of exist key
	if entry, ok := c.hashMap[key]; ok {
		c.doubleLinkedList.MoveToFront(entry)
		entry.Value.(*cacheEntry).value = value
		return
	}

	// add new key
	newEntry := c.doubleLinkedList.PushFront(&cacheEntry{key, value})
	c.hashMap[key] = newEntry

	// remove oldest if reach the capacity
	if c.capacity != 0 && c.doubleLinkedList.Len() > c.capacity {
		removeEntry := c.doubleLinkedList.Back()
		c.doubleLinkedList.Remove(removeEntry)
		delete(c.hashMap, removeEntry.Value.(*cacheEntry).key)
	}
}

// Get looks up a key's value from the cache.
func (c *LRUCache) Get(key interface{}) (value interface{}, ok bool) {
	if c.hashMap == nil {
		return
	}

	if foundEntry, ok := c.hashMap[key]; ok {
		c.doubleLinkedList.MoveToFront(foundEntry)
		return foundEntry.Value.(*cacheEntry).value, true
	}
	return
}

func main() {
	fmt.Println("Welcome to the playground!")

	lruCache := New(5)
	lruCache.Add(1, "Monday")
	lruCache.Add(2, "Tuesday")
	lruCache.Add(3, "Wednesday")
	lruCache.Add(4, "Thursday")
	lruCache.Add(5, "Friday")
	lruCache.Add(6, "Satday")

	fmt.Println("Get newest element: ", lruCache.doubleLinkedList.Front().Value.(*cacheEntry).key)

	value, ok := lruCache.Get(5)
	if ok {
		fmt.Println("Get element: ", lruCache.doubleLinkedList.Front().Value.(*cacheEntry).key)
		fmt.Println("Get element from cache: ", value)
	}

	fmt.Println("Get newest element: ", lruCache.doubleLinkedList.Front().Value.(*cacheEntry).key)
}
