package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Heap key points:
// 	1. For node i:
//		a. Its left child node must at: 2i
//		b. Its right child node must at: 2i + 1
//		c. Its parent node must at: i/2
//	2. Array is built from index 1, not 0 (to meet requirement 1.)
//	   (index 0 is reserved)
//  3. For Max Heap, all child nodes are smaller than their parent node
//	   For Min Heap, all child nodes are larger than their parent node
//
// Heap sort steps:
// 	1. For max heap, the root must be the largest number among other numbers
// 	2. Therefore, starts with max heap, swap root and the last node
//  3. Pretend the root node not exist, rebuild the max heap,
//     then we can get the 'second' largest number
//  4. Repeat step 2. and 3. keep removing root node from max heap,
//     we can get 'third', 'forth', 'fifth'... largest number
//  5. Until max heap contains only one node, sort is completed
//	   and the numbers are sorted from small to large since we keep swapping
//	   root (max number) with last node (end of array) during each iteration
func HeapSort(arr *[]int) {
	// Fill index 0, since heap is built from index 1
	heapArr := append([]int{0}, *arr...)

	// Build max heap
	BuildMaxHeap(heapArr)

	last := len(heapArr) - 1
	for i := 1; i < len(heapArr)-1; i++ {
		// Swap max heap's root node with last node
		swap(heapArr, 1, last)

		// Rebuild max heap, ignore the swapped root node of max heap
		// Note:
		// 	We don't need to call BuildMaxHeap() to rebuild the
		//  entire max heap since max heap was already built and we
		//  just need to adjust the root node and it affected subtrees
		MaxHeapify(heapArr[0:last], 1)

		// Substract last to ignore the swapped root node for next iteration
		last--
	}

	heapArr = heapArr[1:]
	*arr = heapArr
}

// MaxHeapify rebuilds the max heap top-down from the specified root node
// Note:
//	It won't rebuild the entire max heap but only the specified root node
//  and it affected subtrees
func MaxHeapify(arr []int, root int) {
	left := 2 * root
	right := 2*root + 1
	largest := root

	// Find the largest number among root and its left, right child node
	if left < len(arr) && arr[left] > arr[largest] {
		largest = left
	}

	if right < len(arr) && arr[right] > arr[largest] {
		largest = right
	}

	if largest != root {
		// Swap root with largest child node and rebuild
		// the max heap for the swapped child node subtree
		swap(arr, largest, root)
		MaxHeapify(arr, largest)
	}
}

// BuildMaxHeap rebuilds the entire max heap bottom-up
// from last subtree's root node to the root node
func BuildMaxHeap(arr []int) {
	for i := len(arr) / 2; i >= 1; i-- {
		MaxHeapify(arr, i)
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

// Reference: https://goo.gl/BRyBGY
func main() {
	var arr []int

	arr = []int{}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	HeapSort(&arr)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{1}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	HeapSort(&arr)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{9, 2}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	HeapSort(&arr)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{9, 2, 8}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	HeapSort(&arr)
	fmt.Print("After sort: ")
	PrintArray(arr)

	arr = []int{3, 5, 2, 9, 6, 5, 0, 1}
	fmt.Print("Before sort: ")
	PrintArray(arr)
	HeapSort(&arr)
	fmt.Print("After sort: ")
	PrintArray(arr)
}
