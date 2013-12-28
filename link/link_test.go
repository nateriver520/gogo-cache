package link

import (
	"fmt"
	"testing"
)

func Test_PopAndPush(t *testing.T) {
	fmt.Println("[Test] pop and push function")

	link := New_Link()
	link.Push(New_Node("a", "av", 0))
	link.Push(New_Node("b", "bv", 0))
	link.Push(New_Node("c", "cv", 0))
	link.Push(New_Node("d", "dv", 0))

	link.Print()

	if link.Length != 4 {
		t.Error("push error")
	}

	node := link.Pop()

	if link.Length != 3 {
		t.Error("pop error")
	}

	if node.Key != "d" {
		t.Error("pop node should have a key d, but got ", node.Key)
		node.Print()
	}

	link.Print()
}

func Test_Del(t *testing.T) {
	fmt.Println("[Test] del function")

	link := New_Link()
	link.Push(New_Node("a", "av", 0))
	node1 := link.Push(New_Node("b", "bv", 0))
	link.Push(New_Node("c", "cv", 0))
	link.Del(node1)

	if link.Length != 2 {
		t.Error("del node should make link to 2, but got ", link.Length)
	}

	link.Print()
}

func Test_Unshift(t *testing.T) {
	fmt.Println("[Test] unshift function")

	link := New_Link()
	link.Push(New_Node("a", "av", 0))
	link.Push(New_Node("b", "bv", 0))
	link.Push(New_Node("c", "cv", 0))

	node := link.Unshift(New_Node("d", "dv", 0))

	if link.Length != 4 {
		t.Error("unshift should insert d to head ")
	}

	if node.Next.Key != "a" {
		t.Error("d next node's key should be a ")
	}

}

func Test_ForwardAndBackward(t *testing.T) {
	fmt.Println("[Test] Forward and Backward function")

	link := New_Link()
	node1 := link.Push(New_Node("a", "av", 0))
	node2 := link.Push(New_Node("b", "bv", 0))
	node3 := link.Push(New_Node("c", "cv", 0))
	node4 := link.Push(New_Node("d", "dv", 0))

	link.Backward(node4)

	if node4.Next != nil {
		t.Error("backward node4 should nothing happened")
	}

	link.Backward(node3)

	if node3.Next != nil {
		t.Error("node 3 should be the tail")
	}

	link.Forward(node1)

	if node1.Pre != nil {
		t.Error("node 1 should be the head, and nothing happened")
	}

	link.Forward(node2)

	if node2.Next.Key != "a" {
		t.Error("node 2 should be the head, and next is node1")
	}

	link.Print()

	link.Forward(node4)
	link.Forward(node4)
	link.Forward(node4)

	link.Print()

}

func Test_MoveHead(t *testing.T) {
	fmt.Println("[Test] move head")

	link := New_Link()
	link.Push(New_Node("a", "av", 0))
	link.Push(New_Node("b", "bv", 0))
	link.Push(New_Node("c", "cv", 0))
	node4 := link.Push(New_Node("d", "dv", 0))

	link.MoveHead(node4)
	link.Print()

	if node4.Pre != nil {
		t.Error("node 4 should be on the head ")
		node4.Print()
	}

}
