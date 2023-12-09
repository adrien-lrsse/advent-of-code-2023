package part_two

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func parseSeeds(line string) []int {
	var seeds []int
	line_decoupe := strings.Split(line, `:`)

	runes := []rune(line_decoupe[1])

	tempon := -1

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) && tempon == -1 {
			tempon = int(runes[i] - '0')
		} else if unicode.IsDigit(runes[i]) {
			tempon = tempon*10 + int(runes[i]-'0')
		}
		if !(unicode.IsDigit(runes[i])) && (tempon != -1) {
			seeds = append(seeds, tempon)
			tempon = -1
		}
	}

	if tempon != -1 {
		seeds = append(seeds, tempon)

	}
	return seeds
}

func parseMap(line string) []int {
	var maps []int

	runes := []rune(line)

	tempon := -1

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) && tempon == -1 {
			tempon = int(runes[i] - '0')
		} else if unicode.IsDigit(runes[i]) {
			tempon = tempon*10 + int(runes[i]-'0')
		}
		if !(unicode.IsDigit(runes[i])) && (tempon != -1) {
			maps = append(maps, tempon)
			tempon = -1
		}
	}

	if tempon != -1 {
		maps = append(maps, tempon)
	}

	return maps
}

func findMin(list [][]int) int {
	min := list[0][0]
	for i := 1; i < len(list); i++ {
		if min > list[i][0] {
			min = list[i][0]
		}
	}
	return min
}

func findValueOfIntervals(destination int, source int, length int, interval []int) ([][]int, bool) {
	var intervalDest [][]int
	intervalSource := []int{source, source + length - 1}

	isChange := false

	switch {
	case interval[0] >= intervalSource[0] && interval[1] <= intervalSource[1]: // cas 1
		intervalDest = append(intervalDest, []int{destination + (interval[0] - intervalSource[0]), destination + (interval[1] - intervalSource[0])})
		isChange = true
	case interval[0] < intervalSource[0] && interval[1] <= intervalSource[1] && interval[1] >= intervalSource[0]: // cas 2
		intervalDest = append(intervalDest, []int{destination, destination + (interval[1] - intervalSource[0])})
		intervalDest = append(intervalDest, []int{interval[0], intervalSource[0] - 1})
		isChange = true
	case interval[0] >= intervalSource[0] && interval[0] <= intervalSource[1] && interval[1] > intervalSource[1]: // cas 3
		intervalDest = append(intervalDest, []int{destination + (interval[0] - intervalSource[0]), destination + length - 1})
		intervalDest = append(intervalDest, []int{intervalSource[1] + 1, interval[1]})
		isChange = true
	case interval[0] < intervalSource[0] && interval[1] > intervalSource[1]:
		intervalDest = append(intervalDest, []int{destination, destination + (intervalSource[1] - intervalSource[0])})
		intervalDest = append(intervalDest, []int{interval[0], intervalSource[0] - 1})
		intervalDest = append(intervalDest, []int{intervalSource[1] + 1, interval[1]})
		isChange = true
	default:
		intervalDest = append(intervalDest, interval)
	}

	return intervalDest, isChange
}

func createSeedsInterval(list []int) [][]int {
	var seedsInterval [][]int
	for i := 0; i < len(list)-1; i += 2 {
		seedsInterval = append(seedsInterval, []int{list[i], list[i] + list[i+1] - 1})
	}
	return seedsInterval
}

func Part_Two() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}
	defer file.Close()
	var interval_from_map [][]int

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	extracted_seeds := parseSeeds(line)

	seeds_intervals := createSeedsInterval(extracted_seeds)

	var isCheck []bool
	for i := 0; i < len(seeds_intervals); i++ {
		isCheck = append(isCheck, false)
	}
	for scanner.Scan() {

		line = scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			isCheck = []bool{}
			for i := 0; i < len(seeds_intervals); i++ {
				isCheck = append(isCheck, false)
			}
			continue
		} else if !unicode.IsDigit(rune(line[0])) {
			isCheck = []bool{}
			for i := 0; i < len(seeds_intervals); i++ {
				isCheck = append(isCheck, false)
			}
			continue
		} else {
			map_i := parseMap(line)

			if len(seeds_intervals) == 0 {
				continue
			}
			k := 0
			for {
				if k < len(seeds_intervals) {
					if !isCheck[k] {
						interval_from_map, isCheck[k] = findValueOfIntervals(map_i[0], map_i[1], map_i[2], seeds_intervals[k])
						seeds_intervals[k] = interval_from_map[0]
						for j := 1; j < len(interval_from_map); j++ {
							seeds_intervals = append(seeds_intervals, interval_from_map[j])
							isCheck = append(isCheck, false)
						}
					}
					k++
				} else {
					break
				}

			}
		}
	}

	return findMin(seeds_intervals)
}
