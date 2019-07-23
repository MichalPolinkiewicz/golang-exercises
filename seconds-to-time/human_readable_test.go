package main

import "testing"

func TestMakeHumanReadable(t *testing.T) {
	//given
	testCases := map[int]string{
		0:      "00:00:00",
		5:      "00:00:05",
		60:     "00:01:00",
		86399:  "23:59:59",
		359999: "99:59:59",
	}

	//when && then
	for k := range testCases {
		result := convertToHumanReadable(k)
		if result != testCases[k] {
			t.Errorf("should be %s, is %s", testCases[k], result)
		}
	}
}
