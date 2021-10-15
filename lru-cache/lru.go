package lru

import (
	"strings"
)

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
	place    map[string]*Node
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
	lru.place[val] = &n
}

func (lru *LRU) Find(val string) *Node {
	return lru.place[val]
}

func (lru *LRU) Delete(n *Node) {
	val := n.Val
	if n.Prev != nil {
		n.Prev.Next = n.Next
	} else {
		lru.first = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	} else {
		lru.last = n.Prev
	}
	lru.lenght--
	delete(lru.place, val)
}

func (lru *LRU) Length() int {
	return lru.lenght
}

func LRUCache(calls []string) string {
	cache := LRU{capacity: cacheSize, place: make(map[string]*Node)}
	for _, v := range calls {
		if node := cache.Find(v); node != nil {
			cache.Delete(node)
		}
		if cache.lenght == cache.capacity {
			cache.Delete(cache.first)
		}
		cache.PushBack(v)
	}

	var sb strings.Builder
	sb.Grow(cache.lenght)
	for node := cache.first; node != nil; node = node.Next {
		sb.WriteString(node.Val)
		if node.Next != nil {
			sb.WriteRune('-')
		}
	}
	return sb.String()
}
