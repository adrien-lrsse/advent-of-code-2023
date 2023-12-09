package part_two

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func convertSubStringInDigit(word string) map[int]int {

	occurences := make(map[int]int, 0)

	if strings.Contains(word, "one") {
		occurences[strings.Index(word, "one")] = 1
		occurences[strings.LastIndex(word, "one")] = 1
	}
	if strings.Contains(word, "two") {
		occurences[strings.Index(word, "two")] = 2
		occurences[strings.LastIndex(word, "two")] = 2

	}
	if strings.Contains(word, "three") {
		occurences[strings.Index(word, "three")] = 3
		occurences[strings.LastIndex(word, "three")] = 3

	}
	if strings.Contains(word, "four") {
		occurences[strings.Index(word, "four")] = 4
		occurences[strings.LastIndex(word, "four")] = 4

	}
	if strings.Contains(word, "five") {
		occurences[strings.Index(word, "five")] = 5
		occurences[strings.LastIndex(word, "five")] = 5

	}
	if strings.Contains(word, "six") {
		occurences[strings.Index(word, "six")] = 6
		occurences[strings.LastIndex(word, "six")] = 6

	}
	if strings.Contains(word, "seven") {
		occurences[strings.Index(word, "seven")] = 7
		occurences[strings.LastIndex(word, "seven")] = 7

	}
	if strings.Contains(word, "eight") {
		occurences[strings.Index(word, "eight")] = 8
		occurences[strings.LastIndex(word, "eight")] = 8

	}
	if strings.Contains(word, "nine") {
		occurences[strings.Index(word, "nine")] = 9
		occurences[strings.LastIndex(word, "nine")] = 9

	}

	return occurences

}

func findFirstIndice(m map[int]int) int {
	keyMin := 5000000
	for key := range m {
		if keyMin > key {
			keyMin = key
		}
	}
	return keyMin
}

func findLastIndice(m map[int]int) int {
	keyMax := -1
	for key := range m {
		if keyMax < key {
			keyMax = key
		}
	}
	return keyMax
}

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

func splitStringOnDigit(word string) (string, string, bool, bool) {

	digitPattern := regexp.MustCompile(`\d+`)

	firstIsDigit := false

	lastIsDigit := false

	runes := []rune(word)

	if unicode.IsDigit(runes[0]) {
		firstIsDigit = true
	}

	if unicode.IsDigit(runes[len(runes)-1]) {
		lastIsDigit = true
	}

	result := digitPattern.Split(word, -1)

	return result[0], result[len(result)-1], firstIsDigit, lastIsDigit

}

func analyseLine(line string) int {

	firstDigit := 0
	lastDigit := 0

	firstString, lastString, firstBool, lastBool := splitStringOnDigit(line)

	if firstBool == true {
		firstDigit = findFirstDigit(line)
	} else {
		m := convertSubStringInDigit(firstString)
		if len(m) > 0 {
			tempon := findFirstIndice(m)
			// fmt.Println(tempon, m[tempon])
			firstDigit = m[tempon]
		} else {
			firstDigit = findFirstDigit(line)
		}
	}

	if lastBool == true {
		lastDigit = findLastDigit(line)
	} else {
		m := convertSubStringInDigit(lastString)
		if len(m) > 0 {
			tempon := findLastIndice(m)
			lastDigit = m[tempon]
		} else {
			lastDigit = findLastDigit(line)
		}
	}

	return firstDigit*10 + lastDigit

}

func Part_Two() int {

	// test := "onetwooneonef33 33"

	// fmt.Println(analyseLine(test))

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

		result += analyseLine(line)
	}
	return result
}
