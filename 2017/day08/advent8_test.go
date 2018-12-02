package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleInput = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
`

func TestExitRightAway(t *testing.T) {
	assert.Equal(t, 0, advent81(""))
}

func TestSample(t *testing.T) {
	assert.Equal(t, 1, advent81(sampleInput))
}

func TestSampleSecondExercice(t *testing.T) {
	assert.Equal(t, 10, advent82(sampleInput))
}

func TestEquality(t *testing.T) {
	assert.True(t, condition(3, "==", 3))
	assert.False(t, condition(2, "==", 3))
}

func TestIncrement(t *testing.T) {
	assert.Equal(t, 5, operation(3, "inc", 2))
	assert.Equal(t, 1, operation(3, "dec", 2))
}

func TestFirstExercice(t *testing.T) {
	text, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(advent81(string(text)))
}

func TestSecondExercice(t *testing.T) {
	text, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(advent82(string(text)))
}

func operation(registerValue int, operator string, value int) int {
	fmt.Println(registerValue, operator, value)
	switch operator {
	case "inc":
		return registerValue + value
	case "dec":
		return registerValue - value
	}
	return 0
}

func condition(registerValue int, operator string, value int) bool {
	fmt.Println(registerValue, operator, value)
	switch operator {
	case "==":
		return registerValue == value
	case "!=":
		return registerValue != value
	case "<=":
		return registerValue <= value
	case ">=":
		return registerValue >= value
	case ">":
		return registerValue > value
	case "<":
		return registerValue < value
	}
	return false
}

func advent81(input string) int {
	if input == "" {
		return 0
	}
	registers := map[string]int{}
	fmt.Println(input)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		words := strings.Split(line, " ")
		fmt.Println(words)
		if condition(registers[words[4]], words[5], toInt(words[6])) {
			fmt.Println(true)
			fmt.Println(words[0], "is", registers[words[0]])
			registers[words[0]] = operation(registers[words[0]], words[1], toInt(words[2]))
			fmt.Println(words[0], "is now", registers[words[0]])
		}
	}
	biggest := -1
	for name, value := range registers {
		fmt.Println(name, value)
		if value > biggest {
			biggest = value
		}
	}
	return biggest
}

func advent82(input string) int {
	if input == "" {
		return 0
	}
	registers := map[string]int{}
	highest := -1
	fmt.Println(input)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		words := strings.Split(line, " ")
		fmt.Println(words)
		if condition(registers[words[4]], words[5], toInt(words[6])) {
			fmt.Println(true)
			fmt.Println(words[0], "is", registers[words[0]])
			registers[words[0]] = operation(registers[words[0]], words[1], toInt(words[2]))
			if registers[words[0]] > highest {
				highest = registers[words[0]]
				fmt.Println("Highest is now", words[0], "with", registers[words[0]])
			}
			fmt.Println(words[0], "is now", registers[words[0]])
		}
	}
	return highest
}

func toInt(valueAsString string) int {
	value, err := strconv.Atoi(valueAsString)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return value
}
