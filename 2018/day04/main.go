package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type event struct {
	timeStamp   string
	text        string
	minute      int
	guardNumber string
	action      int
}

const beginShift = 1
const fallAsleep = 2
const wakeUp = 3

func main() {
	input := load()
	step1(input)
	step2(input)
}

func buildEventList(input string) []event {
	events := []event{}
	for _, line := range strings.Split(input, "\n") {
		guardNumber := ""
		action := -1
		index := strings.Index(line, "]")
		timestamp := line[1:index]

		minute := toNumber(timestamp[14:16])

		text := line[index+2:]
		if strings.Contains(text, "Guard") {
			//guardNumber = text[7:9]
			fmt.Sscanf(text, "Guard #%s begins shift", &guardNumber)
			action = beginShift
		}
		if strings.Contains(text, "asleep") {
			action = fallAsleep
		}
		if strings.Contains(text, "wakes") {
			action = wakeUp
		}
		events = append(events, event{timestamp, text, minute, guardNumber, action})
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].timeStamp < events[j].timeStamp
	})
	return events
}

func step1(input string) {
	events := buildEventList(input)
	guardSleepTime := map[string]int{}
	guard := ""
	from := 0
	for _, event := range events {
		fmt.Println(event)
		switch event.action {
		case beginShift:
			guard = event.guardNumber
		case fallAsleep:
			from = event.minute
		case wakeUp:
			to := event.minute
			guardSleepTime[guard] += to - from
			fmt.Println(guard, "has slept", guardSleepTime[guard], "minutes")
		}
	}
	maxSleepTime := 0
	lazyGuard := ""
	for guard, sleepTime := range guardSleepTime {
		fmt.Println("guard", guard, "slept a total of", sleepTime, "minutes")
		if sleepTime > maxSleepTime {
			maxSleepTime = sleepTime
			lazyGuard = guard
		}
	}
	fmt.Println("lazyGuard is", lazyGuard)

	minutes := [60]int{}
	var sleepiestMinute int
	for _, event := range events {
		if event.action == beginShift {
			guard = event.guardNumber
			continue
		}
		if guard != lazyGuard {
			continue
		}
		switch event.action {
		case fallAsleep:
			from = event.minute
		case wakeUp:
			to := event.minute
			for i := from; i < to; i++ {
				minutes[i]++
				if minutes[i] > minutes[sleepiestMinute] {
					sleepiestMinute = i
				}
			}
		}
	}

	fmt.Println("guard", lazyGuard, "on minute", sleepiestMinute, "step1 --->", toNumber(lazyGuard)*sleepiestMinute)
}

func step2(input string) {
	events := buildEventList(input)
	guard := ""
	lazyGuard := ""
	from := 0
	sleepiestMinute := 0
	minutes := map[string][]int{}
	guards := []string{}
	for _, event := range events {
		switch event.action {
		case beginShift:
			guard = event.guardNumber
			if lazyGuard == "" {
				lazyGuard = guard
			}
			if !find(guards, guard) {
				guards = append(guards, guard)
				minutes[guard] = make([]int, 60)
			}
		case fallAsleep:
			from = event.minute
		case wakeUp:
			for minute := from; minute < event.minute; minute++ {
				minutes[guard][minute]++
				if minutes[guard][minute] > minutes[lazyGuard][sleepiestMinute] {
					lazyGuard = guard
					sleepiestMinute = minute
				}
			}
		}
	}

	fmt.Println("guard", lazyGuard, "on minute", sleepiestMinute, "step2 --->", toNumber(lazyGuard)*sleepiestMinute)
}

func find(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func toNumber(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return number
}

func load() string {
	//text, err := ioutil.ReadFile("./input04-test.txt")
	text, err := ioutil.ReadFile("./input04.txt")
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}
