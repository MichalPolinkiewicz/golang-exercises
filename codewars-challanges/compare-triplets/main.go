package main

import "fmt"

func main() {
	a := []int32{17, 28, 30}
	b := []int32{99, 16, 8}

	fmt.Println(compareTriplets(a, b))

}

func compareTriplets(a []int32, b []int32) []int32 {
	if len(a) != len(b) {
		return nil
	}
	results := []int32{0, 0}
	for i, r := range a {
		if r > b[i] {
			results[0]++
		} else if b[i] > r {
			results[1]++
		}
	}
	return results
}
