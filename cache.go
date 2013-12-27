package cache

import (
	//"gogo-cache/cache"
	"gogo-cache/link"
	"time"
)

type Item struct {
	key    string
	value  interface{}
	expire time.Duration
}

type Queue interface {
	Insert(key string)
	Update(node link.Node)
	Del(node link.Node)
}

type Cache struct {
	maxCount int64
	itemMap  map[string]*Item
	queue    Queue
}

func (cache *Cache) Set(key string, value interface{}, expire time.Duration) {
	itemMap := cache.itemMap
	queue := cache.queue

	_, ok := itemMap[key]

	if ok {

		itemMap[key] = &Item{key, value, expire}

	} else {
		queue.Insert(key)
		itemMap[key] = &Item{key, value, expire}
	}

}

func New_Cache(algName string, maxCount int64) *Cache {
	var q Queue
	LURQueue := New_LRU(maxCount)

	switch algName {
	case "LUR":
		q = LURQueue
	default:
		q = LURQueue
	}

	return &Cache{
		maxCount: maxCount,
		itemMap:  make(map[string]*Item),
		queue:    q,
	}
}
