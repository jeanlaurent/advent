package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// step1("small.txt")
	step1("input.txt")
	// step2("small.txt")
	step2("input.txt")
}

func step1(filename string) {
	numberOfIncludedPairs := 0

	for _, line := range strings.Split(load(filename), "\n") {
		pairs := strings.Split(line,",")
		low1, high1 := getPairBound(pairs[0])
		low2, high2 := getPairBound(pairs[1])
		if (low1 <= low2 && high1 >= high2) || (low1 >= low2 && high1 <= high2) {
			numberOfIncludedPairs++
		}
	}

	fmt.Println("step1 (",filename,")-->", numberOfIncludedPairs)
}

func getPairBound(pair string) (int, int) {
	bounds := strings.Split(pair,"-")
	low := atoi(bounds[0])
	high := atoi(bounds[1])
	return low, high
}


func step2(filename string) {
	numberOfOverlappingPairs := 0

	for _, line := range strings.Split(load(filename), "\n") {
		pairs := strings.Split(line,",")
		low1, high1 := getPairBound(pairs[0])
		low2, high2 := getPairBound(pairs[1])
		// if (high1 < low2) || (high2 < low1) {
		// 	// do nothing
		// } else {
		// 	numberOfOverlappingPairs++
		// }
		if (high1 >= low2) && (high2 >= low1) {
			numberOfOverlappingPairs++
		} 
	}

	fmt.Println("step2 (",filename,")-->", numberOfOverlappingPairs)
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