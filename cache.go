package cache

import (
	"gogo-cache/algorithm"
	"gogo-cache/link"
	"time"
)

type Queue interface {
	Insert(node link.Node) *link.Node
	Update(node *link.Node)
	Del(node *link.Node)
}

type Cache struct {
	maxCount int64
	nodeMap  map[string]*link.Node
	queue    Queue
}

func (cache *Cache) Set(key string, value interface{}, expire time.Duration) {
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

func (cache *Cache) get(key string) interface{} {
	nodeMap := cache.nodeMap
	queue := cache.queue

	node, ok := nodeMap[key]

	if !ok {
		return nil
	}

	if node.Expired() {
		queue.Del(node)
		return nil
	}

	queue.Update(node)
	return node.Value

}

func New_Cache(algName string, maxCount int64) *Cache {
	var q Queue
	//LURQueue := algorithm.New_LRU(maxCount)
	LFUQueue := algorithm.New_LFU(maxCount)

	switch algName {
	case "LUR":
		q = LFUQueue
	case "LFU":
		q = LFUQueue
	default:
		q = LFUQueue
	}

	return &Cache{
		maxCount: maxCount,
		nodeMap:  make(map[string]*link.Node),
		queue:    q,
	}
}
