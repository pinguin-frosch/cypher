package vigenere_test

import (
	"cypher/polisubstitution/vigenere"
	"testing"
)

func TestCypher(t *testing.T) {
	// test valid key
	invalidKeys := []string{"", "1234", "not a valid key"}
	for _, invalidKey := range invalidKeys {
		_, err := vigenere.Cypher("message test", invalidKey)
		if err == nil {
			t.Fatalf("invalid key cannot be used: %v", invalidKey)
		}
	}

	// test that cypher keeps capitalization
	original := "did you miss ME?"
	key := "blueBook"
	expected := "etx cpi astd GI?"

	result, _ := vigenere.Cypher(original, key)
	if result != expected {
		t.Fatalf("Got: %v\n, Want: %v", result, expected)
	}

	// test with text that cannot be cyphered
	original = "ершы шы ыщ тшсу"
	key = "randomkey"
	expected = "ершы шы ыщ тшсу"

	result, _ = vigenere.Cypher(original, key)
	if result != expected {
		t.Fatalf("Got: %v\n, Want: %v", result, expected)
	}
}
