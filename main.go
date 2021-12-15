package main

import (
	"fmt"
	"constraints"
)

type Number interface {
	constraints.Integer | constraints.Float	
}

func main() {
	ints := map[string]int64{
		"1st": 32,
		"2nd": 12,
	}

	floats := map[string]float64{
		"1st": 32.13,
		"2nd": 12.15,
	}

	fmt.Printf("Non-generic sums: %v and %v.\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic sums: %v and %v.\n",
		SumGeneric(ints),
		SumGeneric(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

func SumGeneric[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}