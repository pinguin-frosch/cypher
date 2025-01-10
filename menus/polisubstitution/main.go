package polisubstitution

import (
	"cypher/polisubstitution/vigenere"
	"fmt"

	"github.com/pinguin-frosch/menu/pkg/menu"
)

var Main *menu.Menu

func init() {
	Main = menu.NewMenu("polisubstitution")
	Main.AddOption("vc", "cipher using vigenere", func() {
		text := Main.GetString("text: ")
		key := Main.GetString("key: ")
		cipher, err := vigenere.Cypher(text, key)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		fmt.Printf("cipher: %s\n", cipher)
	})
	Main.AddOption("vd", "decipher using vigenere", func() {
		cipher := Main.GetString("cipher: ")
		key := Main.GetString("key: ")
		text, err := vigenere.Decypher(cipher, key)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		fmt.Printf("text: %s\n", text)
	})
}
