package part_two

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

type Coord struct {
	i, j int
}

func constructExtandedUniverse(input string) ([][]rune, []int, []int) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file :", err)
		return [][]rune{}, []int{}, []int{}
	}
	defer file.Close()
	var universeExtracted [][]rune
	scanner := bufio.NewScanner(file)

	var ligneAjoute []int
	ind := 0
	for scanner.Scan() {
		line := scanner.Text()
		pattern := "^\\.+$"
		match, err := regexp.MatchString(pattern, line)
		if err != nil {
			// Gérer l'erreur si la syntaxe de l'expression régulière est incorrecte
			fmt.Println("Erreur dans l'expression régulière:", err)
			return [][]rune{}, []int{}, []int{}
		}
		if match {
			universeExtracted = append(universeExtracted, []rune(line))
			ligneAjoute = append(ligneAjoute, ind)
		} else {
			universeExtracted = append(universeExtracted, []rune(line))
		}
		ind++
	}

	allColonneareStars := true
	var colonneAAjouter []int
	for i := 0; i < len(universeExtracted[0]); i++ {
		for j := 0; j < len(universeExtracted); j++ {
			if universeExtracted[j][i] != '.' {
				allColonneareStars = false
			}
		}
		if allColonneareStars {
			colonneAAjouter = append(colonneAAjouter, i)
		} else {
			allColonneareStars = true
		}
	}

	return universeExtracted, ligneAjoute, colonneAAjouter
}

func findAllHastag(universe [][]rune) []Coord {
	var result []Coord
	for i := 0; i < len(universe); i++ {
		for j := 0; j < len(universe[i]); j++ {
			if universe[i][j] == '#' {
				result = append(result, Coord{i, j})
			}
		}
	}
	return result
}

func findOptDistance(a Coord, b Coord, ligne []int, colonne []int) int {
	return int(math.Abs(float64(a.i-b.i))) + quantityOfLine(a, b, ligne)*999999 + int(math.Abs(float64(a.j-b.j))) + quantityOfColonne(a, b, colonne)*999999
}

func sortCoords(coords []Coord) []Coord {
	n := len(coords)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if coords[j].i > coords[j+1].i || (coords[j].i == coords[j+1].i && coords[j].j > coords[j+1].j) {
				// Échange des éléments si dans le mauvais ordre
				coords[j], coords[j+1] = coords[j+1], coords[j]
			}
		}
	}

	return coords
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
	universe, ligne, colonne := constructExtandedUniverse("input.txt")

	coords := findAllHastag(universe)

	coords = sortCoords(coords)

	result := 0

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			//fmt.Println(coords[i], coords[j], findOptDistance(coords[i], coords[j], ligne, colonne))
			result += findOptDistance(coords[i], coords[j], ligne, colonne)
		}
	}

	return result
}
