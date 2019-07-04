package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	w := "thisIsWordForTests"
	r := countWordInString(&w)

	fmt.Println(*r)

}

func countWordInString(s *string) *int {
	counter := 0
	runeSlice := []rune(strings.Trim(*s, " "))

	if len(runeSlice) > 0 {
		counter++
		for i, r := range runeSlice {
			if unicode.IsUpper(r) && i != 0 {
				counter++
			}
		}
	}
	return &counter
}
