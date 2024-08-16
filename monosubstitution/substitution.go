package monosubstitution

import (
	"unicode"
)

type State struct {
	input       string
	frequencies map[rune]letterFrequency
}

func NewState() *State {
	s := State{}
	s.frequencies = make(map[rune]letterFrequency)
	return &s
}

func (s *State) AddInputText(input string) {
	s.input = input
	s.reset()
	s.analizeFrequencies()
}

func (s *State) reset() {
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
