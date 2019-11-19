package main

import "fmt"

//https://www.hackerrank.com/challenges/new-year-chaos/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays

func main() {
	a := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	minimumBribes(a) //7
}

func minimumBribes(q []int32) {
	var counter int32

	for i, e := range q {
		if i+1 < len(q) && e > q[i+1]+2 {
			fmt.Println("Too chaotic")
			return
		}

		if e != int32(i+1) {
		counter++
		}

	}
	fmt.Println(counter)

}
