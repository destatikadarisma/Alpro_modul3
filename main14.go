//Seorang atlet maraton mencatat waktu larinya menggunakan stopwatch (dalam detik).
//Tulislah sebuah algoritma untuk mengonversi catatan waktu tersebut menjadi jam, menit, dan detik.

//Masukan berupa satu bilangan bulat positif, waktu lari atlet dalam detik.

//Keluaran berupa catatan waktu yang sama dalam jam, menit, dan detik.
package main

import (
	"fmt"
)

func main() {
	var totalDetik int

	// Membaca masukan total waktu dalam detik
	_, err := fmt.Scan(&totalDetik)
	if err != nil || totalDetik < 0 {
		return
	}

	// Menghitung jam, menit, dan detik
	jam := totalDetik / 3600
	sisaDetik := totalDetik % 3600

	menit := sisaDetik / 60
	detik := sisaDetik % 60

	// Mencetak hasil keluaran
	fmt.Println(jam)
	fmt.Println(menit)
	fmt.Println(detik)
}
