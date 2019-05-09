package main

import (
	"fmt"
)

// If substracting '1' from a number, then its binary representation:
// 	1. Right-most bit '1' will become '0'.
//	2. If there're any bits '0's at its right, they will all become '1'.
//	3. The left bits of the right-most bit '1' will remain the same.
// 		ex: 1100 - 1 = 1011.
// Then, if we 'AND' it with the original number, we will make all
// right bits of the right-most bit '1' become 0.
//
// Thus, we can keep repeating the above operations and check whether
// the number has become 0 or not to count the number of set bits
// of the number.
// 	ex: n = 1100 - 1 = 1011, 1011 & 1100 = 1000
//		n = 1000 - 1 = 0111, 0111 & 1000 = 0000 (n = 0, terminate)
//		Number of set bits of 1100: 2.
func numOfSetBits(n int) int {
	count := 0

	for n != 0 {
		count++
		n = (n - 1) & n
	}

	return count
}

func main() {
	fmt.Println(numOfSetBits(0))
	fmt.Println(numOfSetBits(1))
	fmt.Println(numOfSetBits(2))
	fmt.Println(numOfSetBits(3))
	fmt.Println(numOfSetBits(4))
	fmt.Println(numOfSetBits(5))
	fmt.Println(numOfSetBits(6))
	fmt.Println(numOfSetBits(7))
	fmt.Println(numOfSetBits(8))
	fmt.Println(numOfSetBits(10))
	fmt.Println(numOfSetBits(11))
	fmt.Println(numOfSetBits(12))
	fmt.Println(numOfSetBits(13))
	fmt.Println(numOfSetBits(14))
	fmt.Println(numOfSetBits(15))
	fmt.Println(numOfSetBits(16))
}
