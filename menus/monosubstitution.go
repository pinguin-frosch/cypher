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
	MonoSubstitutionMenu.AddOption("r", "add letter replacement", func() {
		scanner := TranspositionMenu.Scanner
		fmt.Print("from: ")
		scanner.Scan()
		runes := []rune(scanner.Text())
		if len(runes) != 1 {
			fmt.Println("invalid char")
		}
		from := runes[0]
		fmt.Print("to: ")
		scanner.Scan()
		runes = []rune(scanner.Text())
		if len(runes) != 1 {
			fmt.Println("invalid char")
		}
		to := runes[0]
		monoSubstitutionState.AddLetterReplacement(from, to)
	})
	MonoSubstitutionMenu.AddOption("p", "print text with replacements", func() {
		text := monoSubstitutionState.GetReplacedText()
		fmt.Printf("text: %s\n", text)
	})
	MonoSubstitutionMenu.AddOption("s", "show letter replacements", func() {
		replacements, keys := monoSubstitutionState.GetLetterReplacements()
		for _, key := range keys {
			fmt.Printf("%s -> %s\n", string(key), string(replacements[key]))
		}
	})
	MonoSubstitutionMenu.AddOption("t", "print input text", func() {
		text := monoSubstitutionState.GetInputText()
		fmt.Printf("text: %s\n", text)
	})
}
