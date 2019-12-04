package main

import (
	"fmt"
	"math"
)

// s1 input 245182-790572
func main() {
	fmt.Println("isValid 111111", isValid(111111))
	fmt.Println("isValid 113456", isValid(113456))
	fmt.Println("isValid 789997", isValid(789997))

	passwordCount := 0
	for candidate := 245182; candidate < 790572; candidate++ { //245190; candidate++ {
		if isValid(candidate) {
			fmt.Println(candidate)
			passwordCount++
		}
	}
	fmt.Println("step1 -->", passwordCount)
}

func isValid(candidate int) bool {
	previousDigit := getDigit(candidate, 6)
	double := false
	isValid := true
	for digitIndex := 5; digitIndex > 0; digitIndex-- {
		digit := getDigit(candidate, digitIndex)
		// fmt.Print(candidate, digit)
		if digit == previousDigit {
			// fmt.Println("==")
			double = true
			previousDigit = digit
			continue
		}
		// fmt.Println("compare", digit, "<", previousDigit)
		if digit < previousDigit {
			// fmt.Println("<")
			isValid = false
			break
		}
		previousDigit = digit
	}
	return isValid && double
}

func getDigit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
