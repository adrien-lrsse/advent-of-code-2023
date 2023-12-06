package part_two

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func discriminant(a int, b int, c int) int {
	return b*b - 4*a*c
}

func solv_Pol_Deg2(a int, b int, c int) []float64 {
	var sol []float64
	delta := discriminant(a, b, c)
	if delta == 0 {
		unique := -(float64(b)) / (2 * float64(a))
		sol = append(sol, unique)
		return sol
	} else if delta > 0 {
		x_1 := -(float64(b) + math.Sqrt(float64(delta))) / (2 * float64(a))
		x_2 := -(float64(b) - math.Sqrt(float64(delta))) / (2 * float64(a))
		sol = append(sol, x_1)
		sol = append(sol, x_2)
		return sol
	}
	return sol
}

func count_sol(sol_pol []float64) int {
	if len(sol_pol) == 1 {
		return 1
	} else if len(sol_pol) == 2 {
		return int(int64(math.Floor(sol_pol[1])) - int64(math.Floor(sol_pol[0])))
	}
	return 0
}

func parseRace(line_temps string, line_distance string) ([]int, []int) {
	var tps_course []int
	var d_course []int

	decoupe_tps_cours := strings.Split(line_temps, `:`)
	decoupe_distance_course := strings.Split(line_distance, `:`)

	runesTemps := []rune(decoupe_tps_cours[1])
	runesDistance := []rune(decoupe_distance_course[1])

	tempon := 0

	for i := 0; i < len(runesTemps); i++ {
		if unicode.IsDigit(runesTemps[i]) {
			tempon = tempon*10 + int(runesTemps[i]-'0')
		}
	}

	if tempon != 0 {
		tps_course = append(tps_course, tempon)
	}

	tempon = 0

	for i := 0; i < len(runesDistance); i++ {
		if unicode.IsDigit(runesDistance[i]) {
			tempon = tempon*10 + int(runesDistance[i]-'0')
		}
	}

	if tempon != 0 {
		d_course = append(d_course, tempon)
	}

	return tps_course, d_course
}

func Part_Two() int {

	file, err := os.Open("part_two/input.txt")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return 0
	}
	defer file.Close()

	result := 1

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	temps := scanner.Text()
	scanner.Scan()
	distance := scanner.Text()
	tpsSlice, distanceSlice := parseRace(temps, distance)
	for i := 0; i < len(tpsSlice); i++ {
		x := solv_Pol_Deg2(1, -tpsSlice[i], distanceSlice[i])
		result *= count_sol(x)
	}
	return result
}
