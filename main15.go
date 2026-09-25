//Tulislah sebuah algoritma untuk menyisipkan sebuah angka di tengah bilangan empat digit. C
//contoh, untuk masukan 1234 dan 5, hasilnya adalah 12534.
//Masukan berupa dua bilangan bulat positif, yang pertama berdigit empat dan yang kedua berdigit satu.
//Keluaran berupa bilangan bulat baru setelah proses di atas.
package main

import (
	"fmt"
)

func main() {
	var n, d int

	// Membaca masukan: n (bilangan 4 digit) dan d (angka sisipan 1 digit)
	_, err := fmt.Scan(&n, &d)
	if err != nil {
		return
	}

	// Memisahkan dua digit depan dan dua digit belakang
	depan := n / 100
	belakang := n % 100

	// Menyusun kembali bilangan baru
	hasil := (depan * 1000) + (d * 100) + belakang

	// Mencetak hasil
	fmt.Println(hasil)
}