package main

import (
	"bytes"
	"fmt"
)

type numbers struct {
	array     []byte
	elf1index int64
	elf2index int64
}

func (n *numbers) compute() {
	fmt.Println("computing numbers")
	for len(n.array) < 25000000 {
		n.turn()
	}
}

func (n *numbers) turn() {
	sum := n.array[n.elf1index] + n.array[n.elf2index]
	if sum >= 10 {
		n.array = append(n.array, 1)
		n.array = append(n.array, sum-10)
	} else {
		n.array = append(n.array, sum)
	}
	n.elf1index = (n.elf1index + int64(n.array[n.elf1index]) + 1) % int64(len(n.array))
	n.elf2index = (n.elf2index + int64(n.array[n.elf2index]) + 1) % int64(len(n.array))
}

func (n *numbers) findForStep2(goal string) {
	goalArray := []byte{}
	for _, char := range goal {
		goalArray = append(goalArray, byte(char-'0'))
	}
	fmt.Println("step2", goal, " --> ", bytes.Index(n.array, goalArray))
}

func (n *numbers) findForStep1(goal int) {
	fmt.Print(goal, " --> ")
	for index := goal; index < goal+10; index++ {
		fmt.Print(n.array[index])
	}
	fmt.Println()
}

func main() {
	numbers := numbers{[]byte{3, 7}, 0, 1}
	numbers.compute()
	numbers.findForStep1(5)
	numbers.findForStep1(9)
	numbers.findForStep1(18)
	numbers.findForStep1(2018)
	numbers.findForStep1(306281)
	numbers.findForStep2("51589")
	numbers.findForStep2("01245")
	numbers.findForStep2("92510")
	numbers.findForStep2("59414")
	numbers.findForStep2("306281")
}
