package main

import "fmt"

var memo []int

func Fibo(N int) int {
	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	if len(memo) > N {
		fmt.Println("Hit in memo")
		return memo[N]
	}
	memo = append(memo, Fibo(N-1)+Fibo(N-2))

	return memo[len(memo)-1]
}
