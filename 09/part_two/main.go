package part_two

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extractHistory(line string) []int {
	numbersStr := strings.Fields(line)
	var numbers []int
	for _, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Erreur de conversion : %v\n", err)
			return []int{}
		}
		// Ajout du nombre à la slice
		numbers = append(numbers, num)
	}
	return numbers

}

func allZero(list []int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] != 0 {
			return false
		}
	}
	return true
}

func calculSubSequences(sequence []int) []int {
	var subsequence []int
	for i := 1; i < len(sequence); i++ {
		subsequence = append(subsequence, sequence[i]-sequence[i-1])
	}
	return subsequence
}

func extrapolation(firstNumbSequence []int) int {
	actual := 0
	for i := len(firstNumbSequence) - 2; i >= 0; i-- {
		actual = firstNumbSequence[i] - actual
	}
	return actual
}

func Part_Two() int {
	result := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		history := extractHistory(scanner.Text())
		var lastItem []int
		for {
			lastItem = append(lastItem, history[0])

			if allZero(history) {
				break
			} else {
				history = calculSubSequences(history)
			}
		}
		result += extrapolation(lastItem)

	}

	return result
}
