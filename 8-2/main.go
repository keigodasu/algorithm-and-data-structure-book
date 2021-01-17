package main

import (
	"fmt"
)

type Node struct {
	Next *Node
	Name string
}

var sentinel = &Node{
	Next: nil,
	Name: "",
}

func PrintList() {
	printString := ""
	cur := sentinel.Next
	for ; cur != sentinel; cur = cur.Next {
		printString += " " + cur.Name
	}
	fmt.Println(printString)
}

func Insert(v *Node) {
	v.Next = sentinel.Next
	sentinel.Next = v
}

func main() {
	sentinel.Next = sentinel

	list := []string{"yamamot", "watanabe", "ito", "takahashi", "suzuki", "sato"}

	for i, name := range list {
		node := Node{
			Next: nil,
			Name: name,
		}
		Insert(&node)
		fmt.Printf("step %d: ", i)
		PrintList()
	}

}
