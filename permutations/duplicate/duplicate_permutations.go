package main

import (
	"fmt"
	"sort"
)

// For permutation, [1, 2] and [2, 1] are different.
// Find all n-length permutations
// of given array nums with DUPLICATE numbers.
func permutations(nums []int, n int) [][]int {
	ans := [][]int{}

	if n > len(nums) {
		return ans
	}

	used := make([]bool, len(nums))
	sort.Sort(sort.IntSlice(nums))

	// Find all n-length permutations recursively.
	duplicatePermutationDFS(&ans, nums, used, n, []int{})

	return ans
}

// Find all permutate subsets of DUPLICATE numbers in nums array.
func subsets(nums []int) [][]int {
	ans := [][]int{}

	used := make([]bool, len(nums))
	sort.Sort(sort.IntSlice(nums))

	// Equals to find n-length permutations, where n:
	// 	0 (Empty subset) -> m (Length of nums, subset contains all numbers in nums).
	// Find all different-length permutate subsets recursively.
	for i := 0; i <= len(nums); i++ {
		duplicatePermutationDFS(&ans, nums, used, i, []int{})
	}

	return ans
}

// used: The array which indicate whether the number has been used or not
//       in current permutation.
// n: Target length of permutation.
// cur: Current permutation.
func duplicatePermutationDFS(ans *[][]int, nums []int, used []bool, n int, cur []int) {
	if len(cur) == n {
		// cur may be changed by other DFS,
		// create a copy and add it to the answers.
		c := make([]int, len(cur))
		copy(c, cur)
		*ans = append(*ans, c)
		return
	}

	// For permutation, since [1, 2] and [2, 1] are different,
	// we have to iterate from first number to last number in nums array
	// to find all permutations.
	for i := 0; i < len(nums); i++ {
		// Skip if:
		// 1. The number is already used in current permutation.
		// 2. Since nums array has been sorted,
		//    if previous number was not used but current number is equal to previous number,
		//    it means all the sub-permutations are all already created by previous number.
		//    ex: [1, 1, 2]:
		//                  [ ]
		//               /   |   \
		//              1    1    2
		//             / \   └─────── Second 1: used[0] = false && nums[0] == nums[1], skip.
		//            1   2
		//           /     \
		//          2       1
		if used[i] || (i > 0 && !used[i-1] && nums[i] == nums[i-1]) {
			continue
		}

		// Mark current number as used.
		used[i] = true
		cur = append(cur, nums[i])
		duplicatePermutationDFS(ans, nums, used, n, cur)

		// We have to pop the appended number in order to prevent
		// affecting other DFS.
		cur = cur[:len(cur)-1]

		// We have to reset current number as non-used in order to prevent
		// affecting other DFS.
		used[i] = false
	}
}

func main() {
	// Find all C(m, n) of DUPLICATE numbers.
	fmt.Println(permutations([]int{}, 0))
	fmt.Println(permutations([]int{1}, 1))
	fmt.Println(permutations([]int{1, 1}, 1))
	fmt.Println(permutations([]int{1, 2}, 1))
	fmt.Println(permutations([]int{1, 1}, 2))
	fmt.Println(permutations([]int{1, 2}, 2))
	fmt.Println(permutations([]int{1, 1, 2}, 2))
	fmt.Println(permutations([]int{1, 1, 2}, 3))
	fmt.Println(permutations([]int{1, 1, 1}, 2))
	fmt.Println(permutations([]int{1, 2, 1}, 2))
	fmt.Println(permutations([]int{1, 1, 1}, 3))
	fmt.Println(permutations([]int{1, 1, 2, 3}, 3))
	fmt.Println(permutations([]int{3, 2, 1}, 2))
	fmt.Println(permutations([]int{3, 2, 1}, 3))
	fmt.Println(permutations([]int{1, 1, 2, 2}, 2))
	fmt.Println(permutations([]int{1, 1, 2, 2}, 3))
	fmt.Println(permutations([]int{1, 1, 2, 2}, 4))
	fmt.Println(permutations([]int{3, 5, 6, 7, 2, 4}, 4))
	fmt.Println(permutations([]int{3, 5, 6, 7, 2, 4, 9, 8}, 6))
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
