package part_one

import (
	"bytes"
)

func Part_One(input []byte) int {
	tokens := bytes.Split(input, []byte(","))
	res := 0
	for i := 0; i < len(tokens); i++ {
		current := 0
		for j := 0; j < len(tokens[i]); j++ {
			current += int(tokens[i][j])
			current = current * 17
			current = current % 256
		}
		res += current
	}
	return res
}
