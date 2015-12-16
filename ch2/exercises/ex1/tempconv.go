package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroC
}

// FToK converts a Fahrenheit temperature to Kelvin
func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

// KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

func main() {
	for _, arg := range os.Args[1:] {
		if temperature, err := strconv.ParseFloat(arg, 64); err == nil {
			fmt.Printf("%s = %s = %s\n", Celsius(temperature), CToK(Celsius(temperature)), CToF(Celsius(temperature)))
			fmt.Printf("%s = %s = %s\n", Kelvin(temperature), KToC(Kelvin(temperature)), KToF(Kelvin(temperature)))
			fmt.Printf("%s = %s = %s\n", Fahrenheit(temperature), FToC(Fahrenheit(temperature)), FToK(Fahrenheit(temperature)))
		} else {
			fmt.Println(err)
		}
	}
}
