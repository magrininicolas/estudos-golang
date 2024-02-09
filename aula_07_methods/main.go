package main

import "fmt"

type fahrenheit float64
type celsius float64
type kelvin float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9/5 + 32)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var tempAtualF fahrenheit = 32
	var tempAtualC celsius = 0
	var tempAtualK kelvin = 273.15

	fmt.Printf("%.2fºC\n", tempAtualF.celsius())
	fmt.Printf("%.2fºF\n", tempAtualC.fahrenheit())
	fmt.Printf("%.2fK\n", tempAtualK.celsius())
}
