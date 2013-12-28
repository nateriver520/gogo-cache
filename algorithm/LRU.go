package algorithm

import (
	"gogo-cache/link"
)

type LRUQueue struct {
	maxSize int64
	queue   *link.Link
}

func (q *LRUQueue) Insert(node link.Node) *link.Node {

	for q.queue.Length >= q.maxSize {
		q.queue.Pop()
	}

	return q.queue.Unshift(node)
}

func (q *LRUQueue) Del(node *link.Node) {
	q.queue.Del(node)
}

func (q *LRUQueue) Update(node *link.Node) {
	q.queue.MoveHead(node)
}

func (q *LRUQueue) Clear() {
	q.queue.Clear()
}

func (q *LRUQueue) Print() {
	q.queue.Print()
}

func New_LRU(size int64) *LRUQueue {
	q := link.New_Link()
	return &LRUQueue{
		maxSize: size,
		queue:   q,
	}
}
