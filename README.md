gogo-cache
==========

## What?

gogo-cache is an in-memory key:value store/cache similar to memcached 

It can support LRU ,LRU and FIFO . In the future, we will support more :)

The architecture of this project is base on this project [Node-Simple-Cache](https://github.com/hh54188/Node-Simple-Cache) 


##How ?

###Install

```shell
go get github.com/nateriver520/gogo-cache
```

###Sample

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

## Todo List

- add more test cases
- add more cache function
- bench mark vs redis, memcached
- add funtion for regular garbage collection
- save cache to local (Maybe?) 