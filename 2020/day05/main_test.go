package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleUseCase(t *testing.T) {
	ticket := "FBFBBFFRLR"
	assert.Equal(t, findRow(ticket), 44)
	assert.Equal(t, findColumn(ticket), 5)
	seatID, _, _ := findSeatId(ticket)
	assert.Equal(t, seatID, 357)
}

func TestExample1(t *testing.T) {
	ticket := "BFFFBBFRRR"
	assert.Equal(t, findRow(ticket), 70)
	assert.Equal(t, findColumn(ticket), 7)
	seatID, _, _ := findSeatId(ticket)
	assert.Equal(t, seatID, 567)
}

func TestExample2(t *testing.T) {
	ticket := "FFFBBBFRRR"
	assert.Equal(t, findRow(ticket), 14)
	assert.Equal(t, findColumn(ticket), 7)
	seatID, _, _ := findSeatId(ticket)
	assert.Equal(t, seatID, 119)
}

func TestExample3(t *testing.T) {
	ticket := "BBFFBBFRLL"
	assert.Equal(t, findRow(ticket), 102)
	assert.Equal(t, findColumn(ticket), 4)
	seatID, _, _ := findSeatId(ticket)
	assert.Equal(t, seatID, 820)
}
