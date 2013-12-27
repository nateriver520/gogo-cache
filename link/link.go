package link

import (
	"fmt"
)

type Node struct {
	pre   *Node
	next  *Node
	count int64
	value string
}

type Link struct {
	length int64
	head   *Node
	tail   *Node
}

func (link *Link) Push(node Node) {
	link.length++
	if link.tail == nil {
		link.head = &node
		link.tail = &node
	} else {
		link.tail.next = &node
		node.pre = link.tail
		link.tail = &node
	}
}

func (link *Link) Pop() *Node {
	link.length--
	curNote := link.tail

	if curNote.pre != nil {
		curNote.pre.next = nil
		link.tail = curNote.pre
	} else {
		link.head = link.tail
	}

	return curNote
}

func (link *Link) Print() {
	head := link.head
	fmt.Printf("link lenghth is %d \n", link.length)
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}

func (node *Node) Print() {
	fmt.Printf("node's value is %v \n", node.value)
}

func New_Link() Link {
	link := Link{0, nil, nil}
	return link
}

func New_Node(key string) Node {
	return Node{
		nil,
		nil,
		1,
		key,
	}
}
