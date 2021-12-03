package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 3, compute("1 + 2"))
	assert.Equal(t, 191, compute("100 + 91"))
}

func TestMult(t *testing.T) {
	assert.Equal(t, 12, compute("3 * 4"))
	assert.Equal(t, 9100, compute("100 * 91"))
}

func TestMultipleAdd(t *testing.T) {
	assert.Equal(t, 6, compute("1 + 2 + 3"))
	assert.Equal(t, 1119, compute("100 + 10 + 9 + 1000"))
}

func TestMultipleMultiply(t *testing.T) {
	assert.Equal(t, 30, compute("5 * 2 * 3"))
}

func TestMultipleAddAndMultiply(t *testing.T) {
	assert.Equal(t, 13, compute("5 * 2 + 3"))
	assert.Equal(t, 21, compute("5 + 2 * 3"))
}

func TestSimpleParenthesis(t *testing.T) {
	assert.Equal(t, 7, compute("1 + (2 * 3)"))
	assert.Equal(t, 11, compute("1 + (2 * 3) + 4"))
	assert.Equal(t, 40, compute("2 * (2 + 3) * 4"))
	assert.Equal(t, 10, compute("(2 + 3) * (1 + 1)"))
}

func TestSample1(t *testing.T) {
	assert.Equal(t, 71, compute("1 + 2 * 3 + 4 * 5 + 6"))
}

func TestSample2(t *testing.T) {
	assert.Equal(t, 51, compute("1 + (2 * 3) + (4 * (5 + 6))"))
}

func TestSample3(t *testing.T) {
	assert.Equal(t, 26, compute("2 * 3 + (4 * 5)"))
}

func TestSample4(t *testing.T) {
	assert.Equal(t, 437, compute("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
}

func TestSample5(t *testing.T) {
	assert.Equal(t, 12240, compute("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
}

func TestSample6(t *testing.T) {
	assert.Equal(t, 13632, compute("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
}

//1 + 2 * 3 + 4 * 5 + 6

func TestS2Simple(t *testing.T) {
	assert.Equal(t, 3, compute2("1 + 2"))
}

func TestS2Sample1(t *testing.T) {
	assert.Equal(t, 231, compute2("1 + 2 * 3 + 4 * 5 + 6"))
}

func TestS2Sample2(t *testing.T) {
	assert.Equal(t, 51, compute2("1 + (2 * 3) + (4 * (5 + 6))"))
}

func TestS2Sample3(t *testing.T) {
	assert.Equal(t, 46, compute2("2 * 3 + (4 * 5)"))
}

func TestS2Sample4(t *testing.T) {
	assert.Equal(t, 1445, compute2("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
}

func TestS2Sample5(t *testing.T) {
	assert.Equal(t, 669060, compute2("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
}

func TestS2Sample6(t *testing.T) {
	assert.Equal(t, 23340, compute2("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
}
