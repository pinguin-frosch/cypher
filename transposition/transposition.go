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
	jumpSize := (len(cipher) / groupSize) + 1
	text := ""
	for i := 0; i < jumpSize; i++ {
		for j := i; j < len(cipher); j += jumpSize {
			text += string(cipher[j])
		}
	}
	return text, nil
}
