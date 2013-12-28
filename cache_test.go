package cache

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_Basic_Cache(t *testing.T) {
	fmt.Println("[Test LRU LFU] get and set with 10000 item")

	var (
		index int
		size  int = 10000
	)

	cacheLFU := New("LFU", int64(size))
	cacheLRU := New("LRU", int64(size))

	for index = 0; index < size; index++ {
		key := "key_" + strconv.Itoa(index)
		value := "value_" + strconv.Itoa(index)
		cacheLFU.Set(key, value, 20*time.Minute)
		cacheLRU.Set(key, value, 20*time.Minute)

	}

	for index = 0; index < size; index++ {
		key := "key_" + strconv.Itoa(index)
		value1 := cacheLFU.Get(key)
		value2 := cacheLRU.Get(key)

		if value1 == nil && value2 == nil {
			t.Error("every key should be match in the cache, not match key is %v", key)
		}
	}

	defer cacheLFU.Clear()
	defer cacheLRU.Clear()
}
