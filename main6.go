//Tulislah sebuah algoritma untuk menghitung y = (x² + 1/x)².
//Masukan berupa satu bilangan riil, nilai x.
//Keluaran berupa hasil perhitungan y

package main

import (
	"fmt"
)

func main() {
	var x float64

	// Menerima masukan nilai x (bilangan riil)
	_, err := fmt.Scan(&x)
	if err != nil {
		return
	}

	// Memeriksa pembagian dengan nol
	if x == 0 {
		fmt.Println("Error: Pembagian dengan nol tidak terdefinisi")
		return
	}

	// Menghitung y = (x^2 + 1/x)^2
	inner := (x * x) + (1.0 / x)
	y := inner * inner

	// Menampilkan hasil dengan presisi 7 digit di belakang koma
	fmt.Printf("%.7f\n", y)
}