//Diberikan dua bilangan bulat positif, periksa apakah bilangan kedua adalah faktor dari bilangan pertama.

//Masukan berupa dua bilangan bulat positif x dan y, dengan x > y.

//Keluaran berupa bilangan bulat 1 jika y adalah faktor dari x, dan 0 jika bukan.

//Catatan: Gunakan berbagai operasi bilangan bulat.
package main

import (
	"fmt"
)

func main() {
	var x, y int

	// Membaca masukan x dan y
	_, err := fmt.Scan(&x, &y)
	if err != nil || y == 0 {
		return
	}

	// Memeriksa apakah y =adalah faktor dari x menggunakan operasi modulo (%)
	if x%y == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}