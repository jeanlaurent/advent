package main

import "strconv"

const divisor = 2147483647
const generatorAfactor = 16807
const generatorBfactor = 48271
const million = 1000000

// Generator generates number
type Generator struct {
	factor   int64
	previous int64
}

func (g *Generator) next() int64 {
	g.previous = (g.previous * g.factor) % divisor
	return g.previous
}

func (g *Generator) nextBinary() string {
	return toBinaryString(g.next())
}

func (g *Generator) nextRightMost16Bits() int64 {
	return rightmost16bits(g.next())
}

func (g *Generator) nextMultipleRightMost16Bits(multiple int64) int64 {
	for {
		value := g.next()
		if value%multiple == 0 {
			return rightmost16bits(value)
		}
	}
}

func newGeneratorA(previousValue int64) Generator {
	return Generator{factor: generatorAfactor, previous: previousValue}
}

func newGeneratorB(previousValue int64) Generator {
	return Generator{factor: generatorBfactor, previous: previousValue}
}

func advent15(generatorAPrevious int64, generatorBPrevious int64, iteration int64) int {
	generatorA := newGeneratorA(generatorAPrevious)
	generatorB := newGeneratorB(generatorBPrevious)
	count := 0
	for index := int64(0); index < iteration; index++ {
		if generatorA.nextRightMost16Bits() == generatorB.nextRightMost16Bits() {
			count++
		}
	}
	return count
}

func advent15step2(generatorAPrevious int64, generatorBPrevious int64, iteration int64) int {
	generatorA := newGeneratorA(generatorAPrevious)
	generatorB := newGeneratorB(generatorBPrevious)
	count := 0
	for index := int64(0); index < iteration; index++ {
		if generatorA.nextMultipleRightMost16Bits(4) == generatorB.nextMultipleRightMost16Bits(8) {
			count++
		}
	}
	return count
}

func rightmost16bits(value int64) int64 {
	// return value & ((1 << 16) - 1)
	return value & 0xFFFF
}

func toBinaryString(value int64) string {
	return strconv.FormatInt(value, 2)
}
