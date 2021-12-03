package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(readRulesAndData("sample_puzzle.txt"))
	// fmt.Println("step1 ==>", step1(load("puzzle_input.txt")))
	// fmt.Println("step2 ==>", step2(load("puzzle_input.txt")))
}

func readRulesAndData(fileName string) (map[string]string, []string) {
	rules := map[string]string{}
	data := []string{}

	dataAsString := load(fileName)
	rulesLoading := true

	for _, line := range strings.Split(dataAsString, "\n") {
		if len(line) == 0 {
			rulesLoading = false
		}
		if rulesLoading {
			parts := strings.Split(line, ":")
			rules[parts[0]] = parts[1]
		} else {
			data = append(data, line)
		}
	}

	return rules, data
}

func step1(fileName string) int {
	rules, data := readRulesAndData(fileName)

	return -1
}

func step2() int {
	return -1
}

//----

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

func itoa(line int) string {
	return strconv.Itoa(line)
}
