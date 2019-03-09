package main

import (
	"fmt"
)

// C(m, n), find all n-length combinations
// of given array nums with DISTINCT numbers.
func combinations(nums []int, n int) [][]int {
	ans := [][]int{}

	if n > len(nums) {
		return ans
	}

	// Find all n-length combinations recursively.
	distinctCombinationDFS(&ans, nums, 0, n, []int{})

	return ans
}

// Find all subsets of DISTINCT numbers in nums array.
func subsets(nums []int) [][]int {
	ans := [][]int{}

	// Equals to find C(m, n), where n:
	// 	0 (Empty subset) -> m (Length of nums, subset contains all numbers in nums).
	// Find all different-length subsets recursively.
	for i := 0; i <= len(nums); i++ {
		distinctCombinationDFS(&ans, nums, 0, i, []int{})
	}

	return ans
}

// s: Start index of nums to be added to the combination.
// n: Target length of combination.
// cur: Current combination.
func distinctCombinationDFS(ans *[][]int, nums []int, s int, n int, cur []int) {
	if len(cur) == n {
		// cur may be changed by other DFS,
		// create a copy and add it to the answers.
		c := make([]int, len(cur))
		copy(c, cur)
		*ans = append(*ans, c)
		return
	}

	// Add number to the combination and recursively call DFS
	// to find all combinations.
	for i := s; i < len(nums); i++ {
		// Append number to current combination.
		cur = append(cur, nums[i])

		distinctCombinationDFS(ans, nums, i+1, n, cur)

		// We have to pop the appended number in order to prevent
		// affecting other DFS.
		cur = cur[:len(cur)-1]
	}
}

func main() {
	// Find all C(m, n) of DISTINCT numbers.
	fmt.Println(combinations([]int{}, 0))
	fmt.Println(combinations([]int{1}, 1))
	fmt.Println(combinations([]int{1, 2}, 1))
	fmt.Println(combinations([]int{1, 2}, 2))
	fmt.Println(combinations([]int{1, 2, 3}, 2))
	fmt.Println(combinations([]int{3, 2, 1}, 2))
	fmt.Println(combinations([]int{3, 2, 1}, 3))
	fmt.Println(combinations([]int{3, 5, 6, 7, 2, 4}, 4))
	fmt.Println(combinations([]int{3, 5, 6, 7, 2, 4, 9, 8}, 6))
	fmt.Println()

	// Find all subsets of DISTINCT numbers.
	fmt.Println(subsets([]int{}))
	fmt.Println(subsets([]int{1}))
	fmt.Println(subsets([]int{1, 2}))
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{3, 2, 1}))
	fmt.Println(subsets([]int{3, 5, 6, 7, 2, 4}))
}
