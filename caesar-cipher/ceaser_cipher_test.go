package main

import (
	"testing"
)

func TestShouldCorrectlyShuffleAlphabet(t *testing.T) {
	type shuffledRune struct {
		key        int
		checkIndex int
		encrypted  rune
	}

	shuffledRunes := []shuffledRune{
		{0, 0, 'a'},
		{0, 6, 'g'},
		{0, 20, 'w'},

		{1, 0, 'b'},
		{1, 6, 'h'},
		{1, 20, 'y'},

		{7, 0, 'h'},
		{7, 6, 'n'},
		{7, 20, 'e'},

		{10, 0, 'k'},
		{10, 6, 'r'},
		{10, 20, 'h'},
	}

	for _, v := range shuffledRunes {
		shuffleAlphabet(v.key)
		if shuffled[v.checkIndex] != v.encrypted {
			t.Errorf("key= %d, checkIndex= %d, result should be %s, is %s", v.key, v.checkIndex, string(v.encrypted), string(shuffled[v.checkIndex]))
		}
	}
}

func TestShouldCorrectlyEncryptGivenWord(t *testing.T) {
	type shuffledWord struct {
		key int
		originalWord string
		encrypted string
	}

	words := []shuffledWord{
		{1, " ", ""},
		{1, "", ""},
		{2, "hello", "jgnnr"},
		{4, "john", "ntls"},
		{4, " JO H N ", "nt l s"},
	}

	for _, v := range words {
		shuffleAlphabet(v.key)
		result := encrypt(v.originalWord)

		if result != v.encrypted {
			t.Errorf("%s should be encryped as %s, is %s", v.originalWord, v.encrypted, result)
		}
	}
}
