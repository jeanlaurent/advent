package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExitRightAway(t *testing.T) {
	assert.Equal(t, 1, advent5(""))
}

func TestAdventExample(t *testing.T) {
	input := `0
3
0
1
-3`
	assert.Equal(t, 5, advent5(input))
}

func TestPerformFirstStep(t *testing.T) {
	program := []int{0, 3, 0, 1, -3}
	newIndex, newProgram := performStep(0, program)
	assert.Equal(t, 0, newIndex)
	assert.Equal(t, []int{1, 3, 0, 1, -3}, newProgram)
}

func TestPerformSecondStep(t *testing.T) {
	program := []int{1, 3, 0, 1, -3}
	newIndex, newProgram := performStep(0, program)
	assert.Equal(t, 1, newIndex)
	assert.Equal(t, []int{2, 3, 0, 1, -3}, newProgram)
}

func TestPerformThirdStep(t *testing.T) {
	program := []int{2, 3, 0, 1, -3}
	newIndex, newProgram := performStep(1, program)
	assert.Equal(t, 4, newIndex)
	assert.Equal(t, []int{2, 4, 0, 1, -3}, newProgram)
}

func TestPerformFourthStep(t *testing.T) {
	program := []int{2, 4, 0, 1, -3}
	newIndex, newProgram := performStep(4, program)
	assert.Equal(t, 1, newIndex)
	assert.Equal(t, []int{2, 4, 0, 1, -2}, newProgram)
}

func TestPerformLastStep(t *testing.T) {
	program := []int{2, 4, 0, 1, -2}
	newIndex, newProgram := performStep(1, program)
	assert.Equal(t, 5, newIndex)
	assert.Equal(t, []int{2, 5, 0, 1, -2}, newProgram)
}

func TestFinal(t *testing.T) {
	text, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(advent5(string(text)))
}

func performStep(index int, program []int) (int, []int) {
	newIndex := index + program[index]
	program[index]++
	return newIndex, program
}

func performStep52(index int, program []int) (int, []int) {
	newIndex := index + program[index]
	if program[index] >= 3 {
		program[index]--
	} else {
		program[index]++
	}
	return newIndex, program
}

func advent5(input string) int {
	if input == "" {
		return 1
	}
	instructionsAsString := strings.Split(input, "\n")
	program := []int{}
	for _, instruction := range instructionsAsString {
		if instruction != "" {
			number, err := strconv.Atoi(instruction)
			if err != nil {
				fmt.Println(err)
				return -1
			}
			program = append(program, number)
		}
	}
	index := 0
	count := 0
	for index < len(program) {
		count++
		index, program = performStep(index, program)
	}
	return count
}
