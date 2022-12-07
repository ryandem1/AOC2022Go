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

// ReverseString will take in a string a return the reversed string
func ReverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
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

// CheckAll will take in a function that returns a bool and perform it on all items in the slice. Will return True if
// all check function calls return True, else will return False
func CheckAll[__I any](items []__I, check func(__I) bool) bool {
	result := true
	for _, item := range items {
		if !check(item) {
			result = false
			break
		}
	}
	return result
}
