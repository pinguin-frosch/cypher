package transposition

import (
	"errors"
	"strings"
)

func Cipher(text string, groupSize int) (string, error) {
	text = strings.ReplaceAll(text, " ", "")
	if groupSize < 2 {
		return "", errors.New("groupSize needs to be greater than 1")
	}
	if groupSize > len(text) {
		return "", errors.New("groupSize needs to be less than the length of the text")
	}
	cipher := ""
	for i := 0; i < groupSize; i++ {
		for j := i; j < len(text); j += groupSize {
			cipher += string(text[j])
		}
	}
	return cipher, nil
}

func Decipher(cipher string, groupSize int) (string, error) {
	cipher = strings.ReplaceAll(cipher, " ", "")
	if groupSize < 2 {
		return "", errors.New("groupSize needs to be greater than 1")
	}
	if groupSize > len(cipher) {
		return "", errors.New("groupSize needs to be less than the length of the text")
	}

	// Determine the max jump size between letters
	jumpSize := len(cipher) / groupSize
	perfectGroups := len(cipher) % groupSize
	if perfectGroups > 0 {
		jumpSize++
	}

	// Set limit values per substring
	limitValues := make([]int, groupSize+1)
	limitValues[0] = 0
	for i := 1; i <= perfectGroups; i++ {
		limitValues[i] = jumpSize
	}
	for i := perfectGroups + 1; i <= groupSize; i++ {
		if perfectGroups == 0 {
			limitValues[i] = jumpSize
		} else {
			limitValues[i] = jumpSize - 1
		}
	}

	// Adjust limit values per substring to index
	for i := 1; i < groupSize+1; i++ {
		limitValues[i] = limitValues[i] + limitValues[i-1]
	}

	// Retrieve substrings
	substrings := make([]string, groupSize)
	for i := 0; i < len(limitValues)-1; i++ {
		start := limitValues[i]
		end := limitValues[i+1]
		substrings[i] = cipher[start:end]
	}

	// Setup matrix to save substring in the right format
	matrix := make([][]string, groupSize)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]string, jumpSize)
	}

	// Arrange letters in the matrix format
	for i := 0; i < groupSize; i++ {
		for j, letter := range substrings[i] {
			matrix[i][j] = string(letter)
		}
	}

	// Reconstruct original message
	text := ""
	for i := 0; i < jumpSize; i++ {
		for j := 0; j < groupSize; j++ {
			letter := matrix[j][i]
			if letter == "" {
				break
			}
			text += letter
		}
	}

	return text, nil
}
