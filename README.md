gogo-cache
==========

## What?

gogo-cache is a cache package, can suppot LRU and LFU algorithm. And more algorithm will on the way :)

The architecture of this project is base on this project [Node-Simple-Cache](https://github.com/hh54188/Node-Simple-Cache) 


##How ?

###Install

```shell
go get github.com/nateriver520/gogo-cache
```

###Sample

```go

cache := New("LRU", size) // currently you can choose LFU and LRU

cache.Set(key, value, 20*time.Millisecond) //set expire time

value := cache.Get(key) // get value from cache

```

