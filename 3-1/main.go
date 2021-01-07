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

	if findNumber(list, v){
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func findNumber(list []int, v int) (existFlag bool){
	for i := 0; i < len(list); i++ {
		if list[i] == v {
			existFlag = true
		}
	}
	return existFlag
}