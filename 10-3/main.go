package main

import "fmt"

func main() {
	M := 3

	graph := make([][]int, M)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)

		graph[i] = []int{a, b}
	}

	fmt.Println(graph)
}
