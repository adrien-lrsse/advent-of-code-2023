package part_two

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	i, j int
}

func constructExtandedUniverse(input string) ([]Coord, []int, []int) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file :", err)
		return []Coord{}, []int{}, []int{}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var hashtags []Coord
	var ligne []int
	ind := 0
	countEmptyLine := 0
	countEmptyColonne := 0
	initilisation := false
	var colonneBool []bool
	for scanner.Scan() {
		line := scanner.Text()
		if !initilisation {
			colonneBool = make([]bool, len(line))
			initilisation = true
		}
		isStar := true
		for index, caractere := range line {
			if caractere == '#' {
				hashtags = append(hashtags, Coord{ind, index})
				colonneBool[index] = true
				isStar = false
			}
		}
		if isStar {
			countEmptyLine++
		}
		ligne = append(ligne, countEmptyLine)

		ind++
	}

	var colonne []int
	for i := 0; i < len(colonneBool); i++ {
		if !colonneBool[i] {
			countEmptyColonne++
		}
		colonne = append(colonne, countEmptyColonne)

	}
	return hashtags, ligne, colonne

}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func findOptDistance(a Coord, b Coord, ligne []int, colonne []int) int {
	return abs(a.i-b.i) + quantityOfLine(a, b, ligne)*999999 + abs(a.j-b.j) + quantityOfColonne(a, b, colonne)*999999
}

func quantityOfLine(a Coord, b Coord, list []int) int {
	return abs(list[a.i] - list[b.i])
}

func quantityOfColonne(a Coord, b Coord, list []int) int {
	return abs(list[a.j] - list[b.j])
}

func Part_Two() int {
	hashtags, lineTreated, colonneTreated := constructExtandedUniverse("input.txt")

	result := 0

	for i := 0; i < len(hashtags); i++ {
		for j := i + 1; j < len(hashtags); j++ {
			//fmt.Println(coords[i], coords[j], findOptDistance(coords[i], coords[j], ligne, colonne))
			result += findOptDistance(hashtags[i], hashtags[j], lineTreated, colonneTreated)
		}
	}

	return result
}
