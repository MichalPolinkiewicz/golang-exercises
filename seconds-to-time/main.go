package main

import (
	"strconv"
	"strings"
)

func main() {

	convertToHumanReadable(60)
}

func convertToHumanReadable(seconds int) string {
	const (
		secondsInHour   = 3600
		secondsInMinute = 60
	)
	var out []string

	//helper func's
	addZero := func() {
		out = append(out, "0")
	}

	addColon := func() {
		out = append(out, ":")
	}

	addValue := func(v int) {
		out = append(out, strconv.Itoa(v))
	}

	//calc hours
	h := seconds / secondsInHour
	seconds = seconds - secondsInHour*h
	if h < 10 {
		addZero()
	}
	addValue(h)
	addColon()

	//calc minutes
	m := seconds / secondsInMinute //59
	if m < 10 {
		addZero()
	}
	addValue(m)
	addColon()

	//calc seconds
	s := seconds - 60*m
	if s < 10 {
		addZero()
	}
	addValue(s)

	return strings.Join(out, "")
}
