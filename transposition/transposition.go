package transposition

import (
	"errors"
)

func CipherByColumns(text string, groupSize int) (string, error) {
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
		cipher += " "
	}
	return cipher, nil
}
