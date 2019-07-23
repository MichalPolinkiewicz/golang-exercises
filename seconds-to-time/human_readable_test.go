package main

import "testing"

func TestMakeHumanReadable(t *testing.T){
	//given
	inputSeconds := []int {0, 5, 60, 86399, 359999}
	expectedResults := []string{"00:00:00", "00:00:05", "00:01:00", "23:59:59", "99:59:59"}

	//when && then
	for i, r := range inputSeconds {
		result := convertToHumanReadable(r)
		if result != expectedResults[i] {
			t.Errorf("should be %s, is %s", expectedResults[i], result )
		}
	}
}
