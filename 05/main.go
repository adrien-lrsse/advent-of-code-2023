package main

import (
	"day5/part_one"
	"day5/part_two"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Part One :")
	debut := time.Now()
	fmt.Println("Temps d'exécution : ", time.Since(debut))
	fmt.Println("Result : ", part_one.Part_One())
	fmt.Println("Part Two :")
	debut = time.Now()
	fmt.Println("Result : ", part_two.Part_Two())
	fmt.Println("Temps d'exécution : ", time.Since(debut))

}
