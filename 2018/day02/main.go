package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	input := load()
	step1(input)
	step2(input)
}

func step2(input string) {
	ids := []string{}
	for _, line := range strings.Split(input, "\n") {
		ids = append(ids, line)
	}
	for i := 0; i < len(ids); i++ {
		for j := 0; j < len(ids); j++ {
			if i == j {
				break
			}
			diff := compare(ids[i], ids[j])
			if diff == 1 {
				answer := strings.Builder{}
				for k := 0; k < len(ids[i]); k++ {
					if ids[i][k] == ids[j][k] {
						answer.WriteByte(ids[i][k])
					}
				}
				fmt.Println("Step2:", answer.String())
			}
		}
	}
}

func compare(one string, two string) int {
	differ := 0
	for i := 0; i < len(one); i++ {
		if one[i] != two[i] {
			differ++
		}
	}
	return differ
}

func step1(input string) {
	doubleNumber := 0
	doubleFound := false
	tripleNumber := 0
	tripleFound := false
	for _, line := range strings.Split(input, "\n") {
		sorted := sortString(line)
		for i := 0; i < len(sorted); {
			if i+1 < len(sorted) && sorted[i] == sorted[i+1] {
				if i+2 < len(sorted) && sorted[i] == sorted[i+2] {
					if !tripleFound {
						tripleNumber++
						tripleFound = true
					}
					i++
				} else {
					if !doubleFound {
						doubleNumber++
						doubleFound = true
					}
				}
			}
			i++
		}
		doubleFound = false
		tripleFound = false

	}
	fmt.Println("Step1:", doubleNumber*tripleNumber)
}

func load() string {
	text, err := ioutil.ReadFile("./input02.txt")
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
