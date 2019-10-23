package main

import (
	"fmt"
)

func main() {

	a := []int{1,2,3,6,8}
	//sort.Sort(sort.Reverse(sort.IntSlice(a)))
	//fmt.Println(a)

	a = RaKO(a, 2)
	fmt.Println(a) //1, 2, 6, 8

	a = RwKO(a, 2)
	fmt.Println(a) // 1, 2, 8
}

//remove and keep keep order
func RaKO(s []int, i int) []int {
	copy(s[i:], s[i+1:])

	return s[:len(s)-1]
}

//remove without keeping order
func RwKO(s []int, i int)[]int {
	s[i] = s[len(s)-1]

	return s[:len(s)-1]
}