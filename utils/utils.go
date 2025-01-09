package utils

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
