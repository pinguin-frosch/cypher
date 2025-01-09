package menus

import (
	"cypher/polisubstitution/vigenere"
	"fmt"

	"github.com/pinguin-frosch/menu/pkg/menu"
)

var PolisubstitutionMenu *menu.Menu

func init() {
	PolisubstitutionMenu = menu.NewMenu("polisubstitution")
	PolisubstitutionMenu.AddOption("vc", "cipher using vigenere", func() {
		text := PolisubstitutionMenu.GetString("text: ")
		key := PolisubstitutionMenu.GetString("key: ")
		cipher, err := vigenere.Cypher(text, key)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		fmt.Printf("cipher: %s\n", cipher)
	})
	PolisubstitutionMenu.AddOption("vd", "decipher using vigenere", func() {
		cipher := PolisubstitutionMenu.GetString("cipher: ")
		key := PolisubstitutionMenu.GetString("key: ")
		text, err := vigenere.Decypher(cipher, key)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		fmt.Printf("text: %s\n", text)
	})
}
