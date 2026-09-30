//menghitung luas lingkaran berdasarkan jari-jari
package main

import "fmt"

func main() {
	var r, luas float64
	const PI = 3.14159265358979323846

	// Membaca input jari-jari
	fmt.Scan(&r)

	// Menghitung luas
	luas = PI * r * r

	// Menampilkan hasil luas dengan presisi 1 angka di belakang koma
	fmt.Printf("%.1f\n", luas)
}