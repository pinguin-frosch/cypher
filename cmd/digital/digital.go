package main

import (
	"fmt"
	"strconv"
)

func main() {
	problemTwo()
	problemThree()
	problemFour()
}

func problemTwo() {
	key := "HELLO"
	values := []string{"00111111", "00101010", "00111110", "00100000", "00101011"}
	numbers := make([]int, 0, 5)
	for _, value := range values {
		number, err := strconv.ParseInt(value, 2, 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, int(number))
	}
	resultNumbers := make([]rune, 0, 5)
	for i, c := range key {
		r := c ^ rune(numbers[i])
		resultNumbers = append(resultNumbers, r)
	}
	result := ""
	for _, number := range resultNumbers {
		result += string(number)
	}
	fmt.Println(result)
}

func problemThree() {
	numbers := []string{
		"01100010",
		"01110101",
		"01111001",
		"01100011",
		"01110101",
		"01100110",
		"01101001",
	}
	numbersTransposed := make([]string, 8)
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			numbersTransposed[j] += string(numbers[i][j])
		}
	}
	offsets := []int{2, 1, 3, 1, 2, 5, 7, 2}
	numbersOffseted := make([]string, 8)
	for i, offset := range offsets {
		numberOffseted := numbersTransposed[i][7-offset:] + numbersTransposed[i][:7-offset]
		numbersOffseted[i] = numberOffseted
	}
	numbersFinal := make([]string, len(numbers))
	for i := range numbers {
		for j := range numbersOffseted {
			numbersFinal[i] += string(numbersOffseted[j][i])
		}
	}
	result := ""
	for _, number := range numbersFinal {
		value, err := strconv.ParseInt(number, 2, 64)
		if err != nil {
			panic(err)
		}
		result += string(value)
	}
	fmt.Println(result)
}

func problemFour() {
	numbers := []int{
		254, 1, 0, 1,
		254, 254, 254, 0,
		254, 255, 254, 0,
		0, 255, 1, 1,
	}
	// rgb(254, 254, 254), rgb(1,   1,   1),   rgb(0,   0,   0),   rgb(1, 1, 1)
	// rgb(254, 254, 254), rgb(254, 254, 254), rgb(254, 254, 254), rgb(0, 0, 0)
	// rgb(254, 254, 254), rgb(255, 255, 255), rgb(254, 254, 254), rgb(0, 0, 0)
	// rgb(0,   0,   0),   rgb(255, 255, 255), rgb(1,   1,   1),   rgb(1, 1, 1)
	// .J
	fmt.Printf(".J")
	numbers = []int{0b01010000, 0b01000111}
	for _, number := range numbers {
		fmt.Printf("%s", string(number))
	}
	fmt.Println()
}
