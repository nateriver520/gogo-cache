package algorithm

import (
	"fmt"
	"gogo-cache/link"
	"testing"
)

func Test_LFU(t *testing.T) {
	fmt.Println("[Test] LFU")

	queue := New_LFU(3)
	queue.Insert(link.New_Node("a", "av", 0))
	node2 := queue.Insert(link.New_Node("b", "bv", 0))
	node3 := queue.Insert(link.New_Node("c", "cv", 0))

	queue.Update(node3)

	if !queue.Equal([]string{"c", "a", "b"}) {
		t.Error("node3 shoud be in the head")
	}

	queue.Update(node2)
	queue.Insert(link.New_Node("d", "dv", 0))

	if !queue.Equal([]string{"c", "b", "d"}) {
		t.Error("node 1 should be delete and node4 should in the tail")
	}
}
