package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleInput = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)
`

func TestExitRightAway(t *testing.T) {
	assert.Equal(t, "", advent71(""))
}

func TestWithSampleInput(t *testing.T) {
	assert.Equal(t, "tknk", advent71(sampleInput))
}

func TestFirstExercice(t *testing.T) {
	text, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(advent71(string(text)))
}

func advent71(input string) string {
	if input == "" {
		return ""
	}
	lefts := []string{}
	rights := []string{}
	for _, line := range strings.Split(input, "\n") {
		indexFirstSpace := strings.Index(line, " ")
		if indexFirstSpace != -1 {
			word := line[0:indexFirstSpace]
			lefts = append(lefts, word)
		}
		indexArrow := strings.Index(line, ">")
		if indexArrow != -1 {
			rightList := line[indexArrow+2 : len(line)]
			for _, right := range strings.Split(rightList, ", ") {
				rights = append(rights, right)
			}
		}
	}
	for _, left := range lefts {
		found := false
		for _, right := range rights {
			if right == left {
				found = true
				break
			}
		}
		if !found {
			return left
		}
	}
	return ""
}
