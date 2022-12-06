package common

import (
	"sort"
)

// ReverseMap will take a map of __K: __V comparable items and reverse it so the output is __V: __K
func ReverseMap[__K comparable, __V comparable](m map[__K]__V) map[__V]__K {
	n := make(map[__V]__K, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

// SortedKeys will sort a map of string keys by keys and return the keys
func SortedKeys[__V any](m map[string]__V) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
