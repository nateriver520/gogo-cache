gogo-cache
==========

## What?

gogo-cache is an in-memory key:value store/cache similar to memcached 

It can support LRU and LRU. In the future, we will support more :)

The architecture of this project is base on this project [Node-Simple-Cache](https://github.com/hh54188/Node-Simple-Cache) 


##How ?

###Install

```shell
go get github.com/nateriver520/gogo-cache
```

###Sample

```go

import (
  "github.com/nateriver520/gogo-cache"
)

cache := New("LRU", size) // currently you can choose LFU and LRU

cache.Set(key, value, 20*time.Millisecond) //set expire time

value := cache.Get(key) // get value from cache

```

