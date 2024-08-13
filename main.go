package main

import (
	"bufio"
	"cypher/transposition"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">> ")
	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), " ")
		text = strings.ToLower(text)
		switch text {
		case "b":
			fmt.Print("Message: ")
			scanner.Scan()
			transposition.BruteForceDecipher(scanner.Text())
		}
		fmt.Print(">> ")
	}
}
