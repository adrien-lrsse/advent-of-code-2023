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

func findValueOfIntervals(destination int, source int, length int, interval []int) [][]int {
	var intervalDest [][]int
	intervalSource := []int{source, source + length - 1}

	switch {
	case intervalSource[0] <= interval[0] && intervalSource[1] >= interval[1]: // milieu et milieu
		newInterval := []int{destination + (interval[0] - source), destination + (interval[1] - source)}
		intervalDest = append(intervalDest, newInterval)

	case intervalSource[0] > interval[0] && intervalSource[1] < interval[1]: // avant et après
		intervalDest = append(intervalDest, []int{interval[0], intervalSource[0] - 1})
		intervalDest = append(intervalDest, []int{destination, destination + length - 1})
		intervalDest = append(intervalDest, []int{intervalSource[1] + 1, interval[1]})

	case intervalSource[0] > interval[0] && interval[1] >= intervalSource[0] && interval[1] <= intervalSource[1]: // avant et mileu
		newInterval := []int{destination, destination + (interval[1] - source)}
		nonModifie := []int{interval[0], intervalSource[0] - 1}
		intervalDest = append(intervalDest, newInterval, nonModifie)

	case interval[0] >= intervalSource[0] && interval[0] <= intervalSource[1] && interval[1] > intervalSource[1]: // milieu et aprs
		newInterval := []int{destination + (interval[0] - source), destination + length - 1}
		nonModifie := []int{intervalSource[1] + 1, interval[1]}
		intervalDest = append(intervalDest, newInterval, nonModifie)

	default:
		intervalDest = append(intervalDest, interval)
	}

	return intervalDest
}

func Part_Two() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	extracted_seeds := parseSeeds(line)

	var seeds_intervals [][]int

	var interval1 []int
	interval1 = append(interval1, extracted_seeds[0])
	interval1 = append(interval1, extracted_seeds[0]+extracted_seeds[1]-1)
	var interval2 []int
	interval2 = append(interval2, extracted_seeds[2])
	interval2 = append(interval2, extracted_seeds[2]+extracted_seeds[3]-1)

	seeds_intervals = append(seeds_intervals, interval1)
	seeds_intervals = append(seeds_intervals, interval2)

	var isCheck []bool

	len_inter := len(seeds_intervals)

	isCheck = make([]bool, len_inter)

	for scanner.Scan() {
		line = scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			len_inter = len(seeds_intervals)
			isCheck = make([]bool, len_inter)
			continue
		} else if !unicode.IsDigit(rune(line[0])) {
			len_inter = len(seeds_intervals)
			isCheck = make([]bool, len_inter)
			continue
		} else {
			map_i := parseMap(line)
			m := len(seeds_intervals) + len(seeds_intervals)
			for i := 0; i < m; i++ {
				fmt.Println(i, m)
				if i < len(seeds_intervals) {
					fmt.Println(line, len(seeds_intervals))
					if !isCheck[i] {
						tabInter := findValueOfIntervals(map_i[0], map_i[1], map_i[2], seeds_intervals[i])
						fmt.Println(tabInter)

						if seeds_intervals[i][0] != tabInter[0][0] || seeds_intervals[i][1] != tabInter[0][1] {
							isCheck[i] = true
							seeds_intervals[i] = tabInter[0]
							for j := 1; j < len(tabInter); j++ {
								seeds_intervals = append(seeds_intervals, tabInter[j])
								isCheck = append(isCheck, false)
								m = len(seeds_intervals) + len(seeds_intervals)
							}
						}
					}

				} else {
					break
				}

			}

		}
	}

	var test []int
	test = append(test, 49)
	test = append(test, 96)

	fmt.Println(findValueOfIntervals(40, 50, 20, test))

	fmt.Println(seeds_intervals)

	return findMin(seeds_intervals)
}
