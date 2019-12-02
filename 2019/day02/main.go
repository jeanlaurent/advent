package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,10,23,1,23,13,27,1,6,27,31,1,9,31,35,2,10,35,39,1,39,6,43,1,6,43,47,2,13,47,51,1,51,6,55,2,6,55,59,2,59,6,63,2,63,13,67,1,5,67,71,2,9,71,75,1,5,75,79,1,5,79,83,1,83,6,87,1,87,6,91,1,91,5,95,2,10,95,99,1,5,99,103,1,10,103,107,1,107,9,111,2,111,10,115,1,115,9,119,1,13,119,123,1,123,9,127,1,5,127,131,2,13,131,135,1,9,135,139,1,2,139,143,1,13,143,0,99,2,0,14,0"

//const input = "1,9,10,3,2,3,11,0,99,30,40,50"
//const input = "1,0,0,0,99"
//const input = "2,3,0,3,99"
//const input = "2,4,4,5,99,0"
//const input = "1,1,1,4,99,5,6,0,99"

func main() {
	step1(load())
	step2(load())
}

func load() []int64 {
	intcodes := []int64{}
	chars := strings.Split(input, ",")
	for _, char := range chars {
		intcode, err := strconv.Atoi(char)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		intcodes = append(intcodes, int64(intcode))
	}
	return intcodes
}

func step1(intcodes []int64) {
	intcodes[1] = 12
	intcodes[2] = 2
	fmt.Println(intcodes)
	fmt.Println(gravityAssist(intcodes))
	fmt.Println("step 1 --> ", intcodes[0])
}

func step2(inputIntCodes []int64) []int64 {
	intcodes := make([]int64, len(inputIntCodes))
	for i := 1; i < 100; i++ {
		for j := 1; j < 100; j++ {
			copy(intcodes, inputIntCodes)
			intcodes[1] = int64(i)
			intcodes[2] = int64(j)
			gravityAssist(intcodes)
			if intcodes[0] == 19690720 {
				fmt.Println("step 2 --> ", 100*intcodes[1]+intcodes[2])
				return intcodes
			}
		}
	}
	return intcodes
}

func gravityAssist(intcodes []int64) []int64 {
	index := 0
	for {
		if intcodes[index] == 99 {
			return intcodes
		}
		if intcodes[index] == 1 {
			value := intcodes[intcodes[index+1]] + intcodes[intcodes[index+2]]
			intcodes[intcodes[index+3]] = value
		} else if intcodes[index] == 2 {
			value := intcodes[intcodes[index+1]] * intcodes[intcodes[index+2]]
			intcodes[intcodes[index+3]] = value
		} else {
			fmt.Println("invalid intcode", intcodes[index], "at", index)
		}
		index += 4
	}
}
