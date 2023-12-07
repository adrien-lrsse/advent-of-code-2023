package part_one

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func parseSeeds(line string) []int {
	var seeds []int
	line_decoupe := strings.Split(line, `:`)

	runes := []rune(line_decoupe[1])

	tempon := 0

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			tempon = tempon*10 + int(runes[i]-'0')
		}
		if !(unicode.IsDigit(runes[i])) && (tempon != 0) {
			seeds = append(seeds, tempon)
			tempon = 0
		}
	}

	if tempon != 0 {
		seeds = append(seeds, tempon)

	}
	return seeds
}

func parseMap(line string) []int {
	var maps []int

	runes := []rune(line)

	tempon := -1

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) && tempon == -1 {
			tempon = int(runes[i] - '0')
		} else if unicode.IsDigit(runes[i]) {
			tempon = tempon*10 + int(runes[i]-'0')
		}
		if !(unicode.IsDigit(runes[i])) && (tempon != -1) {
			maps = append(maps, tempon)
			tempon = -1
		}
	}

	if tempon != -1 {
		maps = append(maps, tempon)
	}

	return maps
}

func findMin(list []int) int {
	min := list[0]
	for i := 1; i < len(list); i++ {
		if min > list[i] {
			min = list[i]
		}
	}
	return min
}

func findValueOfSeeds(destination int, source int, length int, seed int) int {
	var intervalleSource []int

	intervalleSource = append(intervalleSource, source)
	intervalleSource = append(intervalleSource, source+length-1)

	if seed >= intervalleSource[0] && seed <= intervalleSource[1] {
		ecart := seed - intervalleSource[0]

		return destination + ecart
	}

	return seed
}

func Part_One() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}
	defer file.Close()

	reset := false
	var isCheck []bool

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	seeds := parseSeeds(line)
	for i := 0; i < len(seeds); i++ {
		isCheck = append(isCheck, false)
	}
	for scanner.Scan() {
		line = scanner.Text()

		if reset {
			for i := 0; i < len(isCheck); i++ {
				isCheck[i] = false
			}
		}
		if len(strings.TrimSpace(line)) == 0 {
			reset = true
		} else if !unicode.IsDigit(rune(line[0])) {
			reset = true
		} else {
			reset = false
			map_i := parseMap(line)
			for i := 0; i < len(seeds); i++ {
				if !isCheck[i] {
					temp := seeds[i]
					seeds[i] = findValueOfSeeds(map_i[0], map_i[1], map_i[2], seeds[i])
					if temp != seeds[i] {
						isCheck[i] = true

					}
				}
			}

		}

	}

	return findMin(seeds)
}
