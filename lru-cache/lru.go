package lru

const cacheSize = 5

type LRUCache interface {
	LRUCache([]string) string
}
