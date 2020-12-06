package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	dataAsString := load("puzzle_input.txt")
	//dataAsString := load("sample_input.txt")

	fmt.Println("step1 -->", step1(dataAsString))
	fmt.Println("step1 -->", step2(dataAsString))
}

func step2(dataAsString string) int {
	sum := 0
	last := true
	firstLine := ""
	for _, line := range strings.Split(dataAsString, "\n") {
		if last {
			firstLine = line
			last = false
		}
		if len(line) == 0 {
			last = true
			sum += len(firstLine)
		}

		currentLine := ""

		for _, char := range line {
			if strings.Contains(firstLine, string(char)) {
				currentLine += string(char)
			}
		}

		firstLine = currentLine
	}
	return sum
}

func step1(dataAsString string) int {
	sum := 0
	yes := map[string]bool{}
	for _, line := range strings.Split(dataAsString, "\n") {
		if len(line) == 0 {
			sum += len(yes)
			yes = map[string]bool{}
		}
		for _, char := range line {
			yes[string(char)] = true
		}
	}
	return sum
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
