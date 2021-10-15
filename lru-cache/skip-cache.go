package lru

import "strings"

type SkipLRU struct{}

func (SkipLRU) LRUCache(calls []string) string {
	reversedResult := make([]string, 0, cacheSize)

	// from the end of calls, if not in result - add
	for i := len(calls) - 1; i >= 0; i-- {
		call := calls[i]
		if !contains(reversedResult, call) {
			reversedResult = append(reversedResult, call)
		}

		// got full result - breaking
		if len(reversedResult) == cap(reversedResult) {
			break
		}
	}

	// reversing
	result := make([]string, 0, len(reversedResult))
	for i := len(reversedResult) - 1; i >= 0; i-- {
		result = append(result, reversedResult[i])
	}

	return strings.Join(result, "-")
}

func contains(arr []string, elem string) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}

	return false
}
