package main

import (
	"fmt"
)

func main() {
	a := [][]int32 {
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}

	fmt.Println(calc(a))

}

func calc(arr [][]int32) int32 {
	var results []int32
	var sum int32 = 0
	for i:= 0; i+2 < len(arr); i++ {
			arr1 := arr[i]
			arr2 := arr[i+1]
			arr3 := arr[i+2]
			for j:= 0; j+2 < len(arr[j]); j++ {
				sum = arr1[j] + arr1[j+1] + arr1[j+2] + arr2[j+1] + arr3[j] + arr3[j+1] + arr3[j+2]
				results = append(results, sum)
		}
	}

	max := results[0]
	for _, e := range results {
		if e > max {
			max = e
		}
	}
	return max
}
