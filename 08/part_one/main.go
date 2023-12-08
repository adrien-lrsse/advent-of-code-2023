package part_one

import (
	"bufio"
	"fmt"
	"os"
)

func extractInstruction(line string) []rune {
	return []rune(line)
}

func runesToString(runes []rune) (outString string) {
	// don't need index so _
	for _, v := range runes {
		outString += string(v)
	}
	return
}

func addCoord(line string) []string {
	runes := []rune(line)

	var actual []rune
	actual = append(actual, runes[0], runes[1], runes[2])

	var X []rune
	X = append(X, runes[7], runes[8], runes[9])

	var Y []rune
	Y = append(Y, runes[12], runes[13], runes[14])

	return []string{runesToString(actual), runesToString(X), runesToString(Y)}
}

func addToMap(myMap map[string][]string, key string, value []string) {
	myMap[key] = value
}
func Part_One() int {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}
	defer file.Close()

	coordMap := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	instruction := extractInstruction(line)
	scanner.Scan()
	for scanner.Scan() {
		line = scanner.Text()
		line_result := addCoord(line)
		addToMap(coordMap, line_result[0], []string{line_result[1], line_result[2]})
	}

	count := 0
	current := "AAA"

	for {
		count += 1
		for i := 0; i < len(instruction); i++ {
			if instruction[i] == 'L' {
				current = coordMap[current][0]
			}
			if instruction[i] == 'R' {
				current = coordMap[current][1]
			}
		}
		if current == "ZZZ" {
			break
		}
	}

	return count * len(instruction)
}
