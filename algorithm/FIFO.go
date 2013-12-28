package algorithm

import (
	"gogo-cache/link"
)

type FIFOQueue struct {
	maxSize int64
	queue   *link.Link
}

func (q *FIFOQueue) Insert(node link.Node) *link.Node {

	for q.queue.Length >= q.maxSize {
		q.queue.Pop()
	}

	return q.queue.Unshift(node)
}

func (q *FIFOQueue) Del(node *link.Node) {
	q.queue.Del(node)
}

func (q *FIFOQueue) Update(node *link.Node) {

}

func (q *FIFOQueue) Clear() {
	q.queue.Clear()
}

func (q *FIFOQueue) Print() {
	q.queue.Print()
}

func (q *FIFOQueue) Equal(keys []string) bool {
	return q.queue.Equal(keys)
}

func New_FIFO(size int64) *FIFOQueue {
	q := link.New_Link()
	return &FIFOQueue{
		maxSize: size,
		queue:   q,
	}
}
