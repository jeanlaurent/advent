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
	max := -1
	values := []int{}
	perms := permutations([]int{0, 1, 2, 3, 4})
	for _, perm := range perms {
		_, output1 := diagnostic(loadAndParse(), []int{perm[0], 0})
		_, output2 := diagnostic(loadAndParse(), []int{perm[1], output1})
		_, output3 := diagnostic(loadAndParse(), []int{perm[2], output2})
		_, output4 := diagnostic(loadAndParse(), []int{perm[3], output3})
		_, output5 := diagnostic(loadAndParse(), []int{perm[4], output4})
		if output5 > max {
			values = perm
			max = output5
		}
	}
	fmt.Println("step1 --> ", values, max)
}

func diagnostic(intcodes []int, input []int) ([]int, int) {
	index := 0
	indexInput := 0
	//output := ""
	output := -1
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
			intcodes[intcodes[index+1]] = input[indexInput]
			indexInput++
			index += 2
		case 4: // output
			output = getValueOfPointer(intcodes, param1, intcodes[index+1])
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
			return intcodes, output
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
	//dataAsString := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	//dataAsString := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	//dataAsString := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
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

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
