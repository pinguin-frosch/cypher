package main

import (
	"cypher/transposition"
	"fmt"
	"log"
)

func main() {
	cipher, err := transposition.Cipher("thejewelsarehiddenbeneaththeelmtree", 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cipher)
}
