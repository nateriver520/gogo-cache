package cache

import (
	"gogo-cache/algorithm"
	"gogo-cache/link"
	"sync"
	"time"
)

type Queue interface {
	Insert(node link.Node) *link.Node
	Update(node *link.Node)
	Del(node *link.Node)
	Clear()
}

type Cache struct {
	sync.Mutex
	maxCount int64
	nodeMap  map[string]*link.Node
	queue    Queue
}

//insert an item to the cache, replacing any existing item
//If the expire <= 0, the item will never expires
func (cache *Cache) Set(key string, value interface{}, expire time.Duration) {
	cache.Lock()
	defer cache.Unlock()

	nodeMap := cache.nodeMap
	queue := cache.queue
	newNode := link.New_Node(key, value, expire)

	node, ok := nodeMap[key]

	if ok {
		node.Value = newNode.Value
		node.Expire = newNode.Expire
	} else {
		inserNode := queue.Insert(newNode)
		nodeMap[key] = inserNode
	}
}

//Get an item from the cache. Returns the item or nil
func (cache *Cache) Get(key string) interface{} {
	cache.Lock()
	defer cache.Unlock()

	nodeMap := cache.nodeMap
	queue := cache.queue

	node, ok := nodeMap[key]

	if !ok {
		return nil
	}

	if node.Expired() {
		queue.Del(node)
		delete(nodeMap, key)
		return nil
	}

	queue.Update(node)
	return node.Value

}

// Delete an item from the cache. Does nothing if the key is not in the cache.
func (cache *Cache) Del(key string) {
	cache.Lock()
	defer cache.Unlock()

	nodeMap := cache.nodeMap
	queue := cache.queue

	node, ok := nodeMap[key]

	if !ok {
		return
	}

	queue.Del(node)
	delete(nodeMap, key)

}

//delete all items from cache
func (cache *Cache) Clear() {
	cache.Lock()
	defer cache.Unlock()

	cache.queue.Clear()
	cache.nodeMap = map[string]*link.Node{}
}

//Returns the number of items in the cache.
//This may include items that have expired, but have not yet been cleaned up
func (cache *Cache) Count() int64 {
	cache.Lock()
	defer cache.Unlock()

	return int64(len(cache.nodeMap))
}

func New(algName string, maxCount int64) *Cache {
	var q Queue
	LURQueue := algorithm.New_LRU(maxCount)
	LFUQueue := algorithm.New_LFU(maxCount)
	FIFOQueue := algorithm.New_LFU(maxCount)

	switch algName {
	case "LUR":
		q = LURQueue
	case "LFU":
		q = LFUQueue
	case "FIFO":
		q = FIFOQueue
	default:
		q = LFUQueue
	}

	return &Cache{
		maxCount: maxCount,
		nodeMap:  make(map[string]*link.Node),
		queue:    q,
	}
}
