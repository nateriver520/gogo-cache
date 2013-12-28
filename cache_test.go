package cache

import (
	"fmt"
	"testing"
	"time"
)

func Test_Set(t *testing.T) {
	cache := New_Cache("LFU", 100)
	cache.Set("word", "123", 20*time.Millisecond)
	value := cache.get("word")

	fmt.Printf("key is %v  and value is %v \n", "word", value)
}
