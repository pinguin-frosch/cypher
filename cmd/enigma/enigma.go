package main

import (
	"cypher/polisubstitution/enigma"
	"fmt"
	"strings"
)

func main() {
	problemOne()
	problemTwo()
	problemThree()
}

func problemOne() {
	cipher := "ZYDNI"
	e := enigma.New()
	rotorI, _ := enigma.NewRotor("UWYGADFPVZBECKMTHXSLRINQOJ")
	reflector, _ := enigma.NewRotor("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	e.AddRotor(rotorI)
	e.AddReflector(reflector)
	result := e.Type(cipher)
	fmt.Println(result)
}

func problemTwo() {
	cipher := "QHSGUWIG"
	for i := range 26 {
		e := enigma.New()
		rotorI, _ := enigma.NewRotor("UWYGADFPVZBECKMTHXSLRINQOJ")
		rotorI.Rotate(i)
		reflector, _ := enigma.NewRotor("YRUHQSLDPXNGOKMIEBFZCWVJAT")
		e.AddRotor(rotorI)
		e.AddReflector(reflector)
		result := e.Type(cipher)
		if strings.HasPrefix(result, "XV") {
			fmt.Println(result)
		}
	}
}

func problemThree() {
	cipher := "GYHRVFLRXY"
	e := enigma.New()
	e.AddPlugboardConnection('A', 'B')
	e.AddPlugboardConnection('S', 'Z')
	e.AddPlugboardConnection('U', 'Y')
	e.AddPlugboardConnection('G', 'H')
	e.AddPlugboardConnection('L', 'Q')
	e.AddPlugboardConnection('E', 'N')

	rotorII, _ := enigma.NewRotor("AJPCZWRLFBDKOTYUQGENHXMIVS")
	rotorII.Rotate(strings.Index(enigma.Alphabet, "A"))
	e.AddRotor(rotorII)

	rotorI, _ := enigma.NewRotor("UWYGADFPVZBECKMTHXSLRINQOJ")
	rotorI.Rotate(strings.Index(enigma.Alphabet, "E"))
	e.AddRotor(rotorI)

	rotorIII, _ := enigma.NewRotor("TAGBPCSDQEUFVNZHYIXJWLRKOM")
	rotorIII.Rotate(strings.Index(enigma.Alphabet, "B"))
	e.AddRotor(rotorIII)

	reflector, _ := enigma.NewRotor("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	e.AddReflector(reflector)

	result := e.Type(cipher)
	fmt.Println(result)
}
