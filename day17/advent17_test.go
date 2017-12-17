package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialState(t *testing.T) {
	spinlock := newDefaultSpinLock(3)
	assert.Equal(t, []int{0}, spinlock.buffer)
}

func TestStep3(t *testing.T) {
	spinlock := newDefaultSpinLock(3)
	spinlock.run()
	assert.Equal(t, []int{0, 1}, spinlock.buffer)
	assert.Equal(t, 1, spinlock.position)
	spinlock.run()
	assert.Equal(t, []int{0, 2, 1}, spinlock.buffer)
	assert.Equal(t, 1, spinlock.position)
	spinlock.run()
	assert.Equal(t, []int{0, 2, 3, 1}, spinlock.buffer)
	assert.Equal(t, 2, spinlock.position)
	spinlock.run()
	assert.Equal(t, []int{0, 2, 4, 3, 1}, spinlock.buffer)
	assert.Equal(t, 2, spinlock.position)
	spinlock.run()
	assert.Equal(t, []int{0, 5, 2, 4, 3, 1}, spinlock.buffer)
	assert.Equal(t, 1, spinlock.position)
}

func TestPart1(t *testing.T) {
	spinlock := newDefaultSpinLock(343)
	for index := 0; index < 2017; index++ {
		spinlock.run()
	}
	result := -1
	for index, value := range spinlock.buffer {
		if value == 2017 {
			result = spinlock.buffer[(index+1)%len(spinlock.buffer)]
		}
	}
	assert.Equal(t, 1914, result)
}

func TestPart2(t *testing.T) {
	last := -1
	index := 0
	for i := 1; i <= 50000000; i++ {
		index = (index+343)%i + 1
		if index == 1 {
			last = i
		}
	}
	assert.Equal(t, 41797835, last)
}

type Spinlock struct {
	buffer   []int
	position int
	steps    int
	value    int
}

func (s *Spinlock) run() {
	// fmt.Println(s.buffer)
	s.position = (s.position+s.steps)%len(s.buffer) + 1
	// fmt.Println("new insertion position is ", s.position)
	// fmt.Println("value is now ", s.value)
	// LOL
	s.buffer = append(s.buffer, 0)
	copy(s.buffer[s.position+1:], s.buffer[s.position:])
	s.buffer[s.position] = s.value
	// fmt.Println(s.buffer)
	s.value++
	// fmt.Println()
}

func newDefaultSpinLock(steps int) Spinlock {
	buffer := []int{0}
	return Spinlock{buffer: buffer, position: 0, steps: steps, value: 1}
}
