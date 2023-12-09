package main

import (
	"day9/part_one"
	"day9/part_two"
	"fmt"
	"time"
)

func main() {
	debut := time.Now()
	fmt.Println("Part One :")
	fmt.Println("Temps d'exécution : ", time.Since(debut))
	fmt.Println("Result : ", part_one.Part_One())
	debut = time.Now()
	fmt.Println("Part Two :")
	fmt.Println("Result : ", part_two.Part_Two())
	fmt.Println("Temps d'exécution : ", time.Since(debut))

}
