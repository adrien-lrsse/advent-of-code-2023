package main

import (
	"day9/part_one"
	"day9/part_two"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Part One :")
	debut1 := time.Now()
	fmt.Println("Result : ", part_one.Part_One())
	fmt.Println("Temps d'exécution : ", time.Since(debut1))
	fmt.Println("Part Two :")
	debut2 := time.Now()
	fmt.Println("Result : ", part_two.Part_Two())
	fmt.Println("Temps d'exécution : ", time.Since(debut2))

}
