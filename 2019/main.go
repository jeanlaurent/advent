package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	step1()
	step2()
}

func step1() {
	input := load()
	sumOfFuel := 0
	for _, line := range strings.Split(input, "\n") {
		mass, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		fuel := findFuelRequirement(mass)
		sumOfFuel += fuel
	}
	fmt.Println("step 1 -->", sumOfFuel)
}

func step2() {
	input := load()
	sumOfFuel := 0
	for _, line := range strings.Split(input, "\n") {
		mass, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		for mass > 0 {
			fuelRequirement := findFuelRequirement(mass)
			if fuelRequirement > 0 {
				sumOfFuel += fuelRequirement
			}
			mass = fuelRequirement
		}
	}
	fmt.Println("step 2 -->", sumOfFuel)
}

func findFuelRequirement(mass int) int {
	return int(math.Floor(float64(mass/3)) - 2)
}

func load() string {
	text, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}
