package main

import "fmt"

type Heap struct {
	heap []int
}

func (h *Heap) Push(x int) {
	h.heap = append(h.heap, x)

	i := len(h.heap) - 1
	for i > 0 {
		p := (i - 1) / 2
		if h.heap[p] < x {
			h.heap[i] = h.heap[p]
			i = p
		} else {
			break
		}
	}
	h.heap[i] = x
}

func (h *Heap) Top() int {
	if !h.Empty() {
		return h.heap[0]
	} else {
		return -1
	}
}

func (h *Heap) Pop() {
	if h.Empty() {
		return
	}

	x := h.heap[len(h.heap)-1]
	h.heap = h.heap[1:]
	i := 0

	for i*2+1 < len(h.heap) {
		child1 := i*2 + 1
		child2 := i*2 + 2
		if child2 < len(h.heap) && h.heap[child2] > h.heap[child1] {
			child1 = child2
		}
		if h.heap[child1] <= x {
			break
		}
		h.heap[i] = h.heap[child1]
		i = child1
	}

	h.heap[i] = x
}

func (h *Heap) Empty() bool {
	return len(h.heap) == 0
}

func main() {
	h := Heap{heap: []int{}}
	h.Push(5)
	h.Push(3)
	h.Push(7)
	h.Push(1)

	fmt.Println(h.Top())
	h.Pop()
	fmt.Println(h.Top())
	h.Push(11)
	fmt.Println(h.Top())
}
