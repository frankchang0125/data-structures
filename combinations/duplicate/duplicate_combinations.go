package main

import (
	"fmt"
	"sort"
)

// C(m, n), find all n-length combinations
// of given array nums with DUPLICATE numbers.
func combinations(nums []int, n int) [][]int {
	ans := [][]int{}

	if n > len(nums) {
		return ans
	}

	// We have to sort the nums array first in order to
	// prevent creating duplicate combinations.
	sort.Sort(sort.IntSlice(nums))

	// Find all n-length combinations recursively.
	duplicateCombinationDFS(&ans, nums, 0, n, []int{})

	return ans
}

// Find all subsets of DUPLICATE numbers in nums array.
func subsets(nums []int) [][]int {
	ans := [][]int{}

	// We have to sort the nums array first in order to
	// prevent creating duplicate subsets.
	sort.Sort(sort.IntSlice(nums))

	// Equals to find C(m, n), where n:
	// 	0 (Empty subset) -> m (Length of nums, subset contains all numbers in nums).
	// Find all different-length subsets recursively.
	for i := 0; i <= len(nums); i++ {
		duplicateCombinationDFS(&ans, nums, 0, i, []int{})
	}

	return ans
}

// s: Start index of nums to be added to the combination.
// n: Target length of combination.
// cur: Current combination.
func duplicateCombinationDFS(ans *[][]int, nums []int, s int, n int, cur []int) {
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
		// Skip if:
		// 1. Since nums array has been sorted,
		//    for numbers other than first number on the same level of tree,
		//    if current number is equal to previous number,
		//    it means all the sub-combinations are all already created by previous number.
		//    ex: [1, 1, 2]:
		//                  [ ]
		//               /   |   \
		//              1    1    2
		//             / \   └─────── Second 1: i > s && nums[0] == nums[1], skip.
		//            1   2
		//           /
		//          2
		if i > s && nums[i] == nums[i-1] {
			continue
		}

		// Append number to current combination.
		cur = append(cur, nums[i])

		duplicateCombinationDFS(ans, nums, i+1, n, cur)

		// We have to pop the appended number in order to prevent
		// affecting other DFS.
		cur = cur[:len(cur)-1]
	}
}

func main() {
	// Find all C(m, n) of DUPLICATE numbers.
	fmt.Println(combinations([]int{}, 0))
	fmt.Println(combinations([]int{1}, 1))
	fmt.Println(combinations([]int{1, 1}, 1))
	fmt.Println(combinations([]int{1, 2}, 1))
	fmt.Println(combinations([]int{1, 1}, 2))
	fmt.Println(combinations([]int{1, 2}, 2))
	fmt.Println(combinations([]int{1, 1, 2}, 2))
	fmt.Println(combinations([]int{1, 1, 2}, 3))
	fmt.Println(combinations([]int{3, 2, 1}, 2))
	fmt.Println(combinations([]int{3, 2, 1}, 3))
	fmt.Println(combinations([]int{1, 1, 2, 2}, 2))
	fmt.Println(combinations([]int{1, 1, 2, 2}, 3))
	fmt.Println(combinations([]int{1, 1, 2, 2}, 4))
	fmt.Println(combinations([]int{3, 5, 6, 7, 2, 4}, 4))
	fmt.Println(combinations([]int{3, 5, 6, 7, 2, 4, 9, 8}, 6))
	fmt.Println()

	// Find all subsets of DUPLICATE numbers.
	fmt.Println(subsets([]int{}))
	fmt.Println(subsets([]int{1}))
	fmt.Println(subsets([]int{1, 2}))
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{1, 1, 2}))
	fmt.Println(subsets([]int{3, 2, 1}))
	fmt.Println(subsets([]int{1, 2, 1}))
	fmt.Println(subsets([]int{1, 1, 1}))
	fmt.Println(subsets([]int{1, 1, 1, 2}))
	fmt.Println(subsets([]int{1, 1, 2, 2}))
	fmt.Println(subsets([]int{1, 1, 2, 2, 3}))
	fmt.Println(subsets([]int{3, 5, 6, 7, 2, 4}))
}
