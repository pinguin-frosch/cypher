package polisubstitution

import (
	"cmp"
	"cypher/polisubstitution/vigenere"
	"cypher/utils"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/pinguin-frosch/menu/pkg/menu"
)

var analysis *vigenere.Analysis
var VigenereAnalysis *menu.Menu

func init() {
	VigenereAnalysis = menu.NewMenu("analysis")
	analysis = vigenere.NewAnalysis()

	VigenereAnalysis.AddOption("ia", "add input text", func() {
		input := VigenereAnalysis.GetString("input: ")
		analysis.AddInput(input)
	})
	VigenereAnalysis.AddOption("ig", "get input text", func() {
		input := analysis.GetInput()
		fmt.Println(input)
	})
	VigenereAnalysis.AddOption("rg", "find repetitions by length", func() {
		length, err := VigenereAnalysis.GetInt("length: ")
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		repetitions, err := analysis.GetRepetitions(length)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		for substring, detail := range repetitions {
			fmt.Printf("%s: {Positions: %v, Distances: %v}\n", substring, detail.Positions, detail.Distances)
		}
	})
	VigenereAnalysis.AddOption("ks", "set key length", func() {
		keyLength, err := VigenereAnalysis.GetInt("key length: ")
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		err = analysis.SetKeyLength(keyLength)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
	})
	VigenereAnalysis.AddOption("ip", "partition input in key length parts", func() {
		parts, err := analysis.GetInputParts()
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		for _, part := range parts {
			fmt.Println(part)
			frequencies, err := utils.AnalyzeSubstringFrequencies(part, 1)
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			letters := slices.Collect(maps.Keys(frequencies))
			slices.SortFunc(letters, func(a, b string) int {
				return -cmp.Compare(frequencies[a], frequencies[b])
			})
			for _, letter := range letters {
				fmt.Printf("[%s: %d] ", letter, frequencies[letter])
			}
			fmt.Println("")
		}
	})
	VigenereAnalysis.AddOption("id", "decipher input using key", func() {
		key := VigenereAnalysis.GetString("key: ")
		input := analysis.GetInput()
		text, err := vigenere.Decypher(input, key)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		fmt.Printf("text: %s\n", text)
	})
	VigenereAnalysis.AddOption("ibd", "decipher input using partial key", func() {
		fmt.Println("for each char of the key enter the possible options separated by comma with no spaces")
		fmt.Println("example: B,V,O,C,M")

		keyLength := analysis.GetKeyLength()
		options := make([][]string, 0)
		for i := range keyLength {
			options = append(options, make([]string, 0))
			prompt := fmt.Sprintf("char %d: ", i+1)
			chars := VigenereAnalysis.GetString(prompt)
			options[i] = strings.Split(chars, ",")
		}

		keys := make([]string, 0)

		// TODO: understand how to works exactly, didn't write this myself
		// Recursive function to handle the combinations
		var generateCombinations func(int, []string)
		generateCombinations = func(index int, currentKey []string) {
			if index == keyLength {
				// Base case: We've built a full key
				keys = append(keys, strings.Join(currentKey, ""))
				// Here you would use the key to attempt decryption
				return
			}

			for _, char := range options[index] {
				newKey := append(currentKey, char)
				generateCombinations(index+1, newKey)
			}
		}
		generateCombinations(0, []string{}) // Start the recursion

		for _, key := range keys {
			text, err := vigenere.Decypher(analysis.GetInput(), key)
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			if strings.Contains(text, "THE") && strings.Contains(text, "AND") {
				fmt.Printf("%s: %s\n", key, text)
			}
		}
	})
}
