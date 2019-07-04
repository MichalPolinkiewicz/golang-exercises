package main

import (
	"testing"
)

var words = map[string]int{
	"":                                  0,
	" ":                                 0,
	"word":                              1,
	"thisIs":                            2,
	"ThisIs":                            2,
	"A":                                 1,
	"A A":                               2,
	"AAA":                               3,
	"thisIsALongWordToCheckCorrectness": 8,
}

func TestShouldReturnCorrectNumberOfWords(t *testing.T) {
	for k, v := range words {
		result := countWordInString(&k)
		if *result != v {
			t.Errorf("value %d is not correct for '%s'. Correct value is %d", v, k, *result)
		}
	}
}
