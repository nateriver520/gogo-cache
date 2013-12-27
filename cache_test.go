package cache

import (
	"testing"
	"time"
)

func Test_Set(t *testing.T) {
	cache := New_Cache("LUR", 100)
	cache.Set("word", "123", 20*time.Millisecond)
}
