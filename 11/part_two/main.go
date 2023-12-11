package part_two

import (
	"bufio"
	"fmt"
	"math"
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
	return hashtags, ligne, colonne

}

func findOptDistance(a Coord, b Coord, ligne []int, colonne []int) int {
	return int(math.Abs(float64(a.i-b.i))) + quantityOfLine(a, b, ligne)*999999 + int(math.Abs(float64(a.j-b.j))) + quantityOfColonne(a, b, colonne)*999999
}

func quantityOfLine(a Coord, b Coord, list []int) int {
	res := 0
	if a.i < b.i {
		for i := 0; i < len(list); i++ {
			if a.i < list[i] && b.i > list[i] {
				res++
			}
		}
		//fmt.Println("Ligne : ", res)
		return res
	} else if a.i > b.i {
		for i := 0; i < len(list); i++ {
			if b.i < list[i] && a.i > list[i] {
				res++
			}
		}
		//fmt.Println("Ligne : ", res)
		return res
	} else {
		return 0
	}

}

func quantityOfColonne(a Coord, b Coord, list []int) int {
	res := 0
	if a.j < b.j {
		for i := 0; i < len(list); i++ {
			if a.j < list[i] && b.j > list[i] {
				res++
			}
		}
		//fmt.Println("Colonne : ", res)
		return res
	} else if a.j > b.j {
		for i := 0; i < len(list); i++ {
			if b.j < list[i] && a.j > list[i] {
				res++
			}
		}
		//fmt.Println("Colonne : ", res)
		return res
	} else {
		return 0
	}
}

func Part_Two() int {
	hashtags, ligne, colonne := constructExtandedUniverse("input.txt")

	result := 0

	for i := 0; i < len(hashtags); i++ {
		for j := i + 1; j < len(hashtags); j++ {
			//fmt.Println(coords[i], coords[j], findOptDistance(coords[i], coords[j], ligne, colonne))
			result += findOptDistance(hashtags[i], hashtags[j], ligne, colonne)
		}
	}

	return result
}
