package common

// Contains can take in any arr where each element __T is comparable and determines if the element is in the arr
func Contains[__T comparable](arr []__T, elem __T) bool {
	for _, val := range arr {
		if val == elem {
			return true
		}
	}

	return false
}
