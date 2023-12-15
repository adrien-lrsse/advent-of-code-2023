package main

import (
	"day15/part_one"
	"day15/part_two"
	"testing"
)

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part_one.Part_One(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part_two.Part_Two(input)
	}
}
