package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findKthSmallest(arr []int, k int) int {
	return QuickSelect(arr, 0, len(arr) - 1, k)
}

// Quick select steps:
//	1. Find a pivot which separates numbers into two parts:
//		a. Left: Numbers smaller than pivot
//		b. Right: Numbers larger than pivot
//	2. Keep repeating Step 1. until we can find a pivot
//	   which is the kth smallest number
//	   (i.e. pivot's index == k - 1)
func QuickSelect(arr []int, front int, end int, k int) int {
	if len(arr) == 0 || len(arr) < k {
		return -1
	}

	// Only one number in this interval, just return it
	if front == end {
		return arr[front]
	}

	// Find the index of pivot
	pivotIdx := partition(arr, front, end)

	// Keep repeating until we find a pivot which is the kth smallest number
	for pivotIdx != k - 1 {
		if pivotIdx > k - 1 {
			// Pivot is at the right hand side of kth smallest number,
			// search left part
			return QuickSelect(arr, front, pivotIdx - 1, k);
		} else {
			// Pivot is at the left hand side of kth smallest number,
			// search right part
			return QuickSelect(arr, pivotIdx + 1, end, k)
		}
	}

	return arr[pivotIdx]
}

func partition(arr []int, front int, end int) int {
	// Pivot can be any number in array, we just choose the last number here
	pivot := arr[end]

	// i will be the final index of pivot,
	// which separates numbers into two parts:
	// numbers smaller than pivot and numbers larger than pivot
	// i is initialized to: (front - 1) since nothing has been
	// compared yet
	i := front - 1

	// Compare all numbers (except pivot itself)
	// and swap numbers: arr[i+1] and arr[j]
	// so that all numbers which are smaller than pivot
	// will be at the left hand side of all numbers which are
	// larger than pivot
	for j := front; j < end; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}

	// All numbers except pivot have been compared
	// Now all numbers: arr[front->i] are smaller than pivot
	// and all numbers: arr[i+1->end-1] are larger than pivot
	// Swap pivot with the first number which is larger than pivot (i++)
	// so that pivot will be at the position which:
	// all numbers at the left hand side of pivot are smaller than pivot
	// and all numbers at the right hand side of pivot are larger than pivot
	i++
	swap(arr, i, end)

	// Return pivot's index
	return i
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func PrintArray(arr []int) {
	values := make([]string, len(arr))

	for i, item := range arr {
		values[i] = strconv.Itoa(item)
	}

	fmt.Printf("[%s]\n", strings.Join(values, ", "))
}

// Reference: https://goo.gl/2AFazY
func main() {
	var arr []int
	var k int
	var kth int

	arr = []int{}
	k = 1
	kth = findKthSmallest(arr, k)
	fmt.Printf("%dth smallest number: %d\n", k, kth)

	arr = []int{1}
	k = 1
	kth = findKthSmallest(arr, k)
	fmt.Printf("%dth smallest number: %d\n", k, kth)

	arr = []int{9, 2}
	k = 2
	kth = findKthSmallest(arr, k)
	fmt.Printf("%dth smallest number: %d\n", k, kth)

	arr = []int{9, 2, 8}
	k = 2
	kth = findKthSmallest(arr, k)
	fmt.Printf("%dth smallest number: %d\n", k, kth)

	arr = []int{3, 5, 2, 9, 6, 5, 0, 1}
	k = 4
	kth = findKthSmallest(arr, k)
	fmt.Printf("%dth smallest number: %d\n", k, kth)

	arr = []int{3, 5, 2, 9, 6, 5, 0, 1}
	k = 6
	kth = findKthSmallest(arr, k)
	fmt.Printf("%dth smallest number: %d\n", k, kth)

	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	k = 7
	kth = findKthSmallest(arr, k)
	fmt.Printf("%dth smallest number: %d\n", k, kth)

	arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	k = 7
	kth = findKthSmallest(arr, k)
	fmt.Printf("%dth smallest number: %d\n", k, kth)
}
