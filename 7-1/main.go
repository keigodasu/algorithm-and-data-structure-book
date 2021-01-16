package main

var value = []int{500, 100, 50, 10, 5, 1}

func Greedy(list []int, X int) int {
	result := 0

	for i := 0; i < 6; i++ {
		add := X / value[i]
		if add > list[i] {
			add = list[i]
		}

		X -= value[i] * add
		result += add
	}

	return result
}
