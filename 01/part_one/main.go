package part_one

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func findFirstDigit(word string) int {
	for _, char := range word {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	return -1
}

func findLastDigit(word string) int {
	runes := []rune(word)

	for i := len(runes) - 1; i >= 0; i-- {
		char := runes[i]
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	return -1
}

func Part_One() int {
	file, err := os.Open("part_one/input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		result += findFirstDigit(line)*10 + findLastDigit(line)
	}
	return result
}
