package monosubstitution

import (
	"cypher/monosubstitution/rot"
	"fmt"

	"github.com/pinguin-frosch/menu/pkg/menu"
)

var Main *menu.Menu

func init() {
	Main = menu.NewMenu("monosubstitution")
	Main.AddOption("m", "manual substitution", func() {
		Manual.Start()
	})
	Main.AddOption("rc", "cipher using rotary", func() {
		text := Main.GetString("text: ")
		amount, err := Main.GetInt("amount: ")
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}
		cipher, err := rot.Cipher(text, amount)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}
		fmt.Printf("cipher: %s\n", cipher)
	})
	Main.AddOption("rd", "decipher using rotary", func() {
		cipher := Main.GetString("cipher: ")
		amount, err := Main.GetInt("amount: ")
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}
		text, err := rot.Decipher(cipher, amount)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}
		fmt.Printf("text: %s\n", text)
	})
	Main.AddOption("rb", "brute force using rotary", func() {
		cipher := Main.GetString("cipher: ")
		results, err := rot.BruteForce(cipher)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		for i, result := range results {
			fmt.Printf("rot-%d: %s\n", i+1, result)
		}
	})
}
