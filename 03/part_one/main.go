package part_one

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func detectDigit(line string) ([]int, [][]int) {
	var numberInLine []int

	var posNumber [][]int

	runes := []rune(line)

	tempon := 0

	temp_i := 0
	temp_j := 0

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) && tempon == 0 {
			temp_i = i

		}
		if unicode.IsDigit(runes[i]) {
			tempon = tempon*10 + int(runes[i]-'0')
		}
		if !unicode.IsDigit(runes[i]) && (tempon != 0) {
			numberInLine = append(numberInLine, tempon)
			tempon = 0
			temp_j = i - 1
			posNumber = append(posNumber, []int{temp_i, temp_j})

		}
	}
	if tempon != 0 {
		numberInLine = append(numberInLine, tempon)
		tempon = 0
		temp_j = len(runes) - 1
		posNumber = append(posNumber, []int{temp_i, temp_j})

	}

	return numberInLine, posNumber

}

func detectSymbol(line string) []int {
	var symbolInLine []rune

	var posSymbol []int

	runes := []rune(line)

	for i := 0; i < len(runes); i++ {
		if !unicode.IsDigit(runes[i]) && runes[i] != '.' {
			symbolInLine = append(symbolInLine, runes[i])
			posSymbol = append(posSymbol, i)
		}
	}
	return posSymbol
}

func adjacent(pos_i int, pos_j int, symbole_Moins []int, symbole_Plus []int, symbole []int) bool {
	// checking on the previous line
	for i := 0; i < len(symbole_Moins); i++ {
		if symbole_Moins[i] >= pos_i-1 && symbole_Moins[i] <= pos_j+1 {
			return true
		}
	}

	// cheking on the line
	for i := 0; i < len(symbole); i++ {
		if symbole[i] == pos_i-1 || symbole[i] == pos_j+1 {
			return true
		}
	}

	//checking on the next line
	for i := 0; i < len(symbole_Plus); i++ {
		if symbole_Plus[i] >= pos_i-1 && symbole_Plus[i] <= pos_j+1 {
			return true
		}
	}
	return false

}

// Fonction pour copier une tranche

func Part_One() int {
	var previous_symbol_pos_line []int = []int{}
	var actual_pos_line [][]int
	var actual_symbol_pos_line []int
	var actual_number_line []int
	var next_pos_line [][]int
	var next_symbol_pos_line []int
	var next_number_line []int

	file, err := os.Open("part_one/input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	scanner.Scan()
	line := scanner.Text()

	actual_symbol_pos_line = detectSymbol(line)
	actual_number_line, actual_pos_line = detectDigit(line)

	scanner.Scan()
	line = scanner.Text()

	next_symbol_pos_line = detectSymbol(line)
	next_number_line, next_pos_line = detectDigit(line)

	for scanner.Scan() {
		for i := 0; i < len(actual_number_line); i++ {
			if adjacent(actual_pos_line[i][0], actual_pos_line[i][1], previous_symbol_pos_line, next_symbol_pos_line, actual_symbol_pos_line) {
				result += actual_number_line[i]
			}
		}

		line = scanner.Text()

		previous_symbol_pos_line = actual_symbol_pos_line
		actual_number_line = next_number_line
		actual_pos_line = next_pos_line
		actual_symbol_pos_line = next_symbol_pos_line
		actual_pos_line = next_pos_line

		next_symbol_pos_line = detectSymbol(line)
		next_number_line, next_pos_line = detectDigit(line)

	}

	for i := 0; i < len(actual_number_line); i++ {

		if adjacent(actual_pos_line[i][0], actual_pos_line[i][1], previous_symbol_pos_line, next_symbol_pos_line, actual_symbol_pos_line) {

			result += actual_number_line[i]

		}
	}

	previous_symbol_pos_line = actual_symbol_pos_line
	actual_number_line = next_number_line
	actual_pos_line = next_pos_line
	actual_symbol_pos_line = next_symbol_pos_line
	actual_pos_line = next_pos_line

	next_symbol_pos_line = []int{}
	next_number_line, next_pos_line = []int{}, [][]int{}

	for i := 0; i < len(actual_number_line); i++ {

		if adjacent(actual_pos_line[i][0], actual_pos_line[i][1], previous_symbol_pos_line, next_symbol_pos_line, actual_symbol_pos_line) {

			result += actual_number_line[i]

		}
	}
	return result
}
