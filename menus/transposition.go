package menus

import (
	"cypher/menu"
	"cypher/transposition"
	"fmt"
	"strconv"
)

var TranspositionMenu *menu.Menu

func init() {
	TranspositionMenu = menu.NewMenu("transposition")
	TranspositionMenu.AddOption("b", "brute force decipher by columns", func() {
		scanner := TranspositionMenu.Scanner
		fmt.Print("message: ")
		scanner.Scan()
		transposition.BruteForceDecipher(scanner.Text())
	})
	TranspositionMenu.AddOption("c", "cipher by columns", func() {
		scanner := TranspositionMenu.Scanner
		fmt.Print("text: ")
		scanner.Scan()
		input := scanner.Text()
		fmt.Print("columns: ")
		scanner.Scan()
		columns := scanner.Text()
		columnSize, err := strconv.Atoi(columns)
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
		scanner := TranspositionMenu.Scanner
		fmt.Print("cipher: ")
		scanner.Scan()
		input := scanner.Text()
		fmt.Print("columns: ")
		scanner.Scan()
		columns := scanner.Text()
		columnSize, err := strconv.Atoi(columns)
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
