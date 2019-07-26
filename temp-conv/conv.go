package temp_conv

//convert's celsius to fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//convert's celsius to kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

//convert's fahrenheit to celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32)*5/9)
}

//convert's fahrenheit to kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67)* 5/9)
}

//convert's fahrenheit to kelvin
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k* 9/5 - 459.67)
}

//convert's kelvin to celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
