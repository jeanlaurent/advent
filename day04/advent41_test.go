package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItRetursFalseWhenNoArg(t *testing.T) {
	assert.False(t, isValid(""))
}

func TestItRetursTrueWhenNoRepeat(t *testing.T) {
	assert.True(t, isValid("aa bb cc dd ee"))
}

func TestItRetursFalseWhenOneRepeat(t *testing.T) {
	assert.False(t, isValid("aa bb cc dd ee aa"))
}

func TestItRetursTrueWhenOneAlmostRepeat(t *testing.T) {
	assert.True(t, isValid("aa bb cc dd aaa"))
}

func TestItRetursTrueWhenMultipleLinesAlmostRepeat(t *testing.T) {
	myInput := `
bdwdjjo avricm cjbmj ran lmfsom ivsof
mxonybc fndyzzi gmdp gdfyoi inrvhr kpuueel wdpga vkq
`
	assert.Equal(t, 2, advent3(myInput))
}

func TestItRetursTrueWhenRealFile(t *testing.T) {
	text, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("1:", advent3(string(text)))
}

func advent3(input string) int {
	if input == "" {
		return 0
	}
	valid := 0
	for _, line := range strings.Split(input, "\n") {
		if isValid(line) {
			valid++
		}
	}
	return valid
}

func isValid(line string) bool {
	if line == "" {
		return false
	}
	oldTokens := []string{}
	for _, token := range strings.Split(line, " ") {
		for _, oldToken := range oldTokens {
			if oldToken == token {
				return false
			}
		}
		oldTokens = append(oldTokens, token)
	}
	return true
}
