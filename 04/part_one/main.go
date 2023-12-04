package part_one

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"
)

func lineParsing(line string) (string, string) {
	dividedOnPipe := strings.Split(line, `|`)
	leftSide := strings.Split(dividedOnPipe[0], `:`)
	return leftSide[1], dividedOnPipe[1]
}

func convertNumbersInList(word string) []int {
	runes := []rune(word)
	var result []int
	tempon := 0
	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			tempon = tempon*10 + int(runes[i]-'0')
		}
		if !unicode.IsDigit(runes[i]) && (tempon != 0) {
			result = append(result, tempon)
			tempon = 0

		}
	}
	if tempon != 0 {
		result = append(result, tempon)

	}
	return result

}

func countMatching(winningNumber []int, myNumber []int) int {
	result := 0
	for i := 0; i < len(myNumber); i++ {
		if slices.Contains(winningNumber, myNumber[i]) {
			result += 1
		}
	}
	return result

}

func Part_One() int {

	file, err := os.Open("part_one/input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}
	defer file.Close()

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		winningNumber, myNumber := lineParsing(line)
		listWinningNumber, ListMyNumber := convertNumbersInList(winningNumber), convertNumbersInList(myNumber)
		matching := countMatching(listWinningNumber, ListMyNumber)
		temp := 0
		if matching != 0 {
			temp = 1
		}
		for i := 1; i < matching; i++ {
			temp *= 2
		}
		result += temp
	}

	return result
}
