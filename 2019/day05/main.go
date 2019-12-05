package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("step1")
	intcodes := loadAndParse()
	diagnostic(intcodes, 1)

	fmt.Println("step2")
	intcodes = loadAndParse()
	diagnostic(intcodes, 5)
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
		case 1: // addition
			firstParameter := getValueOfPointer(intcodes, param1, intcodes[index+1])
			secondParameter := getValueOfPointer(intcodes, param2, intcodes[index+2])
			value := firstParameter + secondParameter
			intcodes[intcodes[index+3]] = value
			index += 4
		case 2: // multiplication
			firstParameter := getValueOfPointer(intcodes, param1, intcodes[index+1])
			secondParameter := getValueOfPointer(intcodes, param2, intcodes[index+2])
			value := firstParameter * secondParameter
			intcodes[intcodes[index+3]] = value
			index += 4
		case 3: // store input
			intcodes[intcodes[index+1]] = input
			index += 2
		case 4: // output
			fmt.Println(getValueOfPointer(intcodes, param1, intcodes[index+1]))
			index += 2
		case 5: // jump if true
			condition := getValueOfPointer(intcodes, param1, intcodes[index+1])
			if condition != 0 {
				index = getValueOfPointer(intcodes, param2, intcodes[index+2])
			} else {
				index += 3
			}
		case 6: // jump if false
			condition := getValueOfPointer(intcodes, param1, intcodes[index+1])
			if condition == 0 {
				index = getValueOfPointer(intcodes, param2, intcodes[index+2])
			} else {
				index += 3
			}
		case 7: // less than
			firstParameter := getValueOfPointer(intcodes, param1, intcodes[index+1])
			secondParameter := getValueOfPointer(intcodes, param2, intcodes[index+2])
			if firstParameter < secondParameter {
				intcodes[intcodes[index+3]] = 1
			} else {
				intcodes[intcodes[index+3]] = 0
			}
			index += 4
		case 8: // equal than
			firstParameter := getValueOfPointer(intcodes, param1, intcodes[index+1])
			secondParameter := getValueOfPointer(intcodes, param2, intcodes[index+2])
			if firstParameter == secondParameter {
				intcodes[intcodes[index+3]] = 1
			} else {
				intcodes[intcodes[index+3]] = 0
			}
			index += 4
		case 99:
			return intcodes
		default:
			fmt.Println("invalid intcode", intcodes[index], "at", index)
			os.Exit(-1)
		}
	}

}

func getValueOfPointer(intcodes []int, param int, value int) int {
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
