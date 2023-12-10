package main

import (
	"day10/part_one"
	"day10/part_two"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Part One :")
	debut := time.Now()
	fmt.Println("Result : ", part_one.Part_One(), " (in", time.Since(debut), ")")
	fmt.Println("Part Two :")
	debut = time.Now()
	fmt.Println("Result : ", part_two.Part_Two(), "(in", time.Since(debut), ")")

}
