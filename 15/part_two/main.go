package part_two

import (
	"bytes"
)

func Part_Two(input []byte) int {
	var lenses [256]map[string]int
	var tokensOfLenses [256][]string

	// initialisation
	for i := range lenses {
		lenses[i] = make(map[string]int)
	}
	tokens := bytes.Split(input, []byte(","))
	res := 0
	for i := 0; i < len(tokens); i++ {
		current := 0
		label := true
		for j := 0; j < len(tokens[i]); j++ {
			if string(tokens[i][j]) == "=" || string(tokens[i][j]) == "-" {
				label = false
			}
			if label {
				current += int(tokens[i][j])
				current = current * 17
				current = current % 256
			}
		}
		isEquals := false
		intSpecial := 0
		for j := 0; j < len(tokens[i]); j++ {
			if string(tokens[i][j]) == "=" {
				isEquals = true
				intSpecial = j
			} else if string(tokens[i][j]) == "-" {
				intSpecial = j
			}
		}
		//fmt.Println(current)
		if isEquals {
			value := 0
			for j := intSpecial + 1; j < len(tokens[i]); j++ {
				value = value*10 + int(tokens[i][j]-'0')
			}
			if _, exists := lenses[current][string(tokens[i][:intSpecial])]; exists {
				lenses[current][string(tokens[i][:intSpecial])] = value
			} else {
				lenses[current][string(tokens[i][:intSpecial])] = value
				tokensOfLenses[current] = append(tokensOfLenses[current], string(tokens[i][:intSpecial]))
			}
			//fmt.Println(tokens[i], value)
		} else {
			if _, exists := lenses[current][string(tokens[i][:intSpecial])]; exists {
				ind := 0
				for k := 0; k < len(tokensOfLenses[current]); k++ {
					if tokensOfLenses[current][k] == string(tokens[i][:intSpecial]) {
						ind = k
					}
				}
				tokensOfLenses[current] = append(tokensOfLenses[current][:ind], tokensOfLenses[current][ind+1:]...)
			}
			delete(lenses[current], string(tokens[i][:intSpecial]))
		}
		//fmt.Println(string(tokens[i]), tokensOfLenses)
		//fmt.Println(lenses)
	}
	//fmt.Println(tokensOfLenses)
	//fmt.Println(lenses)
	for i := 0; i < len(tokensOfLenses); i++ {
		for j := 0; j < len(tokensOfLenses[i]); j++ {
			res += (i + 1) * (j + 1) * lenses[i][tokensOfLenses[i][j]]
		}
	}
	return res
}
