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

func findForStep1(goal int) {
	tick := 0
	stopNumber := goal + 10
	numbers := numbers{[]byte{3, 7}, 0, 1}
	// fmt.Println(numbers.array, numbers.elf1index, numbers.elf2index)
	for len(numbers.array) < stopNumber {
		numbers.turn()
		// fmt.Println(numbers.array, numbers.elf1index, numbers.elf2index)
		tick++
	}
	fmt.Print(goal, " --> ")
	for index := goal; index < goal+10; index++ {
		fmt.Print(numbers.array[index])
	}
	fmt.Println()
}

func computeForStep2() numbers {
	fmt.Println("computing number for step2")
	numbers := numbers{[]byte{3, 7}, 0, 1}
	for len(numbers.array) < 25000000 {
		numbers.turn()
	}
	return numbers
}

func findForStep2(numbers numbers, goal string) {
	goalArray := []byte{}
	for _, char := range goal {
		goalArray = append(goalArray, byte(char-'0'))
	}
	fmt.Println("step2", goal, " --> ", bytes.Index(numbers.array, goalArray))
}

func main() {
	findForStep1(5)
	findForStep1(9)
	findForStep1(18)
	findForStep1(2018)
	findForStep1(306281)
	num := computeForStep2()
	findForStep2(num, "51589")
	findForStep2(num, "01245")
	findForStep2(num, "92510")
	findForStep2(num, "59414")
	findForStep2(num, "306281")
}
