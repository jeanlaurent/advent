package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	dataAsString := load("puzzle_input.txt")
	floor := 0
	position := -1
	for i := 0; i < len(dataAsString); i++ {
		if string(dataAsString[i]) == "(" {
			floor++
		} else if string(dataAsString[i]) == ")" {
			floor--
		}
		if floor == -1 && position == -1 {
			position = i + 1
		}
	}
	fmt.Println("step1", floor)
	fmt.Println("step2", position)
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
