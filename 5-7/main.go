package main

import (
	"fmt"
	"math"
)

func chmax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Knapsack(weight, value []int, W int) int {

	DPTable := make([][]int, len(weight)+1)

	for i := 0; i < len(weight)+1; i++ {
		DPTable[i] = make([]int, W+1)
	}

	for i := 0; i < len(weight); i++ {
		for w := 0; w <= W; w++ {
			if w-weight[i] >= 0 {
				DPTable[i+1][w] = chmax(DPTable[i+1][w], DPTable[i][w-weight[i]]+value[i])
			}

			DPTable[i+1][w] = chmax(DPTable[i+1][w], DPTable[i][w])
		}
	}
	fmt.Println(DPTable)
	return DPTable[len(weight)][W]
}
