package monosubstitution

import (
	"cypher/monosubstitution"
	"fmt"
	"maps"
	"slices"

	"github.com/pinguin-frosch/menu/pkg/menu"
)

var Manual *menu.Menu
var state *monosubstitution.State

func init() {
	state = monosubstitution.NewState()
	Manual = menu.NewMenu("manual")
	Manual.AddOption("a", "add input text", func() {
		input := Manual.GetString("text: ")
		state.AddInputText(input)
	})
	Manual.AddOption("n", "get n letter frequencies", func() {
		n, err := Manual.GetInt("length: ")
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		frequencies, err := state.GetNFrecuencies(n)
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
	Manual.AddOption("r", "add letter replacement", func() {
		runes := []rune(Manual.GetString("from: "))
		if len(runes) != 1 {
			fmt.Println("invalid char")
		}
		from := runes[0]
		runes = []rune(Manual.GetString("to: "))
		if len(runes) != 1 {
			fmt.Println("invalid char")
		}
		to := runes[0]
		state.AddLetterReplacement(from, to)
	})
	Manual.AddOption("p", "print text with replacements", func() {
		text := state.GetReplacedText()
		fmt.Printf("text: %s\n", text)
	})
	Manual.AddOption("s", "show letter replacements", func() {
		replacements, keys := state.GetLetterReplacements()
		for _, key := range keys {
			fmt.Printf("%s -> %s\n", string(key), string(replacements[key]))
		}
	})
	Manual.AddOption("t", "print input text", func() {
		text := state.GetInputText()
		fmt.Printf("text: %s\n", text)
	})
}
