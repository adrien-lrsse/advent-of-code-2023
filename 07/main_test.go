package main

import (
	"day7/part_one"
	"day7/part_two"
	"testing"
)

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part_one.Part_One()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part_two.Part_Two()
	}
}
