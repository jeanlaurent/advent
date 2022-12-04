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
	score := 0

	for _, rucksack := range strings.Split(load(filename), "\n") {

		firstHalf := rucksack[0:len(rucksack)/2]
		secondHalf := rucksack[len(rucksack)/2:]

		common := ""
		for _, char1 := range firstHalf {
			for _, char2 := range secondHalf {
				if char1 == char2 {
					common = string(char1)
				}
			}
		}

		ascii := byte(common[0])
		if ascii >= byte('a') {
			score += int(ascii - byte('a')) + 1
		} else {
			score += int(ascii) - int(byte('A')) + 27
		}

	}

	fmt.Println("step1 -->", score)
}

func step2(filename string) {
	score := 0

	group := []string{}
	for _, rucksack := range strings.Split(load(filename), "\n") {
		if len(group) < 2 {
			group = append(group,rucksack)
		} else {
			group = append(group,rucksack)
			common := ""
			for _, char1 := range group[0] {
				for _, char2 := range group[1] {
					if char1 == char2 {
						for _, char3 := range group[2] {
							if char2 == char3 {
								common = string(char2)
							}
						}	
					}
				}
			}
			ascii := byte(common[0])

			if ascii >= byte('a') {
				score += int(ascii - byte('a')) + 1
			} else {
				score += int(ascii) - int(byte('A')) + 27
			}

			group = []string{}				
		}
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
