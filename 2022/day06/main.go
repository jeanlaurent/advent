package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	step1("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	step1("bvwbjplbgvbhsrlpgdmjqwftvncz")
	step1("nppdvjthqldpwncqszvftbrmjlhg")
	step1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	step1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	step1(load("input.txt"))

	step2("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	step2("bvwbjplbgvbhsrlpgdmjqwftvncz")
	step2("nppdvjthqldpwncqszvftbrmjlhg")
	step2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	step2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	step2(load("input.txt"))
}

func step1(input string) {
	findStartOfBufferMarker(input, 4, 1)
}

func step2(input string) {
	findStartOfBufferMarker(input, 14, 2)
}

func findStartOfBufferMarker(input string, span int, step int) {
	for i := 0; i < len(input)-span; i++ {
		buffer := input[i : i+span]
		if !findDuplicate(buffer) {
			fmt.Println("step", step, " --> ", i+span, buffer)
			break
		}
	}
}

func findDuplicate(substring string) bool {
	for i := 0; i < len(substring); i++ {
		for j := i + 1; j < len(substring); j++ {
			if substring[i] == substring[j] {
				return true
			}
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
