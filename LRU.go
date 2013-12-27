package cache

import (
	"fmt"
	"gogo-cache/link"
)

type LRUQueue struct {
	maxSize int64
	length  int64
	queue   link.Link
}

func (q *LRUQueue) Insert(key string) {
	fmt.Printf("this is insert function\n")
	node := link.New_Node(key)
	q.queue.Push(node)
}

func (q *LRUQueue) Del(node link.Node) {
	fmt.Printf("this is del function\n")
	node.Print()
}

func (q *LRUQueue) Update(node link.Node) {
	fmt.Printf("this is update function \n")
	node.Print()
}

func New_LRU(size int64) *LRUQueue {
	q := link.New_Link()
	return &LRUQueue{
		maxSize: size,
		length:  0,
		queue:   q,
	}
}