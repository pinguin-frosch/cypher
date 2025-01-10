package transposition

import (
	"cypher/transposition"
	"fmt"

	"github.com/pinguin-frosch/menu/pkg/menu"
)

var Main *menu.Menu

func init() {
	Main = menu.NewMenu("transposition")
	Main.AddOption("cb", "brute force decipher by columns", func() {
		message := Main.GetString("message: ")
		transposition.BruteForceDecipher(message)
	})
	Main.AddOption("cc", "cipher by columns", func() {
		input := Main.GetString("text: ")
		columnSize, err := Main.GetInt("columns: ")
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
	Main.AddOption("cd", "decipher by columns", func() {
		input := Main.GetString("text: ")
		columnSize, err := Main.GetInt("columns: ")
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
