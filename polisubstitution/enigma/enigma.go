package enigma

import (
	"cypher/utils"
	"errors"
	"fmt"
	"slices"
	"strings"
)

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Enigma struct {
	rotors    []Rotor
	reflector Rotor
	plugboard map[rune]rune
}

type Rotor struct {
	alphabet string
	rotor    string
}

func New() *Enigma {
	e := Enigma{}
	e.rotors = make([]Rotor, 0)
	e.plugboard = make(map[rune]rune)
	return &e
}

func (e *Enigma) AddPlugboardConnection(from, to rune) error {
	if !utils.IsLatinLetter(from) {
		return fmt.Errorf("%s is an invalid character", string(from))
	}
	if !utils.IsLatinLetter(to) {
		return fmt.Errorf("%s is an invalid character", string(to))
	}
	if from == to {
		return fmt.Errorf("%s cannot connect to itself", string(from))
	}
	if len(e.plugboard) >= 10 {
		return errors.New("cannot have more than 9 connections")
	}
	for k, v := range e.plugboard {
		if k == from || k == to {
			return fmt.Errorf("%s is already connected with %s", string(k), string(v))
		}
		if v == from || v == to {
			return fmt.Errorf("%s is already connected with %s", string(k), string(v))
		}
	}
	e.plugboard[from] = to
	return nil
}

// Simulates typing on the enigma machine
func (e *Enigma) Type(message string) string {
	message = strings.ToUpper(message)
	result := ""
	for _, letter := range message {
		if !utils.IsLatinLetter(letter) {
			result += string(letter)
		}
		e.RotateRotors()
		letter = e.replacePlugboard(letter)
		for _, rotor := range e.rotors {
			letter = e.passThroughRotor(letter, rotor, false)
		}
		letter = e.passThroughRotor(letter, e.reflector, false)
		for _, rotor := range slices.Backward(e.rotors) {
			letter = e.passThroughRotor(letter, rotor, true)
		}
		letter = e.replacePlugboard(letter)
		result += string(letter)
	}
	return result
}

func (e *Enigma) replacePlugboard(letter rune) rune {
	if replacement, ok := e.plugboard[letter]; ok {
		return replacement
	}
	for k, v := range e.plugboard {
		if v == letter {
			return k
		}
	}
	return letter
}

func (e *Enigma) passThroughRotor(letter rune, rotor Rotor, reverse bool) rune {
	var a, r string
	if !reverse {
		a, r = rotor.alphabet, rotor.rotor
	} else {
		a, r = rotor.rotor, rotor.alphabet
	}
	index := strings.Index(Alphabet, string(letter))
	newIndex := strings.Index(r, string(a[index]))
	// fmt.Printf("%s -> %s\n", string(index+65), string(newIndex+65))
	return rune(Alphabet[newIndex])
}

// Rotates the rotors before every key press
func (e *Enigma) RotateRotors() {
	// TODO: rotate not just the first rotor
	e.rotors[0].Rotate(1)
}

func RotateRotorString(rotor string, n int) string {
	return rotor[n:] + rotor[:n]
}

// Rotates a rotor by moving each letter one position
func (r *Rotor) Rotate(n int) {
	r.rotor = RotateRotorString(r.rotor, n)
	r.alphabet = RotateRotorString(r.alphabet, n)
}

// Creates a new rotor, and adds an alphabet alongside it
func NewRotor(rotor string) (Rotor, error) {
	rotor = strings.ToUpper(rotor)
	err := IsValidRotor(rotor)
	if err != nil {
		return Rotor{}, err
	}
	r := Rotor{alphabet: Alphabet, rotor: rotor}
	return r, nil
}

// Adds a new rotor, they run in the order that they are added
func (e *Enigma) AddRotor(rotor Rotor) error {
	e.rotors = append(e.rotors, rotor)
	return nil
}

// Adds the reflector, calling it again will replace the reflector
func (e *Enigma) AddReflector(reflector Rotor) error {
	e.reflector = reflector
	return nil
}

// Checks whether a rotor is valid or not
func IsValidRotor(rotor string) error {
	if len(rotor) != 26 {
		return errors.New("invalid rotor: it has to be 26 characters long")
	}
	frequencies := make(map[rune]int)
	for _, c := range rotor {
		if !utils.IsLatinLetter(c) {
			return fmt.Errorf("invalid letter in rotor: %s", string(c))
		}
		frequencies[c] += 1
	}
	for k, v := range frequencies {
		if v > 1 {
			return fmt.Errorf("%s appears more than once", string(k))
		}
	}
	return nil
}

func (e *Enigma) Print() {
	for i, rotor := range e.rotors {
		fmt.Printf("Rotor %d: %s\n", i+1, rotor)
	}
	fmt.Printf("Reflector: %s\n", e.reflector)
}
