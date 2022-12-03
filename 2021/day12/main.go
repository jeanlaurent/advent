package main

import "fmt"

type node struct {
	name string
	child []*node
	parent []*node
}

func main() {
	fmt.Println("hello world")
	var start = node{"start", nil, nil}
	var b = node{"b", nil, nil}
	start.child = append(start.child, *b)
	b.parent = append(b.parent, *start)
}