package main

import (
	"fmt"
)

// best[i][j], best value of:
//  At most objects[0->i] can be picked for j-capacity knapsack.
//  i: At most objects[0->i] can be picked.
//  j: Knapsack capacity.
func knapsack(values, weights []int, capacity int) int {
	// Initialization.
	best := make([][]int, len(values))
	for i := 0; i < len(values); i++ {
		best[i] = make([]int, capacity+1)
	}

	for i := 0; i < len(values); i++ {
		for j := 0; j <= capacity; j++ {
			// Boundary of dynamic programming array.
			if j == 0 {
				// 0-capacity knapsack, nothing can be picked,
				// best value must be zero.
				best[i][j] = 0
				continue
			} else if i == 0 {
				// Only objects[0] can be picked,
				// thus if j-capacity is larger or equal than objects[0]'s weight,
				// values must be objects[0]'s value.
				// Otherwise, value must be zero, since objects[0] cannot be picked.
				if j >= weights[i] {
					best[i][j] = values[i]
				} else {
					best[i][j] = 0
				}
				continue
			}

			if j < weights[i] {
				// If j-capacity cannot hold objects[i] weight,
				// objects[i] must cannot be hold, thus,
				// the base value must be same as:
				// At most object[0->i-1], j-capacity knapsack's value.
				// i.e. best[i-1][j].
				best[i][j] = best[i-1][j]
			} else {
				// If j-capacity can hold objects[i] weight,
				// than best value must be the maximum value of:
				//  1. Try to put objects[i] into knapsack.
				//     i.e. best[i-1][j-weights[i]] + values[i],
				//          which is the best values of picking objects[0->i-1]
				//          but also save room for objects[i]
				//  2. Don't put objects[i] into knapsack.
				//     i.e. best[i-1][j].
				if best[i-1][j-weights[i]]+values[i] > best[i-1][j] {
					best[i][j] = best[i-1][j-weights[i]] + values[i]
				} else {
					best[i][j] = best[i-1][j]
				}
			}
		}
	}

	return best[len(values)-1][capacity]
}

// Reference: https://goo.gl/Cpzv5f
func main() {
    fmt.Println(knapsack([]int{3, 4, 5}, []int{2, 3, 4}, 5))
    fmt.Println()
	fmt.Println(knapsack([]int{4500, 5700, 2250, 1100, 6700},
		[]int{4, 5, 2, 1, 6}, 8))
}
