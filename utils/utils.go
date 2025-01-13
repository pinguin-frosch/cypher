package utils

import "errors"

// Checks whether a letter is a latin letter or not
func IsLatinLetter(letter rune) bool {
	if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
		return true
	}
	return false
}

// Checks whether a letter is an uppercase latin letter
func IsUppercaseLatinLetter(letter rune) bool {
	if letter >= 65 && letter <= 90 {
		return true
	}
	return false
}

// Checks whether a letter is a lowercase latin letter
func IsLowercaseLatinLetter(letter rune) bool {
	if letter >= 97 && letter <= 122 {
		return true
	}
	return false
}

// Counts how many times each substring appears in the input text
func AnalyzeSubstringFrequencies(input string, length int) (map[string]int, error) {
	if length <= 0 {
		return nil, errors.New("invalid length: use at least 1")
	}
	frequencies := make(map[string]int)
	for i := 0; i < len(input)-(length-1); i++ {
		substring := input[i : i+length]
		if _, ok := frequencies[substring]; !ok {
			frequencies[substring] = 0
		}
		frequencies[substring] += 1
	}
	return frequencies, nil
}
