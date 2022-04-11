package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	dataset := [][]string{
		{"AABC", "CBSAA"},
		{"attack", "tattoo"},
		{"AABCD", "DAABC"},
		{"cat", "act"},
		{"save", "vase"},
	}

	expected := []bool{
		false,
		false,
		true,
		true,
		true,
	}

	for index, input := range dataset {
		path := isAnagram(input[0], input[1])
		assert.Equal(t, path, expected[index], "Not correct.")
	}
}
