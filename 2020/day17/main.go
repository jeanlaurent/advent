package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type cube struct {
	x      int
	y      int
	z      int
	active bool
}

func main() {
	fmt.Println("step1 -->", step1(load("sample-input.txt")))
}

func step1(dataAsString string) int {
	cubes := [][][]bool{}
	for y, line := range strings.Split(dataAsString, "\n") {
		cubes[0] = []bool{}
		for x := 0; x < len(line); x++ {
			cubes[0][y*100+x] = append(cubes[0][y*100+x], string(line[x]) == "#")
		}
	}
	for y := 0; y < len(cubes[z]); y++ {
		for x := 0; x < len(cubes[z][y]); x++ {
			if cubes[z][y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return -1
}

// ----

func exist(array []int, item int) bool {
	for _, element := range array {
		if element == item {
			return true
		}
	}
	return false
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
