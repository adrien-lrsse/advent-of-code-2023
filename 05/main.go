package main

import (
	"day5/part_one"
	"day5/part_two"
	"fmt"
	"time"
)

func main() {
	debut := time.Now()
	fmt.Println("Part One :")
	fmt.Println("Result : ", part_one.Part_One(), " (in", time.Since(debut), ")")
	debut = time.Now()
	fmt.Println("Part Two :")
	fmt.Println("Result : ", part_two.Part_Two(), "(in", time.Since(debut), ")")

}
