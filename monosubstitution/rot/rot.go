package rot

import (
	"cypher/utils"
	"errors"
	"log"
)

// Ciphers a text by rotating it by the given amount clockwise, non latin
// letters are ignored
func Cipher(text string, amount int) (string, error) {
	if amount < 0 || amount > 26 {
		return "", errors.New("invalid rot amount")
	}
	cipher := ""

	for _, char := range text {
		// Keep non latin letters untouched
		if !utils.IsLatinLetter(char) {
			cipher += string(char)
			continue
		}
		// Cypher current letter
		cipher += string(rotateLetter(char, amount))
	}

	return cipher, nil
}

// Deciphers a text by rotating it by the given amount counter clockwise, non
// latin letters are ignored
func Decipher(cipher string, amount int) (string, error) {
	if amount < 0 || amount > 26 {
		return "", errors.New("invalid rot amount")
	}
	text := ""

	for _, char := range cipher {
		// Keep non latin letters untouched
		if !utils.IsLatinLetter(char) {
			text += string(char)
			continue
		}
		// Cypher current letter
		text += string(rotateLetter(char, -amount))
	}

	return text, nil
}

// Tries all possible rotations and returns all of them
// latin letters are ignored
func BruteForce(text string) ([]string, error) {
	results := make([]string, 0, 25)
	for i := 1; i < 26; i++ {
		result, err := Decipher(text, i)
		if err != nil {
			return nil, nil
		}
		results = append(results, result)
	}
	return results, nil
}

// Rotates a single letter by the given amount, amount can be a negative value
func rotateLetter(letter rune, amount int) rune {
	letterOffset := 0
	if utils.IsLowercaseLatinLetter(letter) {
		letterOffset = 97
	} else if utils.IsUppercaseLatinLetter(letter) {
		letterOffset = 65
	} else {
		log.Fatalf("invalid letter: %v", letter)
	}
	l := int(letter)
	adjustedLetter := (l - letterOffset) + amount
	for adjustedLetter < 0 {
		adjustedLetter += 26
	}
	return (rune(adjustedLetter) % 26) + rune(letterOffset)
}
