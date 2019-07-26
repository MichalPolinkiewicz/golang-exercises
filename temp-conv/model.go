package temp_conv

import "fmt"

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100


	FreezingK = 273.15

)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%f 8C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%f 8C", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%f 8C", k)
}

