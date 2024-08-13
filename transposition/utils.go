package transposition

import (
	"fmt"
	"strings"
)

func BruteForceDecipher(message string) {
	message = strings.ReplaceAll(message, " ", "")
	if len(message) < 3 {
		return
	}
	for i := 2; i < len(message)-1; i++ {
		text, err := Decipher(message, i)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		fmt.Printf("%d %s\n", i, text)
	}
}
