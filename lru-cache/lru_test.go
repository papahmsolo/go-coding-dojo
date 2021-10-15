package lru_test

import (
	"lru"
	"testing"
)

func TestLRU(t *testing.T) {
	cases := []struct {
		accessSequence []string
		expectedOrder  string
	}{
		{
			accessSequence: []string{},
			expectedOrder:  "",
		},
		{
			accessSequence: []string{"A"},
			expectedOrder:  "A",
		},
		{
			accessSequence: []string{"A", "A", "A"},
			expectedOrder:  "A",
		},
		{
			accessSequence: []string{"A", "B", "C"},
			expectedOrder:  "A-B-C",
		},
		{
			accessSequence: []string{"A", "B", "C", "D", "E"},
			expectedOrder:  "A-B-C-D-E",
		},
		{
			accessSequence: []string{"A", "B", "C", "D", "E", "F"},
			expectedOrder:  "B-C-D-E-F",
		},
		{
			accessSequence: []string{"A", "B", "A", "C", "A", "B"},
			expectedOrder:  "C-A-B",
		},
		{
			accessSequence: []string{"A", "B", "A", "B"},
			expectedOrder:  "A-B",
		},
		{
			accessSequence: []string{"A", "B", "C", "D", "E", "D", "Q", "Z", "C"},
			expectedOrder:  "E-D-Q-Z-C",
		},
	}

	for _, tt := range cases {
		order := lru.LRUCache(tt.accessSequence)
		if order != tt.expectedOrder {
			t.Errorf("expected %q, but got %q", tt.expectedOrder, order)
		}
	}
}
