package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	step1(load("small.txt"))
	step1(load("input.txt"))
	step2(load("small.txt"))
	step2(load("input.txt"))
}

type monkey struct {
	items   []int
	op      string
	opvalue int
	test    int
	success int
	failure int
}

func step1(input string) {
	monkeys, _ := readMonkeys(input)
	active := make([]int, len(monkeys))
	for round := 0; round < 20; round++ {
		for n := 0; n < len(monkeys); n++ {
			//fmt.Println("Monkey", n)
			for i := 0; i < len(monkeys[n].items); i++ {
				active[n]++
				//fmt.Println("  Monkey inspects an item with a worry level of", monkeys[n].items[i])
				switch monkeys[n].op {
				case "*":
					monkeys[n].items[i] *= monkeys[n].opvalue
					//fmt.Println("    Worry level is multiplied by ", monkeys[n].opvalue, " to ", monkeys[n].items[i], ".")
				case "+":
					monkeys[n].items[i] += monkeys[n].opvalue
					//fmt.Println("    Worry level increases by ", monkeys[n].opvalue, " to ", monkeys[n].items[i], ".")
				case "^":
					monkeys[n].items[i] *= monkeys[n].items[i]
					//fmt.Println("    Worry level is multiplied by itself to ", monkeys[n].items[i], ".")
				}
				monkeys[n].items[i] /= 3
				//fmt.Println("    Monkey gets bored with item. Worry level is divided by 3 to ", monkeys[n].items[i], ".")
				if monkeys[n].items[i]%monkeys[n].test == 0 {
					//fmt.Println("    Current worry level is divisible by", monkeys[n].test, ".")
					//fmt.Println("    Item with worry level ", monkeys[n].items[i], " is thrown to monkey ", monkeys[n].success, ".")
					monkeys[monkeys[n].success].items = append(monkeys[monkeys[n].success].items, monkeys[n].items[i])
				} else {
					//fmt.Println("    Current worry level is not divisible by", monkeys[n].test, ".")
					//fmt.Println("    Item with worry level ", monkeys[n].items[i], " is thrown to monkey ", monkeys[n].failure, ".")
					monkeys[monkeys[n].failure].items = append(monkeys[monkeys[n].failure].items, monkeys[n].items[i])
				}

			}
			monkeys[n].items = []int{}
		}
		//fmt.Println("round", round)
		//printMonkeyItems(monkeys)
	}
	activeSlice := active[:]
	sort.Sort(sort.Reverse(sort.IntSlice(activeSlice)))
	fmt.Println("step1 -->", activeSlice[0]*activeSlice[1])
}

func step2(input string) {
	monkeys, bigDiviser := readMonkeys(input)
	active := make([]int, len(monkeys))
	for round := 0; round < 10000; round++ {
		for n := 0; n < len(monkeys); n++ {
			for i := 0; i < len(monkeys[n].items); i++ {
				active[n]++
				switch monkeys[n].op {
				case "*":
					monkeys[n].items[i] *= monkeys[n].opvalue
				case "+":
					monkeys[n].items[i] += monkeys[n].opvalue
				case "^":
					monkeys[n].items[i] *= monkeys[n].items[i]
				}
				// monkeys[n].items[i] /= 3
				// Did try multiple failed idea until I looked at reddit for a hint.
				monkeys[n].items[i] %= bigDiviser
				if monkeys[n].items[i]%monkeys[n].test == 0 {
					monkeys[monkeys[n].success].items = append(monkeys[monkeys[n].success].items, monkeys[n].items[i])
				} else {
					monkeys[monkeys[n].failure].items = append(monkeys[monkeys[n].failure].items, monkeys[n].items[i])
				}

			}
			monkeys[n].items = []int{}
		}
		if round == 0 || round == 19 || round == 999 || round == 1999 {
			fmt.Println("round", round+1)
			for i := 0; i < len(active); i++ {
				fmt.Println("Monkey", i, "inspected items", active[i], "times.")
			}
		}
	}
	activeSlice := active[:]
	sort.Sort(sort.Reverse(sort.IntSlice(activeSlice)))
	fmt.Println("step2 -->", activeSlice[0]*activeSlice[1])
}

func printMonkeyItems(monkeys []monkey) {
	for n, monkey := range monkeys {
		fmt.Print("Monkey", n, ":")
		for _, item := range monkey.items {
			fmt.Print(item, ",")
		}
		fmt.Println()
	}
}

func readMonkeys(input string) ([]monkey, int) {
	bigDiviser := 1
	monkeys := []monkey{}
	lines := strings.Split(input, "\n")
	i := 0
	for i < len(lines) {
		monkey := monkey{[]int{}, "*", 0, 0, 0, 0}
		i++
		lines[i] = string(strings.TrimSpace(lines[i]))
		itemLines := strings.Split(lines[i], ":")
		for _, item := range strings.Split(itemLines[1], ",") {
			monkey.items = append(monkey.items, atoi(strings.TrimSpace(item)))
		}
		i++
		opLines := strings.Split(lines[i], "=")
		monkey.op = string(opLines[1][5])
		operand := string(opLines[1][7:len(opLines[1])])
		if operand == "old" {
			monkey.op = "^"
			monkey.opvalue = -1
		} else {
			monkey.opvalue = atoi(string(opLines[1][7:len(opLines[1])]))
		}
		i++
		fmt.Sscanf(lines[i], "  Test: divisible by %d", &(monkey.test))
		// only needed for step2
		bigDiviser *= monkey.test
		i++
		fmt.Sscanf(lines[i], "  If true: throw to monkey %d", &(monkey.success))
		i++
		fmt.Sscanf(lines[i], "   If false: throw to monkey %d", &(monkey.failure))
		i += 2
		monkeys = append(monkeys, monkey)
	}
	return monkeys, bigDiviser
}

func atoi(line string) int {
	number, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return number
}

func load(filename string) string {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}
