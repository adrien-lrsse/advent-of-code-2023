package main

import (
	"day15/part_one"
	"day15/part_two"
	"fmt"
	"time"
)

func main() {
	debut := time.Now()
	fmt.Println("Part One :")
	fmt.Println("Result : ", part_one.Part_One(inputTest), " (in", time.Since(debut), ")")
	debut = time.Now()
	fmt.Println("Part Two :")
	fmt.Println("Result : ", part_two.Part_Two(input), "(in", time.Since(debut), ")")

}
