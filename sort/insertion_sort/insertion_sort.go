package main

import (
	"fmt"
	"strconv"
	"strings"
)

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		// Compare all the items before: arr[i],
		// swap each item with arr[i] one by one from right to left
		// until the right place for arr[i] is found
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				swap(arr, j, j+1) // arr[j] now becomes arr[i] after swapping
				continue
			}

			// Found the right place, break the loop
			break
		}
	}
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

// Reference: https://goo.gl/KUfQUz
func main() {
	var arr []int

	arr = []int{}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	InsertionSort(arr)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{1}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	InsertionSort(arr)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{9, 2}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	InsertionSort(arr)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{9, 2, 8}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	InsertionSort(arr)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{3, 5, 2, 9, 6, 5, 0, 1}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	InsertionSort(arr)
	fmt.Print("After sort: ")
	PrintArray(arr)
}
