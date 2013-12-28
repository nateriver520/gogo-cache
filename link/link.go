package link

import (
	"fmt"
	"time"
)

type Node struct {
	Pre    *Node
	Next   *Node
	Count  int64
	Key    string
	Value  interface{}
	Expire *time.Time
}

//link will be a two-way link like this a <-> b <-> c
type Link struct {
	Length int64
	head   *Node
	tail   *Node
}

//insert a node to link head
func (link *Link) Unshift(node Node) *Node {
	if link.tail == nil {
		link.head = &node
		link.tail = &node
	} else {
		link.head.Pre = &node
		node.Next = link.head
		link.head = &node
	}

	link.Length++
	return &node
}

//del node from link
func (link *Link) Del(node *Node) {
	preNode := node.Pre
	nextNode := node.Next
	link.Length--

	if preNode != nil && nextNode == nil {
		//tail node
		link.tail = preNode
		preNode = node.Next

	} else if preNode != nil && nextNode != nil {
		//middle node
		preNode.Next = nextNode
		nextNode.Pre = preNode

	} else if preNode == nil && nextNode != nil {
		//first node
		link.head = nextNode
		nextNode.Pre = nil

	} else if preNode == nil && nextNode == nil {
		//only one node
		link.head = nil
		link.tail = nil
	}

	node = nil
}

//clear link
func (link *Link) Clear() {
	node := link.head

	for node != nil {
		node = node.Next
		node = nil
	}

	link.Length = 0
	link.tail = nil
	link.head = nil
}

//push a node
func (link *Link) Push(node Node) *Node {
	link.Length++
	if link.tail == nil {
		link.head = &node
		link.tail = &node
	} else {
		link.tail.Next = &node
		node.Pre = link.tail
		link.tail = &node
	}

	return &node
}

// pop a node
func (link *Link) Pop() *Node {
	link.Length--
	curNote := link.tail

	if curNote.Pre != nil {
		curNote.Pre.Next = nil
		link.tail = curNote.Pre
	} else {
		link.head = link.tail
	}

	return curNote
}

// link like this a -> b -> c -> d
// Forward(a) should nothing happened
// Forward(b) should like b -> a -> c -> d
func (link *Link) Forward(node *Node) {
	preNode := node.Pre
	nextNode := node.Next

	if preNode == nil {
		return
	}

	if preNode.Pre == nil {
		if nextNode == nil {
			node.Pre = nil
			node.Next = preNode

			preNode.Pre = node
			preNode.Next = nil

			link.tail = preNode
			link.head = node
		} else {
			nextNode.Pre = preNode

			preNode.Next = nextNode
			preNode.Pre = node

			node.Next = preNode
			node.Pre = nil

			link.head = node
		}
	} else {
		// should be tail
		prepreNode := preNode.Pre

		if nextNode == nil {
			prepreNode.Next = node

			node.Next = preNode
			node.Pre = prepreNode

			preNode.Pre = node
			preNode.Next = nil

			link.tail = preNode
		} else {
			prepreNode.Next = node

			node.Next = preNode
			node.Pre = prepreNode

			preNode.Pre = node
			preNode.Next = nextNode

			nextNode.Pre = preNode
		}

	}
}

// link like this a -> b -> c -> d
// Backward(d) should nothing happened
// Backward(b) should like a -> c -> b -> d
func (link *Link) Backward(node *Node) {
	preNode := node.Pre
	nextNode := node.Next

	if nextNode == nil {
		return
	}

	if nextNode.Next == nil {
		if preNode == nil {
			nextNode.Next = node
			nextNode.Pre = nil

			node.Pre = nextNode
			node.Next = nil

			link.head = nextNode
			link.tail = node
		} else {
			preNode.Next = nextNode

			nextNode.Next = node
			nextNode.Pre = preNode

			node.Next = nil
			node.Pre = nextNode

			link.tail = node
		}
	} else {
		nextnextNode := nextNode.Next

		if preNode == nil {
			nextNode.Next = node
			nextNode.Pre = nil

			node.Next = nextnextNode
			node.Pre = nextNode

			nextnextNode.Pre = node

			link.head = nextNode

		} else {
			preNode.Next = nextNode

			nextNode.Next = node
			nextNode.Pre = preNode

			node.Next = nextnextNode
			node.Pre = nextNode

			nextnextNode.Pre = node

		}
	}
}

// move a node to the head
// a -> b -> c, MoveHead(b)
// c -> a - > b
func (link *Link) MoveHead(node *Node) {
	preNode := node.Pre
	nextNode := node.Next

	// if node in the tail
	if preNode != nil && nextNode == nil {

		link.tail = preNode
		preNode.Next = node.Next

		node.Next = link.head
		link.head.Pre = node

		link.head = node
		node.Pre = nil

	} else if preNode != nil && nextNode != nil {
		//node in the middle
		preNode.Next = nextNode
		nextNode.Pre = preNode

		node.Next = link.head
		node.Pre = nil
		link.head.Pre = node

		link.head = node

	}
}

// print link, use for debug
func (link *Link) Print() {
	head := link.head

	fmt.Printf("link's head key is %v and tail key is %v link lenghth is %d \n", link.head.Key, link.tail.Key, link.Length)
	for head != nil {
		head.Print()
		head = head.Next
	}
}

// equal keys use for debug
// input keys like ["a","b","c"]
func (link *Link) Equal(keys []string) bool {
	head, index := link.head, 0

	for head != nil {
		if index >= len(keys) || keys[index] != head.Key {
			return false
		}

		index++
		head = head.Next
	}

	if index != len(keys) {
		return false
	}
	return true
}

// Returns true if the item has expired.
func (node *Node) Expired() bool {
	if node.Expire == nil {
		return false
	}
	return node.Expire.Before(time.Now())
}

func (node *Node) Print() {
	fmt.Printf("node's key is %v value is %v and count is %v ", node.Key, node.Value, node.Count)

	var (
		pre  string
		next string
	)

	if node.Pre == nil {
		pre = "none"
	} else {
		pre = node.Pre.Key
	}

	if node.Next == nil {
		next = "none"
	} else {
		next = node.Next.Key
	}

	fmt.Printf("node's pre node key  is %v and next key is %v \n", pre, next)

}

func New_Link() *Link {
	link := Link{0, nil, nil}
	return &link
}

func New_Node(key string, value interface{}, expire time.Duration) Node {

	var e *time.Time

	if expire > 0 {
		t := time.Now().Add(expire)
		e = &t
	} else {
		e = nil
	}

	return Node{
		Pre:    nil,
		Next:   nil,
		Count:  1,
		Key:    key,
		Value:  value,
		Expire: e,
	}
}
