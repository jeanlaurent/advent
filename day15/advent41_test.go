package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestItRetursFalseWhenNoArg(t *testing.T) {
// 	assert.Equal(t, 588, advent15(65, 8921))
// }

func TestGeneratorA(t *testing.T) {
	generator := Generator{factor: generatorAfactor, previous: 65}
	assert.Equal(t, 1092455, generator.next())
	assert.Equal(t, 1181022009, generator.next())
	assert.Equal(t, 245556042, generator.next())
	assert.Equal(t, 1744312007, generator.next())
	assert.Equal(t, 1352636452, generator.next())
}

func TestGeneratorB(t *testing.T) {
	generator := Generator{factor: generatorBfactor, previous: 8921}
	assert.Equal(t, 430625591, generator.next())
	assert.Equal(t, 1233683848, generator.next())
	assert.Equal(t, 1431495498, generator.next())
	assert.Equal(t, 137874439, generator.next())
	assert.Equal(t, 285222916, generator.next())
}

const divisor = 2147483647
const generatorAfactor = 16807
const generatorBfactor = 48271

type Generator struct {
	factor   int
	previous int
}

func (g *Generator) next() int {
	g.previous = g.previous * g.factor % divisor
	return g.previous
}

func advent15(generatorAPrevious int, generatorBPrevious int) int {
	return 0
}
