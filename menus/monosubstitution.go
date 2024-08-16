package menus

import (
	"cypher/menu"
	"cypher/monosubstitution"
	"fmt"
)

var MonoSubstitutionMenu *menu.Menu
var monoSubstitutionState *monosubstitution.State

func init() {
	monoSubstitutionState = monosubstitution.NewState()
	MonoSubstitutionMenu = menu.NewMenu("monosubstitution")
	MonoSubstitutionMenu.AddOption("a", "add input text", func() {
		scanner := TranspositionMenu.Scanner
		fmt.Print("text: ")
		scanner.Scan()
		monoSubstitutionState.AddInputText(scanner.Text())
	})
	MonoSubstitutionMenu.AddOption("f", "show letter frequencies", func() {
		frequencies := monoSubstitutionState.GetLetterFrequencies()
		if len(frequencies) == 0 {
			fmt.Println("no frequencies yet")
			return
		}
		for c, freq := range frequencies {
			fmt.Printf("%c: %.2f%%\n", c, freq.Percentage)
		}
	})
}
