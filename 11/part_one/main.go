package part_one

import (
	"bytes"
)

type Coord struct {
	i, j int
}

func constructExtandedUniverse(input []byte) ([]Coord, []int, []int) {
	lines := bytes.Split(input, []byte("\n"))

	var hashtags []Coord
	var ligne []int
	countEmptyLine := 0
	countEmptyColonne := 0
	initilisation := false
	var colonneBool []bool
	for ind, line := range lines {
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
	return abs(a.i-b.i) + quantityOfLine(a, b, ligne)*1 + abs(a.j-b.j) + quantityOfColonne(a, b, colonne)*1
}

func quantityOfLine(a Coord, b Coord, list []int) int {
	return abs(list[a.i] - list[b.i])
}

func quantityOfColonne(a Coord, b Coord, list []int) int {
	return abs(list[a.j] - list[b.j])
}

func Part_One(input []byte) int {
	hashtags, lineTreated, colonneTreated := constructExtandedUniverse(input)

	result := 0

	for i := 0; i < len(hashtags); i++ {
		for j := i + 1; j < len(hashtags); j++ {
			//fmt.Println(coords[i], coords[j], findOptDistance(coords[i], coords[j], ligne, colonne))
			result += findOptDistance(hashtags[i], hashtags[j], lineTreated, colonneTreated)
		}
	}

	return result
}
