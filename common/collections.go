package common

// ReverseMap will take a map of __K: __V comparable items and reverse it so the output is __V: __K
func ReverseMap[__K comparable, __V comparable](m map[__K]__V) map[__V]__K {
	n := make(map[__V]__K, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}
