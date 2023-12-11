package part_one

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

func constructExtandedUniverse(input string) [][]rune {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file :", err)
		return [][]rune{}
	}
	defer file.Close()
	var universeExtracted [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pattern := "^\\.+$"
		match, err := regexp.MatchString(pattern, line)
		if err != nil {
			// Gérer l'erreur si la syntaxe de l'expression régulière est incorrecte
			fmt.Println("Erreur dans l'expression régulière:", err)
			return [][]rune{}
		}
		if match {
			universeExtracted = append(universeExtracted, []rune(line))
			universeExtracted = append(universeExtracted, []rune(line))
		} else {
			universeExtracted = append(universeExtracted, []rune(line))
		}
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

	for k := 0; k < len(colonneAAjouter); k++ {
		for i := range universeExtracted {
			nouvelleLigne := make([]rune, len(universeExtracted[i])+1)
			copy(nouvelleLigne[:colonneAAjouter[k]+k+1], universeExtracted[i][:colonneAAjouter[k]+k])
			nouvelleLigne[colonneAAjouter[k]+k] = '.'
			copy(nouvelleLigne[colonneAAjouter[k]+k+1:], universeExtracted[i][colonneAAjouter[k]+k:])
			universeExtracted[i] = nouvelleLigne
		}
	}

	return universeExtracted
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

func findOptDistance(a Coord, b Coord) int {
	return int(math.Abs(float64(a.i-b.i))) + int(math.Abs(float64(a.j-b.j)))
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

func Part_One() int {
	universe := constructExtandedUniverse("input.txt")

	coords := findAllHastag(universe)

	coords = sortCoords(coords)

	result := 0

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			result += findOptDistance(coords[i], coords[j])
		}
	}

	return result
}
