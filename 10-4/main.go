package main

import "fmt"

type Edge struct {
	to int
	w  int
}

func main() {
	N, M := 3, 3

	graph := make([][]Edge, N)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)

		graph[a] = append(graph[a], Edge{
			to: a,
			w:  b,
		})
	}

	fmt.Println(graph)
}
