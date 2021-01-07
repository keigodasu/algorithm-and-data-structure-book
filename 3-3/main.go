package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	list := make([]int, N)

	for i := 0; i < len(list); i++ {
		list[i] = i
	}

	min := FindMin(list)

	fmt.Println(min)
}

func FindMin(list []int) int {
	const INF = 20000000
	min := INF

	for _, num := range list {
		if num < min {
			min = num
		}
	}

	return min
}
