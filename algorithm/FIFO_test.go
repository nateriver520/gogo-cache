package algorithm

import (
	"fmt"
	"gogo-cache/link"
	"testing"
)

func Test_FIFO(t *testing.T) {
	fmt.Println("[Test] FIFO")

	queue := New_FIFO(3)
	queue.Insert(link.New_Node("a", "av", 0))
	queue.Insert(link.New_Node("b", "bv", 0))
	queue.Insert(link.New_Node("c", "cv", 0))
	queue.Insert(link.New_Node("d", "dv", 0))

	if !queue.Equal([]string{"d", "c", "b"}) {
		t.Error("fifo should be d c b ")
	}
}
