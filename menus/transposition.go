package menus

import (
	"cypher/transposition"
	"fmt"

	"github.com/pinguin-frosch/menu/pkg/menu"
)

var TranspositionMenu *menu.Menu

func init() {
	TranspositionMenu = menu.NewMenu("transposition")
	TranspositionMenu.AddOption("b", "brute force decipher by columns", func() {
		message := TranspositionMenu.GetString("message: ")
		transposition.BruteForceDecipher(message)
	})
	TranspositionMenu.AddOption("c", "cipher by columns", func() {
		input := TranspositionMenu.GetString("text: ")
		columnSize, err := TranspositionMenu.GetInt("columns: ")
		if err != nil {
			fmt.Println("error: invalid column size")
			return
		}
		cipher, err := transposition.Cipher(input, columnSize)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		fmt.Printf("cipher: %s\n", cipher)
	})
	TranspositionMenu.AddOption("d", "decipher by columns", func() {
		input := TranspositionMenu.GetString("text: ")
		columnSize, err := TranspositionMenu.GetInt("columns: ")
		if err != nil {
			fmt.Println("error: invalid column size")
			return
		}
		text, err := transposition.Decipher(input, columnSize)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		fmt.Printf("text: %s\n", text)
	})
}
