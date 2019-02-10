package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Merge sort steps:
//  1. Divide numbers into left and right part until
//     there's only one item left in each sub-array
//     If the array to be sorted is empty, simply return
//  2. Merge (sort) left part and right part
func MergeSort(arr []int, front, end int) {
	if end <= front {
		// Array to be sorted contains only one item (end == front)
		// or is empty (end < front)
		return
	}

	mid := (front + end) / 2

	// Divide
	MergeSort(arr, front, mid)
	MergeSort(arr, mid+1, end)

	// Conquer
	Merge(arr, front, mid, end)
}

// Merge steps (sorting):
// 	1. Creates two sub arrays with infinite number at the end:
//      Left array: arr[front:mid+1, Infinite] (including mid)
//  	Right array: arr[mid+1:end+1, Infinite] (including end)
//  2. Compare left and right array items one-by-one and put the smaller one
//     into original array
//     Since last items of left and right array are infinite numbers,
//     all numbers to be sorted will be compared and put into
//     original array correctly
func Merge(arr []int, front, mid, end int) {
	// arr[front:mid+1] => mid+1-front
	leftArr := make([]int, mid+1-front)
	// arr[mid+1:end+1] => (end+1) - (mid+1) = end-mid
	rightArr := make([]int, end-mid)

	// Copy array items
	copy(leftArr, arr[front:mid+1])
	copy(rightArr, arr[mid+1:end+1])

	// Copy infinite number to the end of arrays
	leftArr = append(leftArr, math.MaxInt32)
	rightArr = append(rightArr, math.MaxInt32)

	i := 0
	j := 0

	for idx := front; idx <= end; idx++ {
		if leftArr[i] <= rightArr[j] {
			arr[idx] = leftArr[i]
			i++
		} else {
			arr[idx] = rightArr[j]
			j++
		}
	}
}

func PrintArray(arr []int) {
	values := make([]string, len(arr))

	for i, item := range arr {
		values[i] = strconv.Itoa(item)
	}

	fmt.Printf("[%s]\n", strings.Join(values, ", "))
}

// Reference: https://goo.gl/dE4tQ3
func main() {
	var arr []int

	arr = []int{}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	MergeSort(arr, 0, len(arr)-1)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{1}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	MergeSort(arr, 0, len(arr)-1)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{9, 2}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	MergeSort(arr, 0, len(arr)-1)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{9, 2, 8}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	MergeSort(arr, 0, len(arr)-1)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{3, 5, 2, 9, 6, 5, 0, 1}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	MergeSort(arr, 0, len(arr)-1)
	fmt.Print("After sort: ")
	PrintArray(arr)
}
