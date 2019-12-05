package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// 3,0,4,0,99
// 1002,4,3,4,33

func main() {
	intcodes := loadAndParse()
	//intcodes := []int{3, 0, 4, 0, 99}
	// intcodes := []int{1002, 4, 3, 4, 33}
	diagnostic(intcodes, 1)
}

func diagnostic(intcodes []int, input int) []int {
	index := 0
	for {
		// readOpCode and Param
		opcode := intcodes[index] % 100
		param1 := 0
		if intcodes[index] > 99 {
			param1 = getDigit(intcodes[index], 3)
		}
		param2 := 0
		if intcodes[index] > 999 {
			param2 = getDigit(intcodes[index], 4)
		}
		switch opcode {
		case 1:
			value := getValueOfPointer(intcodes, index, param1, intcodes[index+1]) + getValueOfPointer(intcodes, index, param2, intcodes[index+2])
			intcodes[intcodes[index+3]] = value
			index += 4
		case 2:
			value := getValueOfPointer(intcodes, index, param1, intcodes[index+1]) * getValueOfPointer(intcodes, index, param2, intcodes[index+2])
			intcodes[intcodes[index+3]] = value
			index += 4
		case 3:
			intcodes[intcodes[index+1]] = input
			index += 2
		case 4:
			fmt.Println(intcodes[intcodes[index+1]])
			index += 2
		case 99:
			return intcodes
		default:
			fmt.Println("invalid intcode", intcodes[index], "at", index)
			os.Exit(-1)
		}
	}

}

func getValueOfPointer(intcodes []int, index int, param int, value int) int {
	if param == 0 {
		return intcodes[value]
	}
	return value
}

// utilities
func loadAndParse() []int {
	ops := []int{}
	dataAsString := load("input.txt")
	for _, opsAsString := range strings.Split(dataAsString, ",") {
		ops = append(ops, atoi(opsAsString))
	}
	return ops
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

func getDigit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
