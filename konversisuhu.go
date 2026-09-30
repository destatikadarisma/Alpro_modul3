package main
import (
	"fmt"
	"math"
)
func main() {
	var celcius float64
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scan(&celcius)

	reamur := celcius * 4.0 / 5.0
	fahrenheit := (celcius * 9.0 / 5.0) + 32.0
	
	kelvin := (fahrenheit + 459.67) * 5.0 / 9.0
	fmt.Printf("Derajat Reamur: %.0f\n", math.Round(reamur))
	fmt.Printf("Derajat Fahrenheit: %.0f\n", math.Round(fahrenheit))
	fmt.Printf("Derajat Kelvin: %.0f\n", math.Round(kelvin))
}
