//package woofwoof provides functionality for understanding woofers.
package woofwoof

//Years converts human years to woofwoof years
func Years(n int) int {
	return n * 7
}

// YearsTwo converts human years to dog years.
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
