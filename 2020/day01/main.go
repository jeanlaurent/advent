package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := []int{}
	dataAsString := load("data.txt")
	for _, line := range strings.Split(dataAsString, "\n") {
		numbers = append(numbers, atoi(line))
	}
	fmt.Println("step1-->", step1(numbers))
	fmt.Println("step2-->", step2(numbers))
}

func step1(numbers []int) int {
	for _, x := range numbers {
		for _, y := range numbers {
			if x+y == 2020 {
				return x * y
			}
		}
	}
	return -1
}

func step2(numbers []int) int {
	for _, x := range numbers {
		for _, y := range numbers {
			for _, z := range numbers {
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}
	return -1
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
