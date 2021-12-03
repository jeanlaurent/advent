package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("step1 ==>", step1(load("puzzle_input.txt")))
	fmt.Println("step2 ==>", step2(load("puzzle_input.txt")))
}

func computeArray2(input []string) int {
	multOnlyInput := []string{}
	ops := ""
	result := 0
	for tokenIndex := 0; tokenIndex < len(input); tokenIndex++ {
		token := string(input[tokenIndex])
		fmt.Println("token is ", token, " result is ", result, "ops is ", ops, "multiops", multOnlyInput)
		switch token {
		case "+":
			fmt.Println("tracking +")
			ops = "+"
		case "*":
			ops = ""
			multOnlyInput = append(multOnlyInput, itoa(result))
			multOnlyInput = append(multOnlyInput, "*")
			result = 0
			fmt.Println("found a * adding everything we got to string", multOnlyInput)
		case "(":
			fmt.Println("chasing )")
			indexClosing := tokenIndex + 1
			groundZero := 0
			for input[indexClosing] != ")" || groundZero != 0 {
				fmt.Println("chasing ) found", input[indexClosing], "level", groundZero)
				if input[indexClosing] == "(" {
					groundZero++
					fmt.Println("since we found ( level is now", groundZero)
				}
				if input[indexClosing] == ")" {
					groundZero--
					fmt.Println("since we found ) level is now", groundZero)
				}
				indexClosing++
			}
			fmt.Println("parent expression is ", input[tokenIndex+1:indexClosing])
			intermediaryResult := computeArray2(input[tokenIndex+1 : indexClosing])
			tokenIndex = indexClosing
			if ops == "" {
				result = intermediaryResult
			} else {
				fmt.Println("found a + ", intermediaryResult)
				result += intermediaryResult
				fmt.Println("result is now", result)
				ops = ""
			}

		default:
			if ops == "" {
				fmt.Println("putting", input[tokenIndex], "into result")
				result = atoi(input[tokenIndex])
			} else {
				fmt.Println("found a + ", input[tokenIndex])
				result += atoi(input[tokenIndex])
				fmt.Println("result is now", result)
				ops = ""
			}
		}
	}
	if result > 0 {
		multOnlyInput = append(multOnlyInput, itoa(result))
	}
	fmt.Println("passing", multOnlyInput)
	return computeArray(multOnlyInput)
}

func computeArray(input []string) int {
	if len(input) == 1 {
		return atoi(input[0])
	}
	ops := ""
	result := 0
	for tokenIndex := 0; tokenIndex < len(input); tokenIndex++ {
		token := string(input[tokenIndex])
		fmt.Println("token is ", token, " result is ", result, "ops is ", ops)
		switch token {
		case "(":
			fmt.Println("chasing )")
			indexClosing := tokenIndex + 1
			groundZero := 0
			for input[indexClosing] != ")" || groundZero != 0 {
				fmt.Println("chasing ) found", input[indexClosing], "level", groundZero)
				if input[indexClosing] == "(" {
					groundZero++
					fmt.Println("since we found ( level is now", groundZero)
				}
				if input[indexClosing] == ")" {
					groundZero--
					fmt.Println("since we found ) level is now", groundZero)
				}
				indexClosing++
			}
			fmt.Println("parent expression is ", input[tokenIndex+1:indexClosing])
			parResult := computeArray(input[tokenIndex+1 : indexClosing])
			tokenIndex = indexClosing
			if ops == "" {
				result = parResult
			} else {
				if ops == "+" {
					result += parResult
				} else {
					result *= parResult
				}
				ops = ""
			}
		case "+":
			ops = "+"
		case "*":
			ops = "*"
		default:
			if ops == "" {
				result = atoi(input[tokenIndex])
			} else {
				if ops == "+" {
					fmt.Println("+", input[tokenIndex])
					result += atoi(input[tokenIndex])
				} else {
					result *= atoi(input[tokenIndex])
				}
				ops = ""
			}
		}
	}
	fmt.Println("returning", result)
	return result
}

func compute(input string) int {
	fmt.Println("------")
	newInput := strings.Fields(strings.NewReplacer("(", " ( ", ")", " ) ").Replace(input))
	return computeArray(newInput)
}

func compute2(input string) int {
	fmt.Println("------")
	newInput := strings.Fields(strings.NewReplacer("(", " ( ", ")", " ) ").Replace(input))
	return computeArray2(newInput)
}

func step1(dataAsString string) int {
	sum := 0
	for _, line := range strings.Split(dataAsString, "\n") {
		sum += compute(line)
	}
	return sum
}

func step2(dataAsString string) int {
	sum := 0
	for _, line := range strings.Split(dataAsString, "\n") {
		sum += compute2(line)
	}
	return sum
}

func find(name []string, target string) bool {
	for i := 0; i < len(name); i++ {
		if name[i] == target {
			return true
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

func itoa(line int) string {
	return strconv.Itoa(line)
}
