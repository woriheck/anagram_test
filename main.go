package main

import (
	"reflect"
)

func isAnagram(A string, B string) bool {
	// word length not same, so is not anagram
	if len(A) != len(B) {
		return false
	}

	// Verify two map data is same
	return reflect.DeepEqual(
		mapLetterFreq(A),
		mapLetterFreq(B),
	)
}

// Map string char frequency
// Example:
//   Given: AABC
//   Return: map[A:2 B:1 C:1]
func mapLetterFreq(s string) map[string]int {
	result := map[string]int{}

	for _, letter := range s {
		if _, ok := result[string(letter)]; ok {
			result[string(letter)]++
		} else {
			result[string(letter)] = 1
		}
	}
	return result
}
