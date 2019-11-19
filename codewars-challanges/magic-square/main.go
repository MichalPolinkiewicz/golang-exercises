package main

import (
	"fmt"
	"math"
)

func main() {
	//in := [][]int32{
	//	{4, 9, 2},
	//	{3, 5, 7},
	//	{8, 1, 5},
	//}
	//
	in2 := [][]int32{
		{4, 8, 2},
		{4, 5, 7},
		{6, 1, 6},
	}

	//in3 := [][]int32 {
	//	{4, 5, 8},
	//	{2, 4, 1},
	//	{1, 9, 7},
	//}

	fmt.Println(formingMagicSquare(in2))

}

func formingMagicSquare(in [][]int32) int32 {
	magicSquares := make([][]int, 8)
	magicSquares[0] = []int{4, 9, 2, 3, 5, 7, 8, 1, 6}
	magicSquares[1] = []int{8, 1, 6, 3, 5, 7, 4, 9, 2}
	magicSquares[2] = []int{2, 9, 4, 7, 5, 3, 6, 1, 8}
	magicSquares[3] = []int{6, 1, 8, 7, 5, 3, 2, 9, 4}
	magicSquares[4] = []int{8, 3, 4, 1, 5, 9, 6, 7, 2}
	magicSquares[5] = []int{6, 7, 2, 1, 5, 9, 8, 3, 4}
	magicSquares[6] = []int{4, 3, 8, 9, 5, 1, 2, 7, 6}
	magicSquares[7] = []int{2, 7, 6, 9, 5, 1, 4, 3, 8}

	flatted := flat(in)
	var results []int32
	//comparing all magic squares with flatted input
	for _, r := range magicSquares {
		compareResults := compareSlices(r, flatted)
		//if result == 0 it is magic square; end
		if compareResults == 0 {
			return compareResults
		}
		results = append(results, compareResults)
	}

	result := results[0]
	for _, r := range results {
		if r < result {
			result = r
		}
	}
	return result
}

func compareSlices(e1, e2 []int) int32 {
	//searching for differences between slices
	var results []float64
	for i, e := range e1 {
		if e != e2[i] {
			results = append(results, math.Abs(float64(e-e2[i])))
		}
	}

	//sum
	var result int32
	for _, r := range results {
		result += int32(r)
	}
	return result
}

func flat(s [][]int32) []int {
	var out []int
	for _, r := range s {
		for _, e := range r {
			out = append(out, int(e))
		}
	}
	return out
}
