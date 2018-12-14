package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var scores numbers

func TestMain(m *testing.M) {
	scores = numbers{[]byte{3, 7}, 0, 1}
	scores.compute()
	os.Exit(m.Run())
}

func TestStep1For5(t *testing.T) {
	assert.Equal(t, "0124515891", scores.findForStep1(5))
}

func TestStep1For9(t *testing.T) {
	assert.Equal(t, "5158916779", scores.findForStep1(9))
}

func TestStep1For18(t *testing.T) {
	assert.Equal(t, "9251071085", scores.findForStep1(18))
}

func TestStep1For2018(t *testing.T) {
	assert.Equal(t, "5941429882", scores.findForStep1(2018))
}

func TestStep1(t *testing.T) {
	assert.Equal(t, "3718110721", scores.findForStep1(306281))
}

func TestStep2For51589(t *testing.T) {
	assert.Equal(t, 9, scores.findForStep2("51589"))
}

func TestStep2For01245(t *testing.T) {
	assert.Equal(t, 5, scores.findForStep2("01245"))
}

func TestStep2For92510(t *testing.T) {
	assert.Equal(t, 18, scores.findForStep2("92510"))
}

func TestStep2For59414(t *testing.T) {
	assert.Equal(t, 2018, scores.findForStep2("59414"))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, 20298300, scores.findForStep2("306281"))
}

type numbers struct {
	array     []byte
	elf1index int64
	elf2index int64
}

func (n *numbers) compute() {
	fmt.Println("computing numbers")
	for len(n.array) < 25000000 {
		n.turn()
	}
}

func (n *numbers) turn() {
	sum := n.array[n.elf1index] + n.array[n.elf2index]
	if sum >= 10 {
		n.array = append(n.array, 1)
		n.array = append(n.array, sum-10)
	} else {
		n.array = append(n.array, sum)
	}
	n.elf1index = (n.elf1index + int64(n.array[n.elf1index]) + 1) % int64(len(n.array))
	n.elf2index = (n.elf2index + int64(n.array[n.elf2index]) + 1) % int64(len(n.array))
}

func (n *numbers) findForStep2(goal string) int {
	goalArray := []byte{}
	for _, char := range goal {
		goalArray = append(goalArray, byte(char-'0'))
	}
	return bytes.Index(n.array, goalArray)
}

func (n *numbers) findForStep1(goal int) string {
	var sb strings.Builder
	for index := goal; index < goal+10; index++ {
		sb.WriteString(strconv.Itoa(int(n.array[index])))
	}
	return sb.String()
}
