package part_one

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	i int
	j int
}

func pipeAction(start Coord, actual Coord) (next Coord) {
	next.j = actual.j
	if start.i == actual.i-1 {
		next.i = actual.i + 1
	} else if start.i == actual.i+1 {
		next.i = actual.i - 1
	} else {
		next.i = -1
		next.j = -1
	}
	return next
}

func hyphenAction(start Coord, actual Coord) (next Coord) {
	next.i = actual.i
	if start.j == actual.j-1 {
		next.j = actual.j + 1
	} else if start.j == actual.j+1 {
		next.j = actual.j - 1
	} else {
		next.i = -1
		next.j = -1
	}
	return next
}

func LAction(start Coord, actual Coord) (next Coord) {
	if start.i == actual.i-1 {
		next.i = actual.i
		next.j = actual.j + 1
	} else if start.j == actual.j+1 {
		next.i = actual.i - 1
		next.j = actual.j
	} else {
		next.i = -1
		next.j = -1
	}
	return next
}

func JAction(start Coord, actual Coord) (next Coord) {
	if start.i == actual.i-1 {
		next.i = actual.i
		next.j = actual.j - 1
	} else if start.j == actual.j-1 {
		next.i = actual.i - 1
		next.j = actual.j
	} else {
		next.i = -1
		next.j = -1
	}
	return next
}

func SevenAction(start Coord, actual Coord) (next Coord) {
	if start.j == actual.j-1 {
		next.i = actual.i + 1
		next.j = actual.j
	} else if start.i == actual.i+1 {
		next.i = actual.i
		next.j = actual.j - 1
	} else {
		next.i = -1
		next.j = -1
	}
	return next
}

func FAction(start Coord, actual Coord) (next Coord) {
	if start.j == actual.j+1 {
		next.i = actual.i + 1
		next.j = actual.j
	} else if start.i == actual.i+1 {
		next.i = actual.i
		next.j = actual.j + 1
	} else {
		next.i = -1
		next.j = -1
	}
	return next
}

func addLineToMap(buildingMap [][]rune, line string) ([][]rune, Coord) {
	runes := []rune(line)
	start := Coord{-1, -1}
	for i := 0; i < len(runes); i++ {
		if runes[i] == 'S' {
			start.j = i
		}
	}
	start.i = len(buildingMap)
	return append(buildingMap, []rune(line)), start

}

func createMap(input string) ([][]rune, Coord) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return [][]rune{}, Coord{}
	}
	defer file.Close()
	var mapExtracted [][]rune
	start := Coord{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		temp := Coord{}
		mapExtracted, temp = addLineToMap(mapExtracted, scanner.Text())
		if temp.j != -1 && temp.i != -1 {
			start = temp
		}
	}
	return mapExtracted, start
}

func browseMap(start Coord, firstInstruction Coord, inputMap [][]rune) int {
	step_number := 1
	previous_coord := start
	actual_coord := firstInstruction
	for inputMap[actual_coord.i][actual_coord.j] != 'S' {
		switch inputMap[actual_coord.i][actual_coord.j] {
		case '|':
			tmp := actual_coord
			actual_coord = pipeAction(previous_coord, actual_coord)
			previous_coord = tmp
			step_number++
		case '-':
			tmp := actual_coord
			actual_coord = hyphenAction(previous_coord, actual_coord)
			previous_coord = tmp
			step_number++
		case 'L':
			tmp := actual_coord
			actual_coord = LAction(previous_coord, actual_coord)
			previous_coord = tmp
			step_number++

		case 'J':
			tmp := actual_coord
			actual_coord = JAction(previous_coord, actual_coord)
			previous_coord = tmp
			step_number++
		case '7':
			tmp := actual_coord
			actual_coord = SevenAction(previous_coord, actual_coord)
			previous_coord = tmp
			step_number++
		case 'F':
			tmp := actual_coord
			actual_coord = FAction(previous_coord, actual_coord)
			previous_coord = tmp
			step_number++
		default:
			step_number = -1
		}
		if step_number == -1 || actual_coord.i < 0 || actual_coord.j < 0 {
			actual_coord.i = start.i
			actual_coord.j = start.j
		}

	}
	return step_number

}
func findMax(list []int) int {
	max := list[0]
	for i := 0; i < len(list); i++ {
		if max < list[i] {
			max = list[i]
		}
	}
	return max
}

func Part_One() int {
	inputMap, start := createMap("input.txt")
	var loop_len []int
	loop_len = append(loop_len, browseMap(start, Coord{start.i + 1, start.j}, inputMap))
	loop_len = append(loop_len, browseMap(start, Coord{start.i - 1, start.j}, inputMap))
	loop_len = append(loop_len, browseMap(start, Coord{start.i, start.j + 1}, inputMap))
	loop_len = append(loop_len, browseMap(start, Coord{start.i, start.j - 1}, inputMap))

	return findMax(loop_len) / 2
}
