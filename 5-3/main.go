package main

import "math"

func chmin(a, b float64, i int, list []int) {
	if a > b {
		list[i] = int(b)
	}
}

func DpFrog(h []int) int {
	dp := make([]int, len(h))
	dp[0] = 0
	for i := 1; i < len(h); i++ {
		dp[i] = 1000000000
	}
	dp[0] = 0

	for i := 1; i < len(h); i++ {
		chmin(float64(dp[i]), float64(dp[i-1])+math.Abs(float64(h[i-1])-float64(h[i])), i, dp)
		if i > 1 {
			chmin(float64(dp[i]), float64(dp[i-2])+math.Abs(float64(h[i-2]-h[i])), i, dp)
		}
	}

	return dp[len(dp)-1]
}
