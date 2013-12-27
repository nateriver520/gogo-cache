package link

import (
	"testing"
)

func Test_link(t *testing.T) {
	link := New_Link()
	link.Push(Node{nil, nil, 1, "a"})
	link.Push(Node{nil, nil, 1, "b"})
	link.Push(Node{nil, nil, 1, "c"})
	link.Push(Node{nil, nil, 1, "d"})
	node := link.Pop()
	node.Print()
	link.Print()
}

func Test_node(t *testing.T) {
	node := New_Node("abc")
	node.Print()
}
