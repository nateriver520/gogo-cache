package algorithm

import (
	"fmt"
	"gogo-cache/link"
	"testing"
)

func Test_LRU(t *testing.T) {
	fmt.Println("[Test] LRU")

	queue := New_LRU(3)
	queue.Insert(link.New_Node("a", "av", 0))
	queue.Insert(link.New_Node("b", "bv", 0))
	node3 := queue.Insert(link.New_Node("c", "cv", 0))

	queue.Update(node3)

	if !queue.Equal([]string{"c", "b", "a"}) {
		t.Error("node3 shoud be in the head")
	}

	queue.Insert(link.New_Node("d", "dv", 0))

	if !queue.Equal([]string{"d", "c", "b"}) {
		t.Error("node4 shoud be in the head")
	}
}
