package cache

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_Basic_Cache(t *testing.T) {
	fmt.Println("[Test LRU LFU FIFO] get and set with 10000 item")

	var (
		index int
		size  int = 10000
	)

	cacheLFU := New("LFU", int64(size))
	cacheLRU := New("LRU", int64(size))
	cacheFIFO := New("FIFO", int64(size))

	for index = 0; index < size; index++ {
		key := "key_" + strconv.Itoa(index)
		value := "value_" + strconv.Itoa(index)
		cacheLFU.Set(key, value, 0)
		cacheLRU.Set(key, value, 0)
		cacheFIFO.Set(key, value, 0)
	}

	for index = 0; index < size; index++ {
		key := "key_" + strconv.Itoa(index)
		value1 := cacheLFU.Get(key)
		value2 := cacheLRU.Get(key)
		value3 := cacheFIFO.Get(key)

		if value1 == nil && value2 == nil && value3 == nil {
			t.Error("every key should be match in the cache, not match key is %v", key)
		}
	}

	defer cacheLFU.Clear()
	defer cacheLRU.Clear()
	defer cacheFIFO.Clear()
}

func Test_Del_Expire_Length(t *testing.T) {
	fmt.Println("[Test] del expire count")

	cache := New("LFU", 100)
	cache.Set("10s_key", "value", time.Second*10)

	if cache.Get("10s_key") == nil {
		t.Error("we should get the key")
	}

	time.Sleep(time.Second * 10)

	if cache.Get("10s_key") != nil {
		t.Error("this key should expire")
	}

	cache.Set("10s_key", "value", -1)

	if cache.Get("10s_key") == nil {
		t.Error("we should get the key")
	}

	cache.Del("10s_key")

	if cache.Get("10s_key") != nil {
		t.Error("we should del the key")
	}

	cache.Set("10s_key_1", "value", -1)
	cache.Set("10s_key", "value", -1)

	if cache.Count() != 2 {
		t.Error("cache's count should be 2")
	}

}
