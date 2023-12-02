package part_two

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

	return mathcingColor(number_color, color)
}

func mathcingColor(integers []int, strings []string) int {
	red := 0
	green := 0
	blue := 0

	for i := 0; i < len(integers); i++ {
		switch strings[i] {
		case "red":
			if integers[i] > red {
				red = integers[i]
			}
		case "green":
			if integers[i] > green {
				green = integers[i]
			}
		case "blue":
			if integers[i] > blue {
				blue = integers[i]
			}
		}
	}

	return blue * red * green
}

func Part_Two() {

	file, err := os.Open("part_two/input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	// start := time.Now()
	for scanner.Scan() {
		line := scanner.Text()

		result += parsing(line)
	}
	// fmt.Println("Part 2: ", result)
	// fmt.Println(time.Since(start))

}
