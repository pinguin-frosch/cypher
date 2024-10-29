package monosubstitution

import (
	"errors"
	"sort"
	"strings"
	"unicode"
)

type State struct {
	input        string
	cleanInput   string
	replacements map[rune]rune
	frequencies  map[int]map[string]substringFrequency
}

func NewState() *State {
	s := State{}
	s.frequencies = make(map[int]map[string]substringFrequency)
	s.replacements = make(map[rune]rune)
	return &s
}

func (s *State) AddInputText(input string) {
	s.input = input
	s.getCleanInput()
	s.reset()
}

func (s *State) getCleanInput() {
	s.cleanInput = ""
	for _, c := range s.input {
		if unicode.IsSpace(c) || unicode.IsPunct(c) {
			continue
		}
		s.cleanInput += string(c)
	}
}

func (s *State) GetReplacedText() string {
	text := s.input
	for from, to := range s.replacements {
		text = strings.ReplaceAll(text, string(from), string(to))
	}
	return text
}

func (s *State) GetInputText() string {
	return s.input
}

func (s *State) AddLetterReplacement(from, to rune) {
	s.replacements[from] = to
}

func (s *State) GetLetterReplacements() (map[rune]rune, []rune) {
	keys := make([]rune, 0, len(s.replacements))
	for key := range s.replacements {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return s.replacements[keys[i]] < s.replacements[keys[j]]
	})
	return s.replacements, keys
}

func (s *State) reset() {
	clear(s.replacements)
	clear(s.frequencies)
}

type substringFrequency struct {
	Times      int
	Percentage float64
}

func (s *State) GetNFrecuencies(length int) (map[string]substringFrequency, error) {
	if length <= 0 {
		return nil, errors.New("invalid length")
	}
	if _, ok := s.frequencies[length]; !ok {
		err := s.analizeNFrequencies(length)
		if err != nil {
			return nil, err
		}
	}
	return s.frequencies[length], nil
}

func (s *State) analizeNFrequencies(length int) error {
	if length <= 0 {
		panic("invalid length")
	}
	frequencies := make(map[string]substringFrequency)
	for i := 0; i < len(s.cleanInput)-(length+1); i++ {
		substring := s.cleanInput[i : i+length]
		if f, ok := frequencies[substring]; ok {
			f.Times++
			frequencies[substring] = f
		} else {
			f := substringFrequency{Times: 1}
			frequencies[substring] = f
		}
	}
	for substring, f := range frequencies {
		f.Percentage = float64(f.Times) / float64(len(s.cleanInput))
		frequencies[substring] = f
	}
	s.frequencies[length] = frequencies
	return nil
}
