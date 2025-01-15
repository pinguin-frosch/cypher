package main

import "fmt"

func main() {
	problemSeven()
}

func problemSeven() {
	numbers := []int{50, 20, 37, 95, 74, 00, 25, 51, 05, 78, 34, 32, 72}

	shiftValue := 2
	shifted := make([]int, 0, len(numbers))
	for i := range numbers {
		index := (len(numbers) - shiftValue + i) % len(numbers)
		shifted = append(shifted, numbers[index])
	}

	modValue := 19
	moded := make([]int, 0, len(numbers))
	for _, number := range shifted {
		moded = append(moded, number%modValue)
	}

	addValue := 3
	added := make([]int, 0, len(numbers))
	for _, number := range moded {
		added = append(added, number+addValue)
	}

	message := ""
	for _, number := range added {
		message += string(rune(number + 64))
	}
	fmt.Println(message)
}
