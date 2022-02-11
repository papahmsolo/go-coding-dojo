package lru

import "testing"

func TestSkipLRU_LRUCache(t *testing.T) {
	testLRU(t, SkipLRU{})
}
