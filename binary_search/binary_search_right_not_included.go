package main

import (
	"fmt"
)

// Note: For binary search, nums must be sorted.
//  (For following algorithm, nums must be sorted from small to big.)
func binarySearch(nums []int, target int) int {
	// Left included / Right not included.
	l := 0
	r := len(nums)

	for l < r {
		m := l + (r-l)/2

		if nums[m] == target {
			// Searched number is equals to target,
			// return the index.
			return m
		} else if nums[m] > target {
			// Searched number is greater than target,
			// search Left part.
			r = m
		} else {
			// Searched number is lesser than target,
			// search Right part.
			l = m + 1
		}
	}

	// Target not found.
	return -1
}

// floor(x): The largest number less than or equal to x,
//			 return -1 if we cannot find floor number of x.
// Note: nums must be sorted.
//  (For following algorithm, nums must be sorted from small to big.)
func floor(nums []int, target int) int {
	// Left included / Right not included.
	l := 0
	r := len(nums)
	f := -1

	for l < r {
		m := l + (r-l)/2

		if nums[m] == target {
			return nums[m]
		} else if nums[m] > target {
			// Searched number is greater than target,
			// search Left part.
			r = m
		} else {
			// Searched number is lesser than target,
			// 	Since we are looking for the floor number of target x,
			// 	assign floor to current number.
			// Search Right part.
			f = nums[m]
			l = m + 1
		}
	}

	return f
}

// ceil(x): The smallest number larger than or equal to x.
//			 return -1 if we cannot find ceil number of x.
// Note: nums must be sorted.
//  (For following algorithm, nums must be sorted from small to large.)
func ceil(nums []int, target int) int {
	// Left included / Right not included.
	l := 0
	r := len(nums)
	c := -1

	for l < r {
		m := l + (r-l)/2

		if nums[m] == target {
			return nums[m]
		} else if nums[m] > target {
			// Searched number is greater than target,
			// 	Since we are looking for the ceil number of target x,
			// 	assign ceil to current number.
			// search Left part.
			c = nums[m]
			r = m
		} else {
			// Searched number is lesser than target,
			// 	Since we are looking for the floor number of target x,
			// 	assign floor to current number.
			// Search Right part.
			l = m + 1
		}
	}

	return c
}

// References:
//	http://bit.ly/2GhVloe
//	http://bit.ly/2VdtEqk
//	http://bit.ly/2Ixh5jF
//	http://bit.ly/2IwVojI
//	http://bit.ly/2Xzn7Uv
//	http://bit.ly/2XtNOcS
func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 10))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, -10))
	fmt.Println()

	fmt.Println(binarySearch([]int{2, 5, 7, 15, 18, 25}, 2))
	fmt.Println(binarySearch([]int{2, 5, 7, 15, 18, 25}, 5))
	fmt.Println(binarySearch([]int{2, 5, 7, 15, 18, 25}, 7))
	fmt.Println(binarySearch([]int{2, 5, 7, 15, 18, 25}, 15))
	fmt.Println(binarySearch([]int{2, 5, 7, 15, 18, 25}, 18))
	fmt.Println(binarySearch([]int{2, 5, 7, 15, 18, 25}, 25))
	fmt.Println(binarySearch([]int{2, 5, 7, 15, 18, 25}, -30))
	fmt.Println(binarySearch([]int{2, 5, 7, 15, 18, 25}, 30))
	fmt.Println()

	fmt.Println(floor([]int{1, 2, 3, 4, 5}, 0))
	fmt.Println(floor([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(floor([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(floor([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(floor([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(floor([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(floor([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println()

	fmt.Println(floor([]int{1, 3, 5, 7, 9}, 0))
	fmt.Println(floor([]int{1, 3, 5, 7, 9}, 2))
	fmt.Println(floor([]int{1, 3, 5, 7, 9}, 4))
	fmt.Println(floor([]int{1, 3, 5, 7, 9}, 6))
	fmt.Println(floor([]int{1, 3, 5, 7, 9}, 8))
	fmt.Println(floor([]int{1, 3, 5, 7, 9}, 10))
	fmt.Println()

	fmt.Println(ceil([]int{1, 2, 3, 4, 5}, 0))
	fmt.Println(ceil([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(ceil([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(ceil([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(ceil([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(ceil([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(ceil([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println()

	fmt.Println(ceil([]int{1, 3, 5, 7, 9}, 0))
	fmt.Println(ceil([]int{1, 3, 5, 7, 9}, 2))
	fmt.Println(ceil([]int{1, 3, 5, 7, 9}, 4))
	fmt.Println(ceil([]int{1, 3, 5, 7, 9}, 6))
	fmt.Println(ceil([]int{1, 3, 5, 7, 9}, 8))
	fmt.Println(ceil([]int{1, 3, 5, 7, 9}, 10))
}
