package main

import "fmt"

func Fibo(N int) uint64 {
	var list []uint64
	list = append(list, 0, 1)

	for i := 2; i <= N; i++ {
		list = append(list, list[i-1]+list[i-2])
	}
	fmt.Println(list)
	return list[len(list)-1]
}
