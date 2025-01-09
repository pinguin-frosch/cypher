package rot_test

import (
	"cypher/monosubstitution/rot"
	"testing"
)

func TestCipher(t *testing.T) {
	// use invalid amounts
	invalidAmounts := []int{-2, 38}
	for _, invalidAmount := range invalidAmounts {
		_, err := rot.Cipher("message test", invalidAmount)
		if err == nil {
			t.Fatalf("invalid amount cannot be used: %v", invalidAmount)
		}
	}

	// test using rot-7
	original := "Sample text, nothing important."
	amount := 7
	expected := "Zhtwsl alea, uvaopun ptwvyahua."

	result, _ := rot.Cipher(original, amount)
	if result != expected {
		t.Fatalf("Got: %v\n, Want: %v", result, expected)
	}
}

func TestCipherAndDecipher(t *testing.T) {
	// test using rot-5, followed by rot-21
	text := "random text to execute this test."
	amount := 5
	cipher, _ := rot.Cipher(text, amount)
	result, _ := rot.Decipher(cipher, amount)
	if text != result {
		t.Fatalf("rot-%v followed by rot-%v modifies the text, got: %v, wanted: %v", amount, 26-amount, result, text)
	}
}
