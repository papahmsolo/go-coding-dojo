package lru

import "testing"

func TestSliceLRU_LRUCache(t *testing.T) {
	testLRU(t, SliceLRU{})
}
