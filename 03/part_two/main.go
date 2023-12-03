package part_two

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

func detectStar(line string) []int {
	var posSymbol []int

	runes := []rune(line)

	for i := 0; i < len(runes); i++ {
		if runes[i] == 42 {
			posSymbol = append(posSymbol, i)
		}
	}
	return posSymbol
}

func adjacent(pos_symbol int, previous_line_pos [][]int, actual_line_pos [][]int, next_line_pos [][]int, previous_number []int, actual_number []int, next_number []int) []int {
	var tempo []int

	// checking on the previous line
	for i := 0; i < len(previous_number); i++ {
		if previous_line_pos[i][0]-1 <= pos_symbol && previous_line_pos[i][1]+1 >= pos_symbol {
			tempo = append(tempo, previous_number[i])
		}
	}

	// cheking on the line

	for i := 0; i < len(actual_number); i++ {
		if actual_line_pos[i][0]-1 == pos_symbol || actual_line_pos[i][1]+1 == pos_symbol {
			tempo = append(tempo, actual_number[i])

		}
	}

	//checking on the next line
	for i := 0; i < len(next_number); i++ {
		if next_line_pos[i][0]-1 <= pos_symbol && next_line_pos[i][1]+1 >= pos_symbol {
			tempo = append(tempo, next_number[i])
		}
	}
	return tempo

}

// Fonction pour copier une tranche

func Part_Two() int {
	var previous_number_pos [][]int = [][]int{}
	var previous_number []int
	var actual_star_pos []int
	var actual_number_pos [][]int
	var actual_number []int
	var next_number_pos [][]int
	var next_number []int

	file, err := os.Open("part_two/input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	scanner.Scan()
	line := scanner.Text()

	actual_star_pos = detectStar(line)
	actual_number, actual_number_pos = detectDigit(line)

	scanner.Scan()
	line = scanner.Text()

	next_number, next_number_pos = detectDigit(line)

	for scanner.Scan() {
		for i := 0; i < len(actual_star_pos); i++ {
			adjacent := adjacent(actual_star_pos[i], previous_number_pos, actual_number_pos, next_number_pos, previous_number, actual_number, next_number)
			if len(adjacent) == 2 {
				result += (adjacent[0] * adjacent[1])
			}
		}

		actual_star_pos = detectStar(line)
		line = scanner.Text()

		previous_number = actual_number
		previous_number_pos = actual_number_pos

		actual_number = next_number
		actual_number_pos = next_number_pos

		next_number, next_number_pos = detectDigit(line)

	}

	for i := 0; i < len(actual_star_pos); i++ {
		adjacent := adjacent(actual_star_pos[i], previous_number_pos, actual_number_pos, next_number_pos, previous_number, actual_number, next_number)
		if len(adjacent) == 2 {
			result += (adjacent[0] * adjacent[1])
		}
	}

	previous_number = actual_number
	previous_number_pos = actual_number_pos

	actual_number = next_number
	actual_number_pos = next_number_pos
	actual_star_pos = detectStar(line)

	next_number, next_number_pos = []int{}, [][]int{}

	for i := 0; i < len(actual_star_pos); i++ {
		adjacent := adjacent(actual_star_pos[i], previous_number_pos, actual_number_pos, next_number_pos, previous_number, actual_number, next_number)
		if len(adjacent) == 2 {
			result += (adjacent[0] * adjacent[1])
		}
	}
	return result

}
