package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirth(t *testing.T) {
	assert.True(t, isValidBirthDate("2002"))
	assert.False(t, isValidBirthDate("2003"))
	assert.False(t, isValidBirthDate(""))
}

func TestHeight(t *testing.T) {
	assert.True(t, isValidHeight("60in"))
	assert.True(t, isValidHeight("190cm"))
	assert.False(t, isValidHeight("149cm"))
	assert.False(t, isValidHeight("194cm"))
	assert.False(t, isValidHeight("190in"))
	assert.False(t, isValidHeight("55in"))
	assert.False(t, isValidHeight("190"))
	assert.False(t, isValidHeight(""))
}

func TestHcl(t *testing.T) {
	assert.True(t, isValidHairColor("#123abc"))
	assert.False(t, isValidHairColor("#123abz"))
	assert.False(t, isValidHairColor("123abc"))
}

func TestEcl(t *testing.T) {
	assert.True(t, isValidEyeColor("brn"))
	assert.False(t, isValidHairColor("wat"))
}

func TestPid(t *testing.T) {
	assert.True(t, isValidPid("000000001"))
	assert.False(t, isValidPid("0123456789"))
}
