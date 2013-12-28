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

type Link struct {
	Length int64
	head   *Node
	tail   *Node
}

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

func (link *Link) Del(node *Node) {
	preNode := node.Pre
	nextNode := node.Next
	link.Length--

	if preNode != nil && nextNode == nil {
		fmt.Println("tail node")
		link.tail = preNode
		preNode = node.Next

	} else if preNode != nil && nextNode != nil {
		fmt.Println("middle node")
		preNode.Next = nextNode
		nextNode.Pre = preNode

	} else if preNode == nil && nextNode != nil {
		fmt.Println("first node")
		link.head = nextNode
		nextNode.Pre = nil

	} else if preNode == nil && nextNode == nil {
		fmt.Println("only one node")
		link.head = nil
		link.tail = nil
	}

}

func (link *Link) clear() {
	link.tail = nil
	link.head = nil
}

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

// Returns true if the item has expired.
func (node *Node) Expired() bool {
	if node.Expire == nil {
		return false
	}
	return node.Expire.Before(time.Now())
}

func (link *Link) Print() {
	head := link.head
	fmt.Printf("link lenghth is %d \n", link.Length)
	for head != nil {
		head.Print()
		head = head.Next
	}
}

func (node *Node) Print() {
	fmt.Printf("node's key is %v value is %v \n", node.Key, node.Value)
}

func New_Link() Link {
	link := Link{0, nil, nil}
	return link
}

func New_Node(key string, value interface{}, expire time.Duration) Node {

	var e *time.Time

	if expire > 0 {
		t := time.Now().Add(expire)
		e = &t
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
