package main

import (
	"fmt"
)

// lcs[i][j], length of LCS of:
//  s1[0->i-1] substring and s2[0->j-1] substring.
// Note:
//  lcs[0][j] means LCS of:
//      Empty string vs. s2[0->j-1] substring.
//      (Not choosing any characters from s1.)
//  lcs[i][0] means LCS of:
//      s1[0->i-1] substring vs. empty string.
//      (Not choosing any characters from s2.)
func longestCommonSequence(s1, s2 string) int {
	// Initialization.
	lcs := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		lcs[i] = make([]int, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 || j == 0 {
				// Boundary of dynamic programming array.
				// lcs[0][j] or lcs[i][0] means:
				//  Not choosing any characters from s1 or s2,
				//  thus the length of LCS must be zero.
				lcs[i][j] = 0
				continue
			}

			if s1[i-1] == s2[j-1] {
				// If last characters of s1 and s2 substrings are equal,
				// then the length of longest common sequence of current
				// s1 and s2 substrings must be the length of s1 and s2
				// without current last characters' longest common sequence plus 1.
				// (1 for current common character.)
				// i.e. lcs[i-1][j-1] + 1.
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				// If last characters of s1, s2 substrings are not equal,
				// then the length of LCS of s1, s2 substrings must be
				// the maximum length of:
				//      s1[0->i-2] substring vs. s2[0->j-1] substring
				//      s1[0->i-1] substring vs. s2[0->j-2] substring
				// i.e. Max(lcs[i-1][j], lcs[i][j-1].
				if lcs[i-1][j] >= lcs[i][j-1] {
					lcs[i][j] = lcs[i-1][j]
				} else {
					lcs[i][j] = lcs[i][j-1]
				}
			}
		}
	}

	return lcs[len(s1)][len(s2)]
}

// Reference: https://goo.gl/htgbbR
func main() {
	fmt.Println(longestCommonSequence("acbad", "abcadf"))
	fmt.Println()
	fmt.Println(longestCommonSequence("abcadf", "acbad"))
	fmt.Println()
	fmt.Println(longestCommonSequence("acbade", "abcadfe"))
	fmt.Println()
	fmt.Println(longestCommonSequence("abcadfe", "acbade"))
}
