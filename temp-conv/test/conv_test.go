package test

import (
	tempconv "github.com/MichalPolinkiewicz/golang-exercises/temp-conv"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	//given
	var testCases = map[tempconv.Celsius]tempconv.Fahrenheit{
		-39: -38.2,
		-1:  30.2,
		0:   32,
		9:   48.2,
		39:  102.2,
	}

	//when && then
	for k, v := range testCases {
		result := tempconv.CToF(k)
		if int(result) != int(v) {
			t.Fatalf("should be %f, got %f", v, result)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	//given
	var testCases = map[tempconv.Celsius]tempconv.Kelvin{
		-39: 234.15,
		-1:  272.15,
		0:   273.15,
		9:   282.15,
		39:  312.5,
	}

	//when && then
	for k, v := range testCases {
		result := tempconv.CToK(k)
		if int(result) != int(v) {
			t.Fatalf("should be %f, got %f", v, result)
		}
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	//given
	var testCases = map[tempconv.Fahrenheit]tempconv.Celsius{
		-39: -39.444444,
		-1:  -18.333333,
		0:   -17.777778,
		9:   -12.777777778,
		39:  3.888889,
	}

	//when && then
	for k, v := range testCases {
		result := tempconv.FToC(k)
		if int(result) != int(v) {
			t.Fatalf("should be %f, got %f", v, result)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	//given
	var testCases = map[tempconv.Fahrenheit]tempconv.Kelvin{
		-39: 233.70555556,
		-1:  254.81666667,
		0:   255.37222222,
		9:   260.37222222,
		39:  277.03888889,
	}

	//when && then
	for k, v := range testCases {
		result := tempconv.FToK(k)
		if int(result) != int(v) {
			t.Fatalf("should be %f, got %f", v, result)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	//given
	var testCases = map[tempconv.Kelvin]tempconv.Fahrenheit{
		-39: -529.87,
		-1:  -461.47,
		0:   -459.67,
		9:   -443.47,
		39:  -389.47,
	}

	//when && then
	for k, v := range testCases {
		result := tempconv.KToF(k)
		if int(result) != int(v) {
			t.Fatalf("should be %f, got %f", v, result)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	//given
	var testCases = map[tempconv.Kelvin]tempconv.Celsius{
		-39: -312.15,
		-1:  -274.15,
		0:   -273.15,
		9:   -264.15,
		39:  -234.15,
	}

	//when && then
	for k, v := range testCases {
		result := tempconv.KToC(k)
		if int(result) != int(v) {
			t.Fatalf("should be %f, got %f", v, result)
		}
	}
}
