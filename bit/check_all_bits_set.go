package main

import (
	"fmt"
)

// If all bits in the binary representation of n are set,
// then adding '1' to it will produce a number which will
// be a perfect power of 2.
// So, check whether the new number is a perfect power of 2 or not.
//
// Reference: http://bit.ly/2Xx25FX
func areAllBitsSet(n int) bool {
	if n == 0 {
		return false
	}

	// Check whether the new number is a perfect power of 2 or not.
	if (n+1)&n == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(areAllBitsSet(0))
	fmt.Println(areAllBitsSet(1))
	fmt.Println(areAllBitsSet(2))
	fmt.Println(areAllBitsSet(3))
	fmt.Println(areAllBitsSet(4))
	fmt.Println(areAllBitsSet(5))
	fmt.Println(areAllBitsSet(6))
	fmt.Println(areAllBitsSet(7))
	fmt.Println(areAllBitsSet(8))
}
