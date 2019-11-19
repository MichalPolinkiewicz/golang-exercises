package main

import "fmt"

func main() {
	a := []int32{1,2,3,4,5}
	fmt.Println(rotLeft(a, 4))
}


func rotLeft(a []int32, d int32) []int32 {
	x := a[d:]
	y := a[:d]
	r := append(x, y...)

	return r
}
