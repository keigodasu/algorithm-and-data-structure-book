package main

import (
	"fmt"
	"math"
)

func chmin(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func SeparateCost(N int) int {

	c := make([][]int, N+1)

	for i := 0; i <= N; i++ {
		c[i] = make([]int, N+1)
	}

	for i := 0; i <= N; i++ {
		for j := 0; j < N; j++ {
			c[i][j] = i + j
		}
	}

	dp := make([]int, N+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 10000000000
	}

	dp[0] = 0

	for i := 0; i <= N; i++ {
		for j := 0; j < i; j++ {
			dp[i] = chmin(dp[i], dp[j]+c[i][j])
		}
	}

	fmt.Println(dp)
	return dp[N]
}
