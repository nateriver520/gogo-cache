package link

import (
	"fmt"
	"testing"
)

func Test_PopAndPush(t *testing.T) {
	link := New_Link()
	link.Push(New_Node("a", "av", 0))
	link.Push(New_Node("b", "bv", 0))
	link.Push(New_Node("c", "cv", 0))
	link.Push(New_Node("d", "dv", 0))
	node := link.Pop()
	node.Print()
	link.Print()
}

func Test_Del(t *testing.T) {
	fmt.Printf("test del function \n")
	link := New_Link()
	link.Push(New_Node("a", "av", 0))
	node1 := link.Push(New_Node("b", "bv", 0))
	link.Push(New_Node("c", "cv", 0))

	if node1.Pre == nil {
		fmt.Println("pre is nil")
	}

	if node1.Next == nil {
		fmt.Println("next is nil")
	}

	link.Del(node1)
	link.Print()
}

func Test_node(t *testing.T) {
	node := New_Node("abc", "bcd", 0)
	node.Print()
}
