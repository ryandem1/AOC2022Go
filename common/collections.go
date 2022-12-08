package common

import (
	"sort"
)

// ReverseMap will take a map of K: V comparable items and reverse it so the output is V: K
func ReverseMap[K comparable, V comparable](m map[K]V) map[V]K {
	n := make(map[V]K, len(m))
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
func SortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

// FindOneObj is for looking for a specific obj in a slice of objs by a query function passed in. Will perform the query
// function on every element in the slice until the query returns true, then will return a pointer to the item and exit.
// If there is no object found matching the query, this will return nil
func FindOneObj[T any](sl []T, query func(T) bool) (obj T) {
	for _, elem := range sl {
		if query(elem) {
			obj = elem
			break
		}
	}
	return obj
}
