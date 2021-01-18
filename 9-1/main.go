package main

import "fmt"

const max = 100000

var st = make([]int, max)
var top = 0

func Init() {
	top = 0
}

func IsEmpty() bool {
	return top == 0
}

func IsFull() bool {
	return top == max
}

func Push(x int) {
	if IsFull() {
		return
	}

	st[top] = x
	top++
}

func Pop() int {
	if IsEmpty() {
		return -1
	}

	top--
	return st[top]
}

func main() {
	Init()

	Push(3)
	Push(5)
	Push(7)

	fmt.Println(Pop())
	fmt.Println(Pop())

	Push(9)
}
