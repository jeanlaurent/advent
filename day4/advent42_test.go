package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test42ItRetursFalseWhenNoArg(t *testing.T) {
	assert.False(t, isValidAnagram(""))
}

func Test42Basic(t *testing.T) {
	assert.True(t, isValidAnagram("abcde fghij"))
}

func Test42Anagram(t *testing.T) {
	assert.False(t, isValidAnagram("abcde xyz ecdab"))
}

func TestSortString(t *testing.T) {
	assert.Equal(t, "aetw", sortString("weat"))
}

func Test42ItRetursTrueWhenMultipleLinesAlmostRepeat(t *testing.T) {
	myInput := `
bdwdjjo avricm cjbmj ran lmfsom ivsof
mxonybc fndyzzi gmdp gdfyoi inrvhr kpuueel wdpga vkq
`
	assert.Equal(t, 2, advent42(myInput))
}

func Test42ItRetursTrueWhenRealFile(t *testing.T) {
	text, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("1:", advent42(string(text)))
}

func advent42(input string) int {
	if input == "" {
		return 0
	}
	valid := 0
	for _, line := range strings.Split(input, "\n") {
		if isValidAnagram(line) {
			valid++
		}
	}
	return valid
}

func isValidAnagram(line string) bool {
	if line == "" {
		return false
	}
	oldTokens := []string{}
	for _, token := range strings.Split(line, " ") {
		for _, oldToken := range oldTokens {
			if sortString(oldToken) == sortString(token) {
				return false
			}
		}
		oldTokens = append(oldTokens, token)
	}
	return true
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
