package main

import (
	"fmt"
)

// Recursion
// Time complexity: O(2^n)
// Space complexity: O(n)
func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(12))
}
