package menus

import (
	"cypher/monosubstitution"
	"fmt"
	"maps"
	"slices"

	"github.com/pinguin-frosch/menu/pkg/menu"
)

var MonoSubstitutionMenu *menu.Menu
var monoSubstitutionState *monosubstitution.State

func init() {
	monoSubstitutionState = monosubstitution.NewState()
	MonoSubstitutionMenu = menu.NewMenu("monosubstitution")
	MonoSubstitutionMenu.AddOption("a", "add input text", func() {
		input := MonoSubstitutionMenu.GetString("text: ")
		monoSubstitutionState.AddInputText(input)
	})
	MonoSubstitutionMenu.AddOption("n", "get n letter frequencies", func() {
		n, err := MonoSubstitutionMenu.GetInt("length: ")
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		frequencies, err := monoSubstitutionState.GetNFrecuencies(n)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		substrings := slices.Collect(maps.Keys(frequencies))
		slices.SortFunc(substrings, func(a, b string) int {
			if frequencies[a].Times < frequencies[b].Times {
				return 1
			} else if frequencies[a].Times > frequencies[b].Times {
				return -1
			}
			if a < b {
				return -1
			} else if b > a {
				return 1
			}
			return 0
		})
		for _, substring := range substrings {
			f := frequencies[substring]
			if f.Times == 1 {
				continue
			}
			fmt.Printf("%s\t%d\t%.2f%%\n", substring, f.Times, f.Percentage*100)
		}
	})
	MonoSubstitutionMenu.AddOption("r", "add letter replacement", func() {
		runes := []rune(MonoSubstitutionMenu.GetString("from: "))
		if len(runes) != 1 {
			fmt.Println("invalid char")
		}
		from := runes[0]
		runes = []rune(MonoSubstitutionMenu.GetString("to: "))
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
