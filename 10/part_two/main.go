package part_two

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func browseMap(start Coord, firstInstruction Coord, inputMap [][]rune) (int, []Coord) {
	step_number := 1
	previous_coord := start
	actual_coord := firstInstruction
	list_Coord := []Coord{firstInstruction}
	for inputMap[actual_coord.i][actual_coord.j] != 'S' {
		switch inputMap[actual_coord.i][actual_coord.j] {
		case '|':
			tmp := actual_coord
			actual_coord = pipeAction(previous_coord, actual_coord)
			previous_coord = tmp
			list_Coord = append(list_Coord, actual_coord)
			step_number++
		case '-':
			tmp := actual_coord
			actual_coord = hyphenAction(previous_coord, actual_coord)
			list_Coord = append(list_Coord, actual_coord)
			previous_coord = tmp
			step_number++
		case 'L':
			tmp := actual_coord
			actual_coord = LAction(previous_coord, actual_coord)
			list_Coord = append(list_Coord, actual_coord)
			previous_coord = tmp
			step_number++

		case 'J':
			tmp := actual_coord
			actual_coord = JAction(previous_coord, actual_coord)
			list_Coord = append(list_Coord, actual_coord)
			previous_coord = tmp
			step_number++
		case '7':
			tmp := actual_coord
			actual_coord = SevenAction(previous_coord, actual_coord)
			list_Coord = append(list_Coord, actual_coord)
			previous_coord = tmp
			step_number++
		case 'F':
			tmp := actual_coord
			actual_coord = FAction(previous_coord, actual_coord)
			list_Coord = append(list_Coord, actual_coord)
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
	return step_number, list_Coord

}

func findMax(list []int) int {
	max := list[0]
	max_i := 0
	for i := 0; i < len(list); i++ {
		if max < list[i] {
			max = list[i]
			max_i = i
		}
	}
	return max_i
}

func CoordIn(item Coord, list []Coord) bool {
	for i := 0; i < len(list); i++ {
		if item == list[i] {
			return true
		}

	}
	return false
}

func Part_Two() int {
	inputMap, start := createMap("input.txt")
	var loop_len []int
	res1, list1 := browseMap(start, Coord{start.i + 1, start.j}, inputMap)
	res2, list2 := browseMap(start, Coord{start.i - 1, start.j}, inputMap)
	res3, list3 := browseMap(start, Coord{start.i, start.j + 1}, inputMap)
	res4, list4 := browseMap(start, Coord{start.i, start.j - 1}, inputMap)

	loop_len = append(loop_len, res1)
	loop_len = append(loop_len, res2)
	loop_len = append(loop_len, res3)
	loop_len = append(loop_len, res4)

	var loop_Coord [][]Coord

	loop_Coord = append(loop_Coord, list1)
	loop_Coord = append(loop_Coord, list2)
	loop_Coord = append(loop_Coord, list3)
	loop_Coord = append(loop_Coord, list4)

	h := findMax(loop_len)

	result := 0
	for i := 0; i < len(inputMap); i++ { // thx to https://github.com/Nounoursdestavernes/AdventOfCode2023/blob/main/day10/Part2/part2.py
		tmp := []rune{}
		for j := 0; j < len(inputMap[i]); j++ {
			if CoordIn(Coord{i, j}, loop_Coord[h]) {
				tmp = append(tmp, inputMap[i][j])
			} else {
				tmp = append(tmp, '.')
			}

		}
		re := regexp.MustCompile(`L-*7|F-*J`)
		str := string(tmp)
		str = re.ReplaceAllString(str, "|")

		inside := false
		for _, char := range str {
			if char == '|' {
				inside = !inside
			}
			if char == '.' && inside {
				result++
			}
		}

	}

	return result
}
