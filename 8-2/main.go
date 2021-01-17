package main

import (
	"fmt"
)

type Node struct {
	Next *Node
	Prev *Node
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
	sentinel.Next.Prev = v
	sentinel.Next = v
	v.Prev = sentinel
}

func Erase(v *Node) {
	if v == sentinel {
		return
	}
	v.Prev.Next = v.Next
	v.Next.Prev = v.Prev
}

func main() {
	sentinel.Next = sentinel
	var deleteTargetNode *Node

	list := []string{"yamamot", "watanabe", "ito", "takahashi", "suzuki", "sato"}

	for i, name := range list {
		node := Node{
			Next: nil,
			Name: name,
		}
		if node.Name == "watanabe" {
			deleteTargetNode = &node
		}
		Insert(&node)
		fmt.Printf("step %d: ", i)
		PrintList()
	}
	PrintList()
	Erase(deleteTargetNode)
	PrintList()
}
