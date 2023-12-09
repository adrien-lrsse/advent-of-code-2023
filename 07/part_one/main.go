package part_one

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseLine(line string) []string {
	return strings.Split(line, ` `)
}

func typeOfHand(hand []rune) int {
	value := make(map[rune]int)
	for i := 0; i < len(hand); i++ {
		value[hand[i]] += 1
	}

	typeOfHand := 0
	nbrPaire := 0
	for _, v := range value {
		if v == 5 {
			typeOfHand = 5
			break
		}
		if v == 4 {
			typeOfHand = 4
			break
		}

		if v == 3 {
			typeOfHand = 3
		}
		if v == 2 && typeOfHand < v {
			typeOfHand = 2
		}
		if v == 1 && typeOfHand < v {
			typeOfHand = 1
		}
		if v == 2 {
			nbrPaire++
		}
	}

	handList := []int{typeOfHand, nbrPaire}

	switch {
	case handList[0] == 5:
		return 8
	case handList[0] == 4:
		return 7
	case handList[0] == 3 && handList[1] == 1:
		return 6
	case handList[0] == 3:
		return 5
	case handList[0] == 2 && handList[1] == 2:
		return 4
	case handList[0] == 2:
		return 3
	case handList[0] == 1:
		return 2
	default:
		return 1
	}
}

func valueOfHand(typeOfHand int, hand []rune) float64 {
	cardValues := map[rune]float64{
		'A': 24,
		'K': 23,
		'Q': 22,
		'J': 21,
		'T': 20,
		'9': 19,
		'8': 18,
		'7': 17,
		'6': 16,
		'5': 15,
		'4': 14,
		'3': 13,
		'2': 12,
	}
	sumOfCardValue := 0.0
	for i := 0; i < len(hand); i++ {
		sumOfCardValue += cardValues[hand[i]] * (math.Pow(10, float64(10-2*i)))
	}

	return float64(typeOfHand)*math.Pow(10, 12) + sumOfCardValue
}

func getValue(perm []int, prize []int) int {
	result := 0
	for i := 0; i < len(perm); i++ {
		result += (i + 1) * prize[perm[i]]
	}
	return result
}

func getPerm(list []float64) []int {
	indices := make([]int, len(list))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return list[indices[i]] < list[indices[j]]
	})
	return indices
}

func Part_One() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}

	valueTot := 0
	defer file.Close()
	var price []int
	var representationofHand []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		extractation := parseLine(line)
		num, err := strconv.Atoi(extractation[1])
		valueTot += num

		// Vérification des erreurs
		if err != nil {
			fmt.Println("Erreur de conversion :", err)
			return 0
		}
		price = append(price, num)
		representationofHand = append(representationofHand, valueOfHand(typeOfHand([]rune(extractation[0])), []rune(extractation[0])))
	}
	perm := getPerm(representationofHand)
	return getValue(perm, price)
}
