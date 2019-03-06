package main

import (
	"fmt"
)

// Dynamic programming
// Time complexity: O(n)
// Space complexity: O(n)
func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// Dynamic programming without array
// Time complexity: O(n)
// Space complexity: O(1)
func fibonacci2(n int) int {
	if n <= 1 {
		return 1
	}

	dp1 := 1
	dp2 := 1

	for i := 2; i < n+1; i++ {
		dp2, dp1 = dp1+dp2, dp2
	}

	return dp2
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(12))

	fmt.Println()

	fmt.Println(fibonacci2(0))
	fmt.Println(fibonacci2(1))
	fmt.Println(fibonacci2(4))
	fmt.Println(fibonacci2(5))
	fmt.Println(fibonacci2(12))
}
