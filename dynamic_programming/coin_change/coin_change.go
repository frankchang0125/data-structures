package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	// dp[i]: The minimum changes needed for money: i.
	dp := make([]int, amount+1)
	dp[0] = 0

	// Initialization.
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for m := 1; m <= amount; m++ {
		// Reset minChanges.
		minChanges := amount + 1

		// For money: m
		// dp[m] = Minimum changes needed for money: m - coins[x] + 1.
		// i.e. dp[m] = min(m - coins[x]) + 1, where x = 0 -> len(coins)-1.
		for _, c := range coins {
			change := m - c

			if change < 0 {
				// Cannot make change, continue.
				continue
			}

			// Update minChanges.
			if dp[change] < minChanges {
				minChanges = dp[change]
			}
		}

		dp[m] = minChanges + 1
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

// References:
//  https://goo.gl/r66oR3
//  https://goo.gl/6LiwoH
//  https://goo.gl/H8nftU
func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{1, 2, 5, 6}, 11))
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1, 123123124}, 5))
	fmt.Println(coinChange([]int{2, 123123124}, 5))
	fmt.Println(coinChange([]int{2, 3, 123123124}, 5))
	fmt.Println(coinChange([]int{}, 3))
	fmt.Println(coinChange([]int{1, 2, 3, 4}, 0))
}
