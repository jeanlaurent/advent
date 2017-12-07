package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExitRightAway(t *testing.T) {
	firstCount, _ := advent6("")
	assert.Equal(t, 1, firstCount)
}

func TestPerformFirstStep(t *testing.T) {
	firstCount, secondCount := advent6("0 2 7 0")
	assert.Equal(t, 5, firstCount)
	assert.Equal(t, 4, secondCount)
}

func TestAdvent6(t *testing.T) {
	fmt.Println(advent6("10 3 15 10 5 15 5 15 9 2 5 8 5 2 3 6"))
}

func advent6(input string) (int, int) {
	if input == "" {
		return 1, 0
	}
	err, banks := convertString(input)
	if err != nil {
		fmt.Println(err)
		return -1, -1
	}
	memory := []string{}
	count := 0
	for !checkIfBanksAlreadyExist(banks, memory) {
		memory = append(memory, stringRepresentation(banks))
		indexBankToScatter := findBankToScatter(banks)
		banks = scatter(banks, indexBankToScatter)
		count++
	}
	fmt.Println(banks)
	fmt.Println("1st count", count)

	secondCount := 0
	bankToFind := stringRepresentation(banks)
	for secondCount == 0 || bankToFind != stringRepresentation(banks) {
		indexBankToScatter := findBankToScatter(banks)
		banks = scatter(banks, indexBankToScatter)
		secondCount++
	}
	fmt.Println(banks)
	fmt.Println("2nd count", secondCount)
	return count, secondCount
}

func stringRepresentation(array []int) string {
	arrayOfString := []string{}
	for _, value := range array {
		arrayOfString = append(arrayOfString, strconv.Itoa(value))
	}
	return strings.Join(arrayOfString, "-")
}

func checkIfBanksAlreadyExist(banks []int, memory []string) bool {
	bankAsString := stringRepresentation(banks)
	for _, oldBanks := range memory {
		if oldBanks == bankAsString {
			return true
		}
	}
	return false
}

func scatter(banks []int, index int) []int {
	valueToScatter := banks[index]
	banks[index] = 0
	for i := index + 1; i < valueToScatter+index+1; i++ {
		banks[i%len(banks)]++
	}
	return banks
}

func findBankToScatter(banks []int) int {
	biggestValue := -1
	biggestIndex := -1
	for index, value := range banks {
		if value > biggestValue {
			biggestValue = value
			biggestIndex = index
		}
	}
	return biggestIndex
}

func convertString(input string) (error, []int) {
	numbersAsString := strings.Split(input, " ")
	numbers := []int{}
	for _, numberAsString := range numbersAsString {
		if numberAsString != "" {
			number, err := strconv.Atoi(numberAsString)
			if err != nil {
				return err, nil
			}
			numbers = append(numbers, number)
		}
	}
	return nil, numbers
}
