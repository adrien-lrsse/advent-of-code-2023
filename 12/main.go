package main

import (
	"day12/part_one"
	"day12/part_two"
	"fmt"
	"time"
)

func main() {
	debut := time.Now()
	fmt.Println("Part One :")
	fmt.Println("Result : ", part_one.Part_One(input), " (in", time.Since(debut), ")")
	debut = time.Now()
	fmt.Println("Part Two :")
	fmt.Println("Result : ", part_two.Part_Two(inputTest), "(in", time.Since(debut), ")")

}
