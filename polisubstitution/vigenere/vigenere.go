package vigenere

import (
	"errors"
	"log"
	"strings"
)

// Cyphers a text using the vigenere cypher with the specified key,
// only latin caracters are processed, the rest are ignored.
func Cypher(text string, key string) (string, error) {
	// key has to be at least 1 characters long and valid key
	err := isValidKey(key)
	if err != nil {
		return "", err
	}

	cypher := ""

	// Capitalize to remain consistent
	key = strings.ToUpper(key)

	// Start on the first index of the key
	keyIndex := 0

	for _, char := range text {
		// Keep non latin letters untouched
		if !IsLatinLetter(char) {
			cypher += string(char)
			continue
		}

		// Use the corresponding char based on the index
		keyChar := key[keyIndex]

		// Cypher current letter
		cypher += string(cypherLetter(char, rune(keyChar)))

		// Advance key
		keyIndex = (keyIndex + 1) % len(key)
	}

	return cypher, nil
}

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

// Cyphers a single letter using the given key char, it is assumed that the
// letters are latin letters and the key is in uppercase
func cypherLetter(letter rune, keyChar rune) rune {
	letterOffset := rune(0)
	if IsLowercaseLatinLetter(letter) {
		letterOffset = 97
	} else if IsUppercaseLatinLetter(letter) {
		letterOffset = 65
	} else {
		log.Fatalf("invalid letter: %v", letter)
	}
	return (((letter - letterOffset) + (keyChar - 65)) % 26) + letterOffset
}

// A valid key only has latin letters, no symbols, numbers or anything else
func isValidKey(key string) error {
	if len(key) == 0 {
		return errors.New("key cannot be empty")
	}
	for _, char := range key {
		if !IsUppercaseLatinLetter(char) && !IsLowercaseLatinLetter(char) {
			return errors.New("key contains invalid characters")
		}
	}
	return nil
}

// Decyphers a text using the vigenere cypher with the specified key,
// only latin caracters are processed, the rest are ignored.
func Decypher(cypher string, key string) (string, error) {
	// key has to be at least 1 characters long and valid key
	err := isValidKey(key)
	if err != nil {
		return "", err
	}

	text := ""

	// Capitalize to remain consistent
	key = strings.ToUpper(key)

	// Start on the first index of the key
	keyIndex := 0

	for _, char := range cypher {
		// Keep non latin letters untouched
		if !IsLatinLetter(char) {
			text += string(char)
			continue
		}

		// Use the corresponding char based on the index
		keyChar := key[keyIndex]

		// Cypher current letter
		text += string(decypherLetter(char, rune(keyChar)))

		// Advance key
		keyIndex = (keyIndex + 1) % len(key)
	}

	return text, nil
}

// Decyphers a single letter using the given key char, it is assumed that the
// letters are latin letters and the key is in uppercase
func decypherLetter(letter rune, keyChar rune) rune {
	letterOffset := rune(0)
	if IsLowercaseLatinLetter(letter) {
		letterOffset = 97
	} else if IsUppercaseLatinLetter(letter) {
		letterOffset = 65
	} else {
		log.Fatalf("invalid letter: %v", letter)
	}
	adjustedLetter := (letter - letterOffset) - (keyChar - 65)
	if adjustedLetter < 0 {
		adjustedLetter += 26
	}
	return (adjustedLetter % 26) + letterOffset
}
