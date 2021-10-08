package lru

import "strings"

const cacheSize = 5

type Node struct {
	Val  string
	Next *Node
	Prev *Node
}

type LRU struct {
	first    *Node
	last     *Node
	lenght   int
	capacity int
}

func (lru *LRU) PushBack(val string) {
	n := Node{Val: val}

	if lru.lenght > 0 {
		n.Prev = lru.last
		lru.last.Next = &n
	}

	if lru.lenght == 0 {
		lru.first = &n
	}

	lru.last = &n
	lru.lenght++

	if lru.lenght > lru.capacity {
		second := lru.first.Next
		second.Prev = nil
		lru.first = second
		lru.lenght--
	}
}

func (lru *LRU) Find(val string) *Node {
	for cur := lru.first; cur != nil; cur = cur.Next {
		if cur.Val == val {
			return cur
		}
	}
	return nil
}

func (lru *LRU) Delete(n *Node) {
	if n.Prev != nil {
		n.Prev.Next = n.Next
		lru.first = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
		lru.last = n.Prev
	}
	lru.lenght--
}

func (lru *LRU) Length() int {
	return lru.lenght
}

func LRUCache(calls []string) string {
	cache := make([]string, 0, cacheSize)
	for _, v := range calls {
		idx := index(v, cache)
		if idx > -1 {
			cache = append(cache[:idx], cache[idx+1:]...)
		}
		if len(cache) == cacheSize {
			cache = cache[1:]
		}
		cache = append(cache, v)
	}
	return strings.Join(cache, "-")
}

func index(el string, s []string) int {
	for i, v := range s {
		if v == el {
			return i
		}
	}
	return -1
}
