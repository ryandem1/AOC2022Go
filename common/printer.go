package common

import (
	"encoding/json"
	"fmt"
)

// PrettyPrintMap will take in an input map and pretty print it to the stdout. Helpful for debugging
func PrettyPrintMap[__K comparable, __V any](input map[__K]__V) {
	output, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	fmt.Println()
}
