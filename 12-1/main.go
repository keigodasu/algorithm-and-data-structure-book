package main

import "fmt"

func InsertionSort(list []int) {
	N := len(list)

	for i := 0; i < N; i++ {
		v := list[i]

		j := i
		for ; j > 0; j-- {
			if list[j-1] > v {
				list[j] = list[j-1]
			} else {
				break
			}
		}
		list[j] = v
	}
	fmt.Println(list)
}

func main() {
	InsertionSort([]int{3, 1, 5, 4})
}
