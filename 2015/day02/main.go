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

	fmt.Println("step1 --> ", step1(dataAsString))
	fmt.Println("step2 --> ", step2(dataAsString))
}

func step1(dataAsString string) int {
	giftSurface := 0
	for _, line := range strings.Split(dataAsString, "\n") {
		dimension := strings.Split(line, "x")
		l := atoi(dimension[0])
		w := atoi(dimension[1])
		h := atoi(dimension[2])
		sides := [3]int{l * w, w * h, h * l}
		smallest := 9999999
		for i := 0; i < len(sides); i++ {
			if smallest > sides[i] {
				smallest = sides[i]
			}
		}
		giftSurface += 2*l*w + 2*w*h + 2*h*l + smallest
	}
	return giftSurface
}

func step2(dataAsString string) int {
	giftSurface := 0
	for _, line := range strings.Split(dataAsString, "\n") {
		dimension := strings.Split(line, "x")
		l := atoi(dimension[0])
		w := atoi(dimension[1])
		h := atoi(dimension[2])
		giftSurface += 2*l + 2*w + l*w*h
	}
	return giftSurface
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
