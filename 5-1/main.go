package main

import "math"

func dp(h []int) int {
	var dpList []float64
	dpList = append(dpList, 0)

	for i := 1; i < len(h); i++ {
		if i == 1 {
			dpList = append(dpList, math.Abs(float64(h[i-1]-h[i])))
		} else {
			dpList = append(dpList, math.Min(dpList[i-1]+math.Abs(float64(h[i-1]-h[i])), dpList[i-2]+math.Abs(float64(h[i-2]-h[i]))))
		}
	}

	return int(dpList[len(dpList)-1])
}
