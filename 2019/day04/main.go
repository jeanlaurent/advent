package main

import (
	"fmt"
	"math"
)

func main() {
	passwordCount := []int{0, 0}
	for candidate := 245182; candidate < 790572; candidate++ {
		if isValidStep1(candidate) {
			passwordCount[0]++
		}
		if isValidStep2(candidate) {
			passwordCount[1]++
		}
	}
	fmt.Println("step1 -->", passwordCount[0])
	fmt.Println("step2 -->", passwordCount[1])
}

func isValidStep1(candidate int) bool {
	previousDigit := getDigit(candidate, 6)
	double := false
	isValid := true
	for digitIndex := 5; digitIndex > 0; digitIndex-- {
		digit := getDigit(candidate, digitIndex)
		if digit == previousDigit {
			double = true
			previousDigit = digit
			continue
		}
		if digit < previousDigit {
			isValid = false
			break
		}
		previousDigit = digit
	}
	return isValid && double
}

func isValidStep2(candidate int) bool {
	previousDigit := getDigit(candidate, 6)
	groups := []int{}
	repeat := 1
	isValid := true
	for digitIndex := 5; digitIndex > 0; digitIndex-- {
		digit := getDigit(candidate, digitIndex)
		if digit == previousDigit {
			repeat++
			previousDigit = digit
			continue
		}
		if repeat > 1 {
			groups = append(groups, repeat)
			repeat = 1
		}
		if digit < previousDigit {
			isValid = false
			break
		}
		previousDigit = digit
	}
	if repeat > 1 {
		groups = append(groups, repeat)
	}
	for index := 0; index < len(groups); index++ {
		if groups[index] == 2 {
			return isValid
		}
	}
	return false
}

func getDigit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
