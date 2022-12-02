package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	step1("input.txt")
	step2("input.txt")
}

func step1(filename string) {
	// elem := [3]string{"rock", "paper", "scissor"}
	dataAsString := load(filename)
	score := 0
	for _, line := range strings.Split(dataAsString, "\n") {
		them := int(line[0]) // 65
		us := int(line[2])   // 88
		score += us - 87
		switch (them - 65) - (us - 88) {
		case 0: // draw
			score += 3
		case 1, -2: // lose
			score += 0
		case -1, 2: // win
			score += 6
		}
	}

	fmt.Println("step1 -->", score)
}

func step2(filename string) {
	dataAsString := load(filename)
	score := 0
	for _, line := range strings.Split(dataAsString, "\n") {
		them := int(line[0])    // 65
		outcome := int(line[2]) // 88
		us := 0
		switch outcome - 88 {
		case 0: // lose
			us = them - 65 - 1
			if us == -1 {
				us = 2
			}
			score += 0
		case 1: // draw
			us = them - 65
			score += 3
		case 2: // win
			us = them - 65 + 1
			if us == 3 {
				us = 0
			}
			score += 6
		}
		score += us + 1
	}

	fmt.Println("step2 -->", score)
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
