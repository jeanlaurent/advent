package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID) *optional*
*/

func main() {
	dataAsString := load("puzzle_input.txt")
	//dataAsString := load("sample_input.txt")
	//dataAsString := load("valid.txt")
	//dataAsString := load("invalid.txt")

	docs := []map[string]string{}

	currentDoc := map[string]string{}
	for _, line := range strings.Split(dataAsString, "\n") {
		if len(line) == 0 {
			docs = append(docs, currentDoc)
			currentDoc = map[string]string{}
			continue
		}
		for _, keypair := range strings.Split(line, " ") {
			items := strings.Split(keypair, ":")
			currentDoc[items[0]] = items[1]
		}
	}

	fmt.Println("step1 ==>", step1(docs))
	fmt.Println("step2 ==>", step2(docs))
}

func step1(docs []map[string]string) int {
	mandatory := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validDocCount := 0
	for _, doc := range docs {
		valid := 0
		for _, trait := range mandatory {
			_, exists := doc[trait]
			if exists {
				valid++
			}
		}
		if valid >= 7 {
			validDocCount++
		}
	}
	return validDocCount
}

func step2(docs []map[string]string) int {
	validDocCount := 0
	for _, doc := range docs {
		traitCount := 0
		if isValidBirthDate(doc["byr"]) {
			traitCount++
		}
		if isValidIssueDate(doc["iyr"]) {
			traitCount++
		}
		if isValidExpirationDate(doc["eyr"]) {
			traitCount++
		}
		if isValidHeight(doc["hgt"]) {
			traitCount++
		}
		if isValidHairColor(doc["hcl"]) {
			traitCount++
		}
		if isValidEyeColor(doc["ecl"]) {
			traitCount++
		}
		if isValidPid(doc["pid"]) {
			traitCount++
		}
		if traitCount >= 7 {
			validDocCount++
		}

	}
	return validDocCount
}

func isValidPid(color string) bool {
	if len(color) != 9 { // I'm so bad at regexp...
		return false
	}
	valid, _ := regexp.Match(`[0-9]{9}`, []byte(color))
	return valid
}

func isValidEyeColor(color string) bool {
	valid, _ := regexp.Match(`amb|blu|brn|gry|grn|hzl|oth`, []byte(color))
	return valid
}

func isValidHairColor(color string) bool {
	valid, _ := regexp.Match(`#[a-f0-9]{6}`, []byte(color))
	return valid
}

func isValidHeight(height string) bool {
	if strings.Contains(height, "in") {
		sizeAsString := height[0 : len(height)-2]
		return isValidDate(sizeAsString, 59, 76)
	}
	if strings.Contains(height, "cm") {
		sizeAsString := height[0 : len(height)-2]
		return isValidDate(sizeAsString, 150, 193)
	}
	return false
}

func isValidDate(yearAsString string, lowerBound, upperBound int) bool {
	year := atoi(yearAsString)
	return year >= lowerBound && year <= upperBound
}

func isValidBirthDate(yearAsString string) bool {
	if len(yearAsString) != 4 {
		return false
	}
	return isValidDate(yearAsString, 1920, 2002)
}

func isValidIssueDate(yearAsString string) bool {
	if len(yearAsString) != 4 {
		return false
	}
	return isValidDate(yearAsString, 2010, 2020)
}

func isValidExpirationDate(yearAsString string) bool {
	if len(yearAsString) != 4 {
		return false
	}
	return isValidDate(yearAsString, 2020, 2030)
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
