package main

import "fmt"

func main() {
	var (
		N int
		v int
	)

	fmt.Scan(&N)
	fmt.Scan(&v)

	list := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&list[i])
	}

	if FindIndex(list, 1) != -1 {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}

}

func FindIndex(list []int, number int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == number {
			return i
		}
	}
	return -1
}
