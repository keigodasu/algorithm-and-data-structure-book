package main

const max = 100000

var qu []int
var tail, head = 0, 0

func Init() {
	head, tail = 0, 0
}

func IsEmpty() bool {
	return tail == head
}

func IsFull() bool {
	return head == (tail + 1%max)
}

func Enqueu(x int) {
	if IsFull() {
		return
	}

	qu[tail] = x
	tail++
	if tail == max {
		tail = 0
	}
}

func Dequeu() int {
	if IsEmpty() {
		return -1
	}
	res := qu[head]
	head++
	if head == max {
		head = 0
	}
	return res
}
