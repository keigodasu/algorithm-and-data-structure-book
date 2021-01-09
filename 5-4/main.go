package main

import "math"

func chmin(a, b float64, i int, list []int) {
	if a > b {
		list[i] = int(b)
	}
}

func DpPushBased(h []int) int {
	INF := 1000000000
	dp := make([]int, len(h))
	dp[0] = 0
	for i := 1; i < len(h); i++ {
		dp[i] = INF
	}

	for i := 0; i < len(h); i++ {
		if i+1 < len(h) {
			chmin(float64(dp[i+1]), float64(dp[i])+math.Abs(float64(h[i])-float64(h[i+1])), i+1, dp)
		}

		if i+2 < len(h) {
			chmin(float64(dp[i+2]), float64(dp[i])+math.Abs(float64(h[i]-h[i+2])), i+2, dp)
		}
	}

	return dp[len(dp)-1]
}
