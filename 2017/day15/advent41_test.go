package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert.Equal(t, 1, advent15(65, 8921, 5))
}

func TestExample2With5Miteration(t *testing.T) {
	assert.Equal(t, 309, advent15step2(65, 8921, 5*million))
}

func TestExampleWith40Miteration(t *testing.T) {
	assert.Equal(t, 588, advent15(65, 8921, 40*million))
}

func TestStep1(t *testing.T) {
	assert.Equal(t, 600, advent15(699, 124, 40*million))
}

func TestStep2(t *testing.T) {
	assert.Equal(t, 313, advent15step2(699, 124, 5*million))
}

func TestGeneratorA(t *testing.T) {
	generator := newGeneratorA(65)
	assert.Equal(t, int64(1092455), generator.next())
	assert.Equal(t, int64(1181022009), generator.next())
	assert.Equal(t, int64(245556042), generator.next())
	assert.Equal(t, int64(1744312007), generator.next())
	assert.Equal(t, int64(1352636452), generator.next())
}

func TestGeneratorABinary(t *testing.T) {
	generator := newGeneratorA(65)
	assert.Equal(t, "100001010101101100111", generator.nextBinary())
	assert.Equal(t, "1000110011001001111011100111001", generator.nextBinary())
	assert.Equal(t, "1110101000101110001101001010", generator.nextBinary())
	assert.Equal(t, "1100111111110000001011011000111", generator.nextBinary())
	assert.Equal(t, "1010000100111111001100000100100", generator.nextBinary())
}

func TestGeneratorWithShiftBits(t *testing.T) {
	generator := newGeneratorA(65)
	assert.Equal(t, "1010101101100111", toBinaryString(generator.nextRightMost16Bits()))
	assert.Equal(t, "1111011100111001", toBinaryString(generator.nextRightMost16Bits()))
	assert.Equal(t, "1110001101001010", toBinaryString(generator.nextRightMost16Bits()))
	assert.Equal(t, "1011011000111", toBinaryString(generator.nextRightMost16Bits()))
	assert.Equal(t, "1001100000100100", toBinaryString(generator.nextRightMost16Bits()))
}

func TestGeneratorB(t *testing.T) {
	generator := newGeneratorB(8921)
	assert.Equal(t, int64(430625591), generator.next())
	assert.Equal(t, int64(1233683848), generator.next())
	assert.Equal(t, int64(1431495498), generator.next())
	assert.Equal(t, int64(137874439), generator.next())
	assert.Equal(t, int64(285222916), generator.next())
}
