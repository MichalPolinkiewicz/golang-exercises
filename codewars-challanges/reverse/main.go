package main

import (
	"fmt"
	"strings"
)

//https://www.codewars.com/kata/5264d2b162488dc400000001
func main() {
	fmt.Println(SpinWords("this is a test for reversing wordsss"))
}

func SpinWords(str string) string {
	var out string
	spitted := strings.Split(str, " ")
	for _, e := range spitted {
		if len([]rune(e)) > 4 {
			out += reverse(e) + " "
		} else {
			out += e + " "
		}
	}
	return out[:len(str)]
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
