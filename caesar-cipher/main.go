package main

import "strings"

var (
	alphabet = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'r', 's', 't', 'u', 'w', 'y', 'z'}
	shuffled []rune
)

func main() {
	shuffleAlphabet(3)
	encrypt("michal")
}

func shuffleAlphabet(n int) {
	shuffled = make([]rune, len(alphabet))

	x := 0
	for i := range alphabet {
		if i+n < len(alphabet) {
			shuffled[i] = alphabet[i+n]
		} else {
			shuffled[i] = alphabet[x]
			x++
		}
	}
}

func encrypt(s string) string {
	toLower := strings.ToLower(s)
	encrypted := make([]rune, len(s))
	runsSlice := []rune(toLower)

	for i, r := range runsSlice {
		encrypted[i] = ' '
		if r != ' ' {
			encrypted[i] = shuffled[getIndexOf(&r)]
		}
	}

	return strings.Trim(string(encrypted), " ")
}

func getIndexOf(r *rune) int {
	for k, v := range alphabet {
		if *r == v {
			return k
		}
	}
	return -1
}
