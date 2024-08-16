package monosubstitution

import (
	"strings"
	"unicode"
)

type State struct {
	input        string
	replacements map[rune]rune
	frequencies  map[rune]letterFrequency
}

func NewState() *State {
	s := State{}
	s.frequencies = make(map[rune]letterFrequency)
	s.replacements = make(map[rune]rune)
	return &s
}

func (s *State) AddInputText(input string) {
	s.input = input
	s.reset()
	s.analizeFrequencies()
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

func (s *State) GetLetterReplacements() map[rune]rune {
	return s.replacements
}

func (s *State) reset() {
	clear(s.replacements)
	clear(s.frequencies)
}

type letterFrequency struct {
	Times      int
	Percentage float32
}

func (s *State) analizeFrequencies() {
	length := 0
	for _, c := range s.input {
		if unicode.IsSpace(c) || unicode.IsPunct(c) {
			continue
		}
		if freq, ok := s.frequencies[c]; ok {
			freq.Times++
			s.frequencies[c] = freq
		} else {
			s.frequencies[c] = letterFrequency{}
		}
		length++
	}
	for r, freq := range s.frequencies {
		freq.Percentage = float32(freq.Times) / float32(length) * 100
		s.frequencies[r] = freq
	}
}

func (s *State) GetLetterFrequencies() map[rune]letterFrequency {
	return s.frequencies
}
