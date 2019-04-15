package main

import (
	"fmt"
)

// lowerBound returns the first index: i, such that nums[i] >= target.
// Note: For binary search, nums must be sorted.
//  (For following algorithm, nums must be sorted from small to large.)
func lowerBound(nums []int, target int) int {
	// Left included / Right not included.
	l := 0
	r := len(nums)

	for l < r {
		m := l + (r-l)/2

		if nums[m] >= target {
			// Searched number is larger or equals to target,
			// search Left part.
			//  Since we want to find the lower bound of the target,
			//  thus, search Left part if searched number is larger
			//  or equals to target.
			r = m
		} else {
			// Searched number is smaller than target,
			// search Right part.
			l = m + 1
		}
	}

	// 'l' is the first index, such that nums[l] >= target.
	return l
}

// upperBound returns the first index: i, such that nums[i] > target.
// Note: For binary search, nums must be sorted.
//  (For following algorithm, nums must be sorted from small to large.)
func upperBound(nums []int, target int) int {
	// Left included / Right not included.
	l := 0
	r := len(nums)

	for l < r {
		m := l + (r-l)/2

		if nums[m] > target {
			// Searched number is larger than target,
			// search Left part.
			r = m
		} else {
			// Searched number is smaller or equals to target,
			// search Right part.
			//  Since we want to search the upper bound of the target,
			//  thus, search Right part if searched number is smaller
			//  or equals to target.
			l = m + 1
		}
	}

	// 'l' is the first index, such that nums[l] > target.
	return l
}

// References:
//	http://bit.ly/2GhVloe
//	http://bit.ly/2VdtEqk
//	http://bit.ly/2Ixh5jF
//	http://bit.ly/2IwVojI
//	http://bit.ly/2Xzn7Uv
//	http://bit.ly/2XtNOcS
func main() {
	fmt.Println(lowerBound([]int{1, 2, 2, 2, 4, 4, 5}, 0))
	fmt.Println(upperBound([]int{1, 2, 2, 2, 4, 4, 5}, 0))
	fmt.Println()

	fmt.Println(lowerBound([]int{1, 2, 2, 2, 4, 4, 5}, 1))
	fmt.Println(upperBound([]int{1, 2, 2, 2, 4, 4, 5}, 1))
	fmt.Println()

	fmt.Println(lowerBound([]int{1, 2, 2, 2, 4, 4, 5}, 2))
	fmt.Println(upperBound([]int{1, 2, 2, 2, 4, 4, 5}, 2))
	fmt.Println()

	fmt.Println(lowerBound([]int{1, 2, 2, 2, 4, 4, 5}, 3))
	fmt.Println(upperBound([]int{1, 2, 2, 2, 4, 4, 5}, 3))
	fmt.Println()

	fmt.Println(lowerBound([]int{1, 2, 2, 2, 4, 4, 5}, 4))
	fmt.Println(upperBound([]int{1, 2, 2, 2, 4, 4, 5}, 4))
	fmt.Println()

	fmt.Println(lowerBound([]int{1, 2, 2, 2, 4, 4, 5}, 5))
	fmt.Println(upperBound([]int{1, 2, 2, 2, 4, 4, 5}, 5))
	fmt.Println()

	fmt.Println(lowerBound([]int{1, 2, 2, 2, 4, 4, 5}, 6))
	fmt.Println(upperBound([]int{1, 2, 2, 2, 4, 4, 5}, 6))

}
