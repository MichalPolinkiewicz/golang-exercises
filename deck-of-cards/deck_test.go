package main

import (
	"strings"
	"testing"
)

func TestShouldExcludeExpectedCards(t *testing.T) {
	for _, v := range values {
		//given
		c := New()

		//when
		c.exclude(func(s *string) bool {
			if *s == v {
				return true
			}
			return false
		})

		//then
		for _, r := range c.cards {
			if strings.Contains(r.value, v) {
				t.Error("test failed", r.value, r.color)
			}
		}
	}
}

