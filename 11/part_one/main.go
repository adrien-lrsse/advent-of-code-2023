package part_one

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	i, j int
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func constructExtandedUniverse(input string) ([]Coord, []int, []int, int, int) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file :", err)
		return []Coord{}, []int{}, []int{}, 0, 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var hashtags []Coord
	var ligne []int
	ind := 0
	line_size := 0
	initilisation := false
	var colonneBool []bool
	for scanner.Scan() {
		line := scanner.Text()
		if !initilisation {
			colonneBool = make([]bool, len(line))
			initilisation = true
			line_size = len(line)
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
			ligne = append(ligne, ind)
		}

		ind++
	}

	var colonne []int
	for i := 0; i < len(colonneBool); i++ {
		if !colonneBool[i] {
			colonne = append(colonne, i)
		}
	}
	return hashtags, ligne, colonne, line_size, ind

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

func Part_One() int {
	hashtags, ligne, colonne, col_size, line_size := constructExtandedUniverse("input.txt")

	result := 0

	lineTreated := make([]int, line_size)
	emptyCount := 0
	checkeurEmpty := 0
	for i := 0; i < len(lineTreated); i++ {
		if checkeurEmpty < len(ligne) {
			if ligne[checkeurEmpty] == i {
				emptyCount++
				checkeurEmpty++
			}
		}
		lineTreated[i] = emptyCount
	}
	colonneTreated := make([]int, col_size)
	emptyCount = 0
	checkeurEmpty = 0
	for i := 0; i < len(colonneTreated); i++ {
		if checkeurEmpty < len(colonne) {
			if colonne[checkeurEmpty] == i {
				emptyCount++
				checkeurEmpty++
			}
		}
		colonneTreated[i] = emptyCount
	}

	for i := 0; i < len(hashtags); i++ {
		for j := i + 1; j < len(hashtags); j++ {
			//fmt.Println(coords[i], coords[j], findOptDistance(coords[i], coords[j], ligne, colonne))
			result += findOptDistance(hashtags[i], hashtags[j], lineTreated, colonneTreated)
		}
	}

	return result
}
