package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	validPasswordCountStep1 := 0
	validPasswordCountStep2 := 0

	dataAsString := load("puzzle_input.txt")

	for _, line := range strings.Split(dataAsString, "\n") {
		items := strings.Split(line, " ")

		bounds := strings.Split(items[0], "-")
		min := atoi(bounds[0])
		max := atoi(bounds[1])

		letter := items[1][0]
		password := items[2]

		//step1
		count := strings.Count(password, string(letter))
		if count >= min && count <= max {
			validPasswordCountStep1++
		}

		//step2
		if (password[min-1] == letter) != (password[max-1] == letter) {
			validPasswordCountStep2++
		}
	}

	fmt.Println("step1 -->", validPasswordCountStep1)
	fmt.Println("step2 -->", validPasswordCountStep2)
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
