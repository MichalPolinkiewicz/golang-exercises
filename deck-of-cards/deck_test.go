package main

import (
	"strings"
	"testing"
)

func TestShouldExcludeExpectedCards(t *testing.T) {
	for _, v := range values {
		//given
		c := New()
		sizeBeforeExclude := len(c.cards)

		//when
		c.exclude(func(s *string) bool {
			if *s == v {
				return true
			}
			return false
		})

		//then
		sizeAfterExclude := len(c.cards)
		for _, r := range c.cards {
			if strings.Contains(r.value, v) || sizeAfterExclude != sizeBeforeExclude- 4 {
				t.Error("test failed", r.value, r.color)
			}
		}
	}
}

