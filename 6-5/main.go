package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	H := make([]int, N)
	S := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&H[i])
		fmt.Scan(&S[i])
	}

	fmt.Println(binarysearch(H, S))
}

func binarysearch(h, s []int) int {
	left, right := 0, 100000
	for right-left > 1 {
		mid := (left + right) / 2

		ok := true
		t := make([]int, len(h))
		for i := 0; i < len(h); i++ {
			if mid < h[i] {
				ok = false
			} else {
				t[i] = (mid - h[i]) / s[i]
			}
		}
		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})

		//N-1秒以内に全てを割ることになるため(競技開始時に1つ割るためN-1)
		for i := 0; i < len(h); i++ {
			if t[i] < i {
				ok = false
			}
		}

		if ok {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}
