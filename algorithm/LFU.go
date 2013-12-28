package algorithm

import (
	"gogo-cache/link"
)

type LFUQueue struct {
	maxSize int64
	queue   *link.Link
}

func (q *LFUQueue) Insert(node link.Node) *link.Node {

	for q.queue.Length >= q.maxSize {
		q.queue.Pop()
	}

	return q.queue.Push(node)
}

func (q *LFUQueue) Del(node *link.Node) {
	q.queue.Del(node)
}

func (q *LFUQueue) Update(node *link.Node) {

	node.Count++

	preNode := node.Pre
	nextNode := node.Next
	queue := q.queue

	if preNode != nil && preNode.Count < node.Count {
		for preNode != nil && preNode.Count < node.Count { //Todo: should think about when count are same, should think about expire time
			queue.Forward(node)
			preNode = node.Pre
		}
	} else if nextNode != nil && nextNode.Count > node.Count {
		for preNode != nil && preNode.Count < node.Count {
			queue.Backward(node)
			nextNode = node.Next
		}
	}
}

func (q *LFUQueue) Clear() {
	q.queue.Clear()
}

func (q *LFUQueue) Equal(keys []string) bool {
	return q.queue.Equal(keys)
}

func (q *LFUQueue) Print() {
	q.queue.Print()
}

func New_LFU(size int64) *LFUQueue {
	q := link.New_Link()
	return &LFUQueue{
		maxSize: size,
		queue:   q,
	}
}
