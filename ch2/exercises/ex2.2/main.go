package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"./lengthconv"
	"./tempconv"
	"./weightconv"
)

func main() {
	var inputs []float64
	var inputsAsStrings []string

	if len(os.Args) == 1 {
		// read numbers from standard input:
	} else {
		inputsAsStrings = os.Args[1:]
	}

	for _, number := range inputsAsStrings {
		if value, err := strconv.ParseFloat(number, 64); err != nil {
			log.Fatalf("Invalid number: %s", number)
		} else {
			inputs = append(inputs, value)
		}
	}

	for _, number := range inputs {
		fmt.Printf("%s = %s\n", tempconv.Kelvin(number), tempconv.KToC(tempconv.Kelvin(number)))
		fmt.Printf("%s = %s\n", tempconv.Kelvin(number), tempconv.KToF(tempconv.Kelvin(number)))
		fmt.Printf("%s = %s\n", tempconv.Celsius(number), tempconv.CToK(tempconv.Celsius(number)))
		fmt.Printf("%s = %s\n", tempconv.Celsius(number), tempconv.CToF(tempconv.Celsius(number)))
		fmt.Printf("%s = %s\n", tempconv.Fahrenheit(number), tempconv.FToK(tempconv.Fahrenheit(number)))
		fmt.Printf("%s = %s\n", tempconv.Fahrenheit(number), tempconv.FToC(tempconv.Fahrenheit(number)))
		fmt.Printf("%s = %s\n", weightconv.Kilogram(number), weightconv.KToP(weightconv.Kilogram(number)))
		fmt.Printf("%s = %s\n", weightconv.Pound(number), weightconv.PToK(weightconv.Pound(number)))
		fmt.Printf("%s = %s\n", lengthconv.Meter(number), lengthconv.MToY(lengthconv.Meter(number)))
		fmt.Printf("%s = %s\n", lengthconv.Yard(number), lengthconv.YToM(lengthconv.Yard(number)))
	}
}
