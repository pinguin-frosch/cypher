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

var ManualSubstitutionMenu *menu.Menu
var ManualSubstitutionState *monosubstitution.State

func init() {
	MonoSubstitutionMenu = menu.NewMenu("monosubstitution")
	MonoSubstitutionMenu.AddOption("m", "manual substitution", func() {
		ManualSubstitutionMenu.Start()
	})

	// TODO: move this menu definition somewhere else
	ManualSubstitutionState = monosubstitution.NewState()
	ManualSubstitutionMenu = menu.NewMenu("manual")
	ManualSubstitutionMenu.AddOption("a", "add input text", func() {
		input := ManualSubstitutionMenu.GetString("text: ")
		ManualSubstitutionState.AddInputText(input)
	})
	ManualSubstitutionMenu.AddOption("n", "get n letter frequencies", func() {
		n, err := ManualSubstitutionMenu.GetInt("length: ")
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		frequencies, err := ManualSubstitutionState.GetNFrecuencies(n)
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
	ManualSubstitutionMenu.AddOption("r", "add letter replacement", func() {
		runes := []rune(ManualSubstitutionMenu.GetString("from: "))
		if len(runes) != 1 {
			fmt.Println("invalid char")
		}
		from := runes[0]
		runes = []rune(ManualSubstitutionMenu.GetString("to: "))
		if len(runes) != 1 {
			fmt.Println("invalid char")
		}
		to := runes[0]
		ManualSubstitutionState.AddLetterReplacement(from, to)
	})
	ManualSubstitutionMenu.AddOption("p", "print text with replacements", func() {
		text := ManualSubstitutionState.GetReplacedText()
		fmt.Printf("text: %s\n", text)
	})
	ManualSubstitutionMenu.AddOption("s", "show letter replacements", func() {
		replacements, keys := ManualSubstitutionState.GetLetterReplacements()
		for _, key := range keys {
			fmt.Printf("%s -> %s\n", string(key), string(replacements[key]))
		}
	})
	ManualSubstitutionMenu.AddOption("t", "print input text", func() {
		text := ManualSubstitutionState.GetInputText()
		fmt.Printf("text: %s\n", text)
	})
}
