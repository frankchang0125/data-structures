package main

import (
	"fmt"
)

// Recursion with memoization
// Time complexity: O(n)
// Space complexity: O(n)
func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}

	mem := make([]int, n+1)
	mem[0] = 1
	mem[1] = 1

	var fibMem func(i int) int
	fibMem = func(i int) int {
		if mem[i] != 0 {
			return mem[i]
		}

		mem[i] = fibMem(i-1) + fibMem(i-2)
		return mem[i]
	}

	return fibMem(n)
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(12))
}
