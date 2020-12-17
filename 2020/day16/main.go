package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("step1        --> ", step1(load("sample_input.txt")))
	fmt.Println("step2-sample2 --> ", step2(load("sample2_input.txt")))
	fmt.Println("step1        --> ", step1(load("puzzle_input.txt")))
	fmt.Println("step2        --> ", step2(load("puzzle_input.txt")))
}

type bound struct {
	lower int
	upper int
}

type constraint struct {
	name   string
	bounds [2]bound
}

func (c *constraint) check(value int) bool {
	return (value >= c.bounds[0].lower && value <= c.bounds[0].upper) || (value >= c.bounds[1].lower && value <= c.bounds[1].upper)
}

func step2(dataAsString string) int {
	constraints, ownTickets, allNearbyTicket := readInput(dataAsString)

	nearbyTicket := [][]int{}
	// remove fully invalid ticket (not really needed I would think...)
	for _, ticket := range allNearbyTicket {
		valid := true
		for _, ticketValue := range ticket {
			invalidCount := 0
			for _, contraint := range constraints {
				if !contraint.check(ticketValue) {
					invalidCount++
				}
			}
			if invalidCount == len(constraints) {
				valid = false
				break
			}
		}
		if valid {
			nearbyTicket = append(nearbyTicket, ticket)
		}
	}

	// count all column which are invalid for each ticket for each constraint
	invalidList := map[string]map[int]bool{} //<-- what :D
	for _, constraint := range constraints {
		invalidList[constraint.name] = map[int]bool{}
		for _, ticket := range nearbyTicket {
			for columnIndex, val := range ticket {
				_, exist := invalidList[constraint.name][columnIndex]
				if exist {
					continue
				}
				if !constraint.check(val) {
					invalidList[constraint.name][columnIndex] = false
				}

			}
		}
	}
	// find every ticket that has but one invalid constraint
	// keep the record of that column and that constraint name
	// flag all ither ticket as invalid for that column
	// continue unil all tickets are found
	association := map[string]int{}
	for len(association) != len(constraints) {
		for _, constraint := range constraints {
			invalidCount := 0
			validColumn := -1
			for column := 0; column < len(ownTickets); column++ {
				_, exist := invalidList[constraint.name][column]
				if exist {
					invalidCount++
				} else {
					validColumn = column
				}
			}
			if invalidCount == len(ownTickets)-1 {
				association[constraint.name] = validColumn
				fmt.Println("Column", validColumn, "is associated with ", constraint.name)
				for _, constraint := range constraints {
					invalidList[constraint.name][validColumn] = false
				}
			}
		}
	}
	// compute result looping on the association map
	result := 1
	for _, constraint := range constraints {
		if strings.HasPrefix(constraint.name, "departure") {
			val := ownTickets[association[constraint.name]]
			fmt.Println(constraint.name, "col(", association[constraint.name], ") is", val)
			result *= val
		}
	}
	return result
}

func step1(dataAsString string) int {
	constraints, _, nearbyTicket := readInput(dataAsString)
	sum := 0
	for _, ticket := range nearbyTicket {
		for _, ticketValue := range ticket {
			invalidCount := 0
			for _, contraint := range constraints {
				if !contraint.check(ticketValue) {
					invalidCount++
				}
			}
			if invalidCount == len(constraints) {
				sum += ticketValue
			}
		}
	}
	return sum
}

func readInput(dataAsString string) ([]constraint, []int, [][]int) {
	readOwnTicket := false
	readNearbyTicket := false
	constraints := []constraint{}
	ownTicket := []int{}
	nearbyTicket := [][]int{}
	for _, line := range strings.Split(dataAsString, "\n") {
		if len(line) == 0 {
			continue
		}
		if line == "your ticket:" {
			readOwnTicket = true
			continue
		}
		if line == "nearby tickets:" {
			readOwnTicket = false
			readNearbyTicket = true
			continue
		}
		if readOwnTicket == false && readNearbyTicket == false {
			// read constraints
			parts := strings.Split(line, ":")
			constraint := constraint{name: parts[0], bounds: [2]bound{}}
			boundParts := strings.Split(parts[1], "or")
			for i := 0; i < 2; i++ {
				lowMinPart := strings.Split(boundParts[i], "-")
				lowerPart := atoi(strings.TrimSpace(lowMinPart[0]))
				upperPart := atoi(strings.TrimSpace(lowMinPart[1]))
				constraint.bounds[i] = bound{lower: lowerPart, upper: upperPart}
			}
			constraints = append(constraints, constraint)
		} else if readOwnTicket {
			for _, val := range strings.Split(line, ",") {
				ownTicket = append(ownTicket, atoi(val))
			}
		} else if readNearbyTicket {
			ticket := []int{}
			for _, val := range strings.Split(line, ",") {
				ticket = append(ticket, atoi(val))
			}
			nearbyTicket = append(nearbyTicket, ticket)
		}
	}
	return constraints, ownTicket, nearbyTicket
}

// ----

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
