package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parsing(line string) int {

	isSeparator := func(r rune) bool {
		return r == ',' || r == ';'
	}

	process_game := strings.Split(line, ":")

	devidedLine := strings.FieldsFunc(process_game[1], isSeparator)

	game := strings.Split(process_game[0], " ")

	game_number, err := strconv.Atoi(game[1])
	if err != nil {
		fmt.Println("Erreur de conversion :", err)
		return 0
	}

	var number_color []int

	var color []string

	for i := 0; i < len(devidedLine); i++ {
		intermediate := strings.Split(devidedLine[i], " ")

		integer, err := strconv.Atoi(intermediate[1])
		if err != nil {
			fmt.Println("Erreur de conversion :", err)
			return 0
		}

		number_color = append(number_color, integer)
		color = append(color, intermediate[2])

	}
	fmt.Println(game_number)

	if (mathcingColor(number_color, color)) == false {
		game_number = 0
	}

	fmt.Println(game_number)
	return game_number
}

func mathcingColor(integers []int, strings []string) bool {
	max_red := 12
	red := 0
	max_green := 13
	green := 0
	max_blue := 14
	blue := 0

	fmt.Println(strings, integers)

	for i := 0; i < len(integers); i++ {
		switch strings[i] {
		case "red":
			if integers[i] > max_red {
				return false
			}
		case "green":
			if integers[i] > max_green {
				return false
			}
		case "blue":
			if integers[i] > max_blue {
				return false
			}
		}
	}

	fmt.Println(red, green, blue)
	return true
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		result += parsing(line)
	}
	fmt.Println(result)

}
