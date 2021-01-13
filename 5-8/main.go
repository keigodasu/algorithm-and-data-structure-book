package main

import (
	"fmt"
	"math"
)

func chmin(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func Diff(S, T string) int {
	INF := 2 << 32

	dp := make([][]int, len(S)+1)

	for i := 0; i <= len(S); i++ {
		dp[i] = make([]int, len(T)+1)
	}

	for i := 0; i < len(S)+1; i++ {
		for j := 0; j <= len(T); j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0

	for i := 0; i <= len(S); i++ {
		for j := 0; j <= len(T); j++ {
			if i > 0 && j > 0 {
				if S[i-1] == T[i-1] {
					dp[i][j] = chmin(dp[i][j], dp[i-1][j-1])
					fmt.Println("S[i-1] == T[i-1]")
					fmt.Println(dp)
				} else {
					dp[i][j] = chmin(dp[i][j], dp[i-1][j-1]+1)
					fmt.Println("S[i-1] !== T[i-1]")
					fmt.Println(dp)
				}
			}

			if i > 0 {
				dp[i][j] = chmin(dp[i][j], dp[i-1][j]+1)
				fmt.Println("dp[i-1][j]")
				fmt.Println(dp)
			}

			if j > 0 {
				dp[i][j] = chmin(dp[i][j], dp[i][j-1]+1)
				fmt.Println("dp[i][j-1]")
				fmt.Println(dp)
			}
		}
	}

	return dp[len(S)][len(T)]
}
