package main

import "fmt"

func main() {
	fmt.Println("Start Game!")

	left, right := 20, 36

	for right-left > 1 {
		mid := left + ((right - left) / 2)

		var anser string
		fmt.Printf("Is the age less than %d ? \n", mid)
		fmt.Scan(&anser)

		if anser == "yes" {
			right = mid
			return
		} else {
			left = mid
		}
	}
}
