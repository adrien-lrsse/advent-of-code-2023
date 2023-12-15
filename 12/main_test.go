package main

import (
	"day12/part_one"
	"day12/part_two"
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
