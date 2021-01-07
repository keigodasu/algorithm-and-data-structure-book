package main

import "fmt"

func main() {
	var (
		N, v int
	)
	fmt.Scan(&N)
	fmt.Scan(&v)

	list := make([]int, N)

	for i := 0; i<N; i++ {
		fmt.Scan(&list[i])
	}

	existFlag := false

	for i := 0; i < N; i++ {
		if (list[i] == v) {
			existFlag = true
		}
	}

	if existFlag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}