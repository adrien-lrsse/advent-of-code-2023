package part_two

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

func max(slice []int) int {
	max := slice[0]
	for i := 0; i < len(slice); i++ {
		if max < slice[i] {
			max = slice[i]
		}
	}
	return max
}

func applyPart_One(start string, instruction []rune, coordMap map[string][]string) int {

	count := 0
	current := start

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
		if current[2] == 'Z' {
			break
		}
	}

	return count * len(instruction)
}

func allElementAreEquals(slice []int) bool {
	tmp := slice[0]
	for i := 0; i < len(slice); i++ {
		if tmp != slice[i] {
			return false
		}

	}
	return true

}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pgcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Fonction pour calculer le PPCM de deux nombres
func ppcm(a, b int) int {
	// PPCM(a, b) = |a * b| / PGCD(a, b)
	absA, absB := abs(a), abs(b)
	gcd := pgcd(absA, absB)
	return absA * absB / gcd
}

func Part_Two() int {
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

	var startList []string
	for scanner.Scan() {
		line = scanner.Text()
		line_result := addCoord(line)
		addToMap(coordMap, line_result[0], []string{line_result[1], line_result[2]})
	}

	for key := range coordMap {
		if key[2] == 'A' {
			startList = append(startList, key)
		}
	}

	var tourne []int
	ending := 0
	for i := 0; i < len(startList); i++ {
		ending = applyPart_One(startList[i], instruction, coordMap)
		tourne = append(tourne, ending)
	}

	result := tourne[0]
	for i := 1; i < len(tourne); i++ {
		result = ppcm(result, tourne[i])
	}

	return result
}
