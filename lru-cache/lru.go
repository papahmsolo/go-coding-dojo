package lru

import "strings"

const cacheSize = 5

func LRUCache(calls []string) string {
	cache := make([]string, 0, cacheSize)
	for _, v := range calls {
		idx := index(v, cache)
		if idx > -1 {
			cache = append(cache[:idx], cache[idx+1:]...)
		}
		if len(cache) >= cacheSize {
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
