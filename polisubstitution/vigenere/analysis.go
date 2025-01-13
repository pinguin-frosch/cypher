package vigenere

import (
	"cypher/utils"
	"errors"
)

type Analysis struct {
	input       string
	cleanInput  string
	repetitions map[int]map[string]substringDetail
	keyLength   int
	inputParts  []string
}

func NewAnalysis() *Analysis {
	analysis := Analysis{}
	analysis.repetitions = make(map[int]map[string]substringDetail)
	analysis.inputParts = make([]string, 0)
	return &analysis
}

func (a *Analysis) reset() {
	clear(a.repetitions)
	a.keyLength = 0
	a.resetInputParts()
}

func (a *Analysis) AddInput(input string) {
	a.input = input
	a.getCleanInput()
	a.reset()
}

func (a *Analysis) getCleanInput() {
	a.cleanInput = ""
	for _, c := range a.input {
		if !utils.IsLatinLetter(c) {
			continue
		}
		a.cleanInput += string(c)
	}
}

func (a *Analysis) GetInput() string {
	return a.input
}

func (a *Analysis) GetRepetitions(length int) (map[string]substringDetail, error) {
	if _, ok := a.repetitions[length]; !ok {
		err := a.findRepetitions(length)
		if err != nil {
			return nil, err
		}
	}
	return a.repetitions[length], nil
}

type substringDetail struct {
	Positions []int
	Distances []int
}

// Finds substrings repetitions in the text, only keeping the ones that appear
// at least twice
func (a *Analysis) findRepetitions(length int) error {
	if length <= 1 {
		return errors.New("invalid length: use at least 2")
	}

	// get all substrings repetitions
	repetitions := make(map[string]substringDetail)
	for i := 0; i < len(a.cleanInput)-(length-1); i++ {
		substring := a.cleanInput[i : i+length]
		if detail, ok := repetitions[substring]; ok {
			detail.Positions = append(detail.Positions, i)
			repetitions[substring] = detail
		} else {
			detail.Positions = make([]int, 0)
			detail.Positions = append(detail.Positions, i)
			repetitions[substring] = detail
		}
	}

	// filter the ones that appear at least twice
	filteredRepetitions := make(map[string]substringDetail)
	for substring, detail := range repetitions {
		if len(detail.Positions) > 1 {
			filteredRepetitions[substring] = detail
		}
	}
	repetitions = filteredRepetitions

	// get distances for each pair of positions
	for substring, detail := range repetitions {
		detail.Distances = make([]int, 0, len(detail.Positions)-1)
		for i := 0; i < len(detail.Positions)-1; i++ {
			current, next := detail.Positions[i], detail.Positions[i+1]
			detail.Distances = append(detail.Distances, next-current)
		}
		repetitions[substring] = detail
	}

	// save repetitions for later
	a.repetitions[length] = repetitions

	return nil
}

func (a *Analysis) SetKeyLength(keyLength int) error {
	if keyLength <= 1 {
		return errors.New("invalid key length: has to be at least 2")
	}
	a.keyLength = keyLength
	return nil
}

func (a *Analysis) GetKeyLength() int {
	return a.keyLength
}

func (a *Analysis) GetInputParts() ([]string, error) {
	if a.keyLength <= 1 {
		return nil, errors.New("key length not set: specify a length first")
	}
	if len(a.inputParts) != a.keyLength {
		a.resetInputParts()
		a.partitionInput()
	}
	return a.inputParts, nil
}

func (a *Analysis) resetInputParts() {
	a.inputParts = make([]string, 0)
}

func (a *Analysis) partitionInput() {
	// initialize string for each part
	for range a.keyLength {
		a.inputParts = append(a.inputParts, "")
	}

	// populate every partition
	for i, c := range a.cleanInput {
		bucket := i % a.keyLength
		a.inputParts[bucket] += string(c)
	}
}
