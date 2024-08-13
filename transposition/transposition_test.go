package transposition_test

import (
	"cypher/transposition"
	"strings"
	"testing"
)

func TestCipherAndDecipher(t *testing.T) {
	message := "this is an even longer message"
	message = strings.ReplaceAll(message, " ", "")
	for i := 2; i < len(message)-1; i++ {
		cipher, err := transposition.Cipher(message, i)
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		text, err := transposition.Decipher(cipher, i)
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if message != text {
			t.Errorf("Decipher(cipher) = %s; want %s, failed for jump value %d", text, message, i)
		}
	}
}
