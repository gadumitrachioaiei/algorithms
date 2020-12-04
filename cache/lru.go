package cache

import "container/list"

type LRU struct {
	m    map[string]Value
	l    *list.List
	size int
}

func NewLRU(size int) *LRU {
	return &LRU{
		m:    make(map[string]Value),
		l:    list.New(),
		size: size,
	}
}

func (lru *LRU) Get(key string) string {
	if v, ok := lru.m[key]; ok {
		lru.l.MoveToFront(v.el)
		return v.v
	}
	return ""
}

func (lru *LRU) Set(key, value string) {
	if v, ok := lru.m[key]; ok {
		lru.l.MoveToFront(v.el)
		lru.m[key] = Value{el: v.el, v: value}
		return
	}
	if len(lru.m) < lru.size {
		el := lru.l.PushFront(key)
		lru.m[key] = Value{el: el, v: value}
		return
	}
	oldestKey := lru.l.Remove(lru.l.Back())
	delete(lru.m, oldestKey.(string))
	lru.m[key] = Value{el: lru.l.PushFront(key), v: value}
}

func (lru *LRU) Len() int {
	return len(lru.m)
}

type Value struct {
	el *list.Element
	v  string
}
