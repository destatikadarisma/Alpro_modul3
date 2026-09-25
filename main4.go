package main

import "fmt"

func main() {
	var f int
	var c float64

	// Membaca masukan suhu dalam Fahrenheit
	fmt.Scan(&f)

	// Menghitung konversi ke Celcius
	c = float64(f-32) * 5.0 / 9.0

	// Menampilkan hasil suhu Celcius
	fmt.Println(c)
}