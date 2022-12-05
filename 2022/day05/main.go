package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type stack []string

func (s *stack) Push(a string) {
	*s = append(*s, a)
}

func (s *stack) Peek() string {
	return (*s)[len(*s)-1]
}

func (s *stack) Pop() string {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func (s stack) List() {
	for _, line := range s {
		fmt.Print("[", line, "]")
	}
	fmt.Println()
}

func main() {
	step1("small.txt")
	step1("input.txt")
	step2("small.txt")
	step2("input.txt")
}

func step1(filename string) {

	stacks, moveLines := initStacks(filename)

	for _, moveLine := range moveLines {
		newMoves := strings.Split(moveLine, " ")
		times := atoi(newMoves[1])
		from := atoi(newMoves[3]) - 1
		to := atoi(newMoves[5]) - 1
		for i := 0; i < times; i++ {
			stacks[to].Push(stacks[from].Pop())
		}
	}

	result := ""
	for i := 0; i < len(stacks); i++ {
		result += stacks[i].Peek()
	}
	fmt.Println("step1 (", filename, ")-->", result)
}

func step2(filename string) {

	stacks, moveLines := initStacks(filename)

	for _, moveLine := range moveLines {
		newMoves := strings.Split(moveLine, " ")
		times := atoi(newMoves[1])
		from := atoi(newMoves[3]) - 1
		to := atoi(newMoves[5]) - 1

		pivotIndex := len(stacks[from]) - times
		slice := stacks[from][pivotIndex:]
		stacks[from] = stacks[from][:pivotIndex]

		for i := 0; i < len(slice); i++ {
			stacks[to].Push(slice[i])
		}
	}

	result := ""
	for i := 0; i < len(stacks); i++ {
		result += stacks[i].Peek()
	}

	fmt.Println("step2 (", filename, ")-->", result)
}

func initStacks(filename string) ([]stack, []string) {
	foundEmptyLine := false
	stackLines := []string{}
	moveLines := []string{}
	for _, line := range strings.Split(load(filename), "\n") {
		if line == "" {
			foundEmptyLine = true
			continue
		}
		if foundEmptyLine {
			moveLines = append(moveLines, line)
		} else {
			stackLines = append(stackLines, line)
		}
	}

	// really, no way to easily reverse an array in go... come on...
	for i, j := 0, len(stackLines)-1; i < j; i, j = i+1, j-1 {
		stackLines[i], stackLines[j] = stackLines[j], stackLines[i]
	}

	firstLine := stackLines[0]
	indexes := []int{}
	for i := 0; i < len(firstLine); i++ {
		if string(firstLine[i]) != " " {
			indexes = append(indexes, i)
		}
	}

	maxStack := len(indexes)
	stacks := []stack{}
	for i := 0; i < maxStack; i++ {
		stacks = append(stacks, stack{})
	}

	stackLines = stackLines[1:]

	for j := 0; j < len(stackLines); j++ {
		line := stackLines[j]
		for i := 0; i < len(indexes); i++ {
			if string(line[indexes[i]]) != " " {
				stacks[i].Push(string(line[indexes[i]]))
			}
		}
	}
	return stacks, moveLines
}

func printStack(stacks []stack) {
	for i := 0; i < len(stacks); i++ {
		fmt.Println("stack", i)
		stacks[i].List()
	}
	fmt.Println()
}

func load(filename string) string {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}

func atoi(line string) int {
	number, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return number
}
