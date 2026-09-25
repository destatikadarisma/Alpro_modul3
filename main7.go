//Tulislah sebuah algoritma untuk mencetak karakter 
//yang sama seperti yang dibaca dari masukan.
//Masukan berupa satu karakter.
//Keluaran berupa karakter yang sama.
package main

import (
	"fmt"
)

func main() {
	var c rune

	// Membaca satu karakter dari masukan
	_, err := fmt.Scanf("%c", &c)
	if err != nil {
		return
	}

	// Mencetak karakter yang sama
	fmt.Printf("%c\n", c)
}