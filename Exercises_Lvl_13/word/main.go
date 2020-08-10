//Package word blah blah
package word

import "strings"

// UseCount creates a map of words used in a string, counting how many each is used
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of different words in a string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
