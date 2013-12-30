gogo-cache
==========

## What?

gogo-cache is an in-memory key:value store/cache similar to memcached 

It supports LRU ,LRU and FIFO. 

In the future, I will support more :)

The architecture of this project is based on the project [Node-Simple-Cache](https://github.com/hh54188/Node-Simple-Cache) 


##How?

###Install

```shell
go get github.com/nateriver520/gogo-cache
```

###Example

```go
package main

import (
  "github.com/nateriver520/gogo-cache"
  "time"
)

func main() {
  cache := cache.New("LRU", 1000) // currently you can choose LFU ,LRU and FIFO, and here we set the size of item 1000

  cache.Set(key, value, 20*time.Second) //set expire time, here is 20s

  value := cache.Get(key) // get value from cache

  defer cache.clear() // clear cache

}
```

## APIs

- **Set**(key string, value interface{}, expire time.Duration)
  - insert an item to the cache, replacing any existing item
  - If the expire <= 0, the item will never expires
- **Get**(key string) interface{}
   - Get an item from the cache. Returns the item or nil
- **Del**(key string) 
  - Delete an item from the cache. Does nothing if the key is not in the cache.
- **Clear**()
  - delete all items from cache
- **Count**() int64
  - Returns the number of items in the cache
  - This may include items that have expired, but have not yet been cleaned up



## Todo List

- add more test cases
- add more cache function
- add funtion for regular garbage collection
- save cache to local (Maybe?) 