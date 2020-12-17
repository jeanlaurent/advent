package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type bus struct {
	id     int
	offset int
}

func main() {
	dataAsString := load("sample_input.txt")
	fmt.Println("step1 (sample) -->", step1(dataAsString))

	fmt.Println("step2 (sample) -->", step2(getBusString(dataAsString), 0))
	fmt.Println("step2 (sample2) -->", step2("17,x,13,19", 0))
	fmt.Println("step2 (sample3) -->", step2("67,7,59,61", 0))

	fmt.Println("step2 (sample4) -->", step2("67,x,7,59,61", 0))
	fmt.Println("step2 (sample5) -->", step2("67,7,x,59,61", 0))
	fmt.Println("step2 (sample6) -->", step2("1789,37,47,1889", 0))

	// 	67,x,7,59,61 first occurs at timestamp 779210.
	// 67,7,x,59,61 first occurs at timestamp 1261476.
	// 1789,37,47,1889 first occurs at timestamp 1202161486.

	dataAsString = load("puzzle_input.txt")
	fmt.Println("step1 -->", step1(dataAsString))
	fmt.Println("step2 -->", step2(getBusString(dataAsString), 100000000000000))
}

func step2(busString string, startTime int) int {
	allBus := []bus{}
	for offset, busIdAsString := range strings.Split(busString, ",") {
		if busIdAsString == "x" {
			continue
		}
		fmt.Println(busIdAsString)
		allBus = append(allBus, bus{id: atoi(busIdAsString), offset: offset})
	}
	departureTime := 0
	found := false
	for !found {
		if departureTime%10000000000000 == 1 {
			fmt.Println("tick", departureTime)
		}

		departureTime += allBus[0].id
		found = true
		for _, bus := range allBus[1:] {
			if (departureTime+bus.offset)%bus.id != 0 {
				found = false
				break
			}
		}
	}
	return departureTime
}

func getBusString(dataAsString string) string {
	parts := strings.Split(dataAsString, "\n")
	return parts[1]
}

func step1(dataAsString string) int {
	parts := strings.Split(dataAsString, "\n")
	biggestRest := 0
	closestBudId := -1
	timeToDepart := atoi(parts[0])
	for _, busIdAsString := range strings.Split(parts[1], ",") {
		if busIdAsString == "x" {
			continue
		}
		busId := atoi(busIdAsString)
		rest := timeToDepart % busId
		if rest > biggestRest {
			biggestRest = rest
			closestBudId = busId
		}

	}
	delta := (timeToDepart/closestBudId+1)*closestBudId - timeToDepart
	return delta * closestBudId
}

// ----

func abs(i int) int {
	if i > 0 {
		return i
	}
	return i * -1
}

func exist(array []int, item int) bool {
	for _, element := range array {
		if element == item {
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
