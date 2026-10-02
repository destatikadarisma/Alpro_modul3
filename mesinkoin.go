package main

import "fmt"

func main() {
	var emas, sisa, perak, tembaga, koin int

	// Menerima input total koin
	fmt.Scan(&koin)

	// Proses perhitungan koin
	emas = koin / 9
	sisa = koin % 9
	perak = sisa / 3
	tembaga = koin % 3 // sisa akhir koin tembaga

	// Menampilkan output angka default (misal: 2 0 2)
	fmt.Println(emas, perak, tembaga)

	// Menampilkan tambahan teks penjelasan (Explanation)
	fmt.Println("Explanation:")
	fmt.Printf("Emas = %d (%d koin)\n", emas, emas*9)
	fmt.Printf("Perak = %d (%d koin)\n", perak, perak*3)
	fmt.Printf("Tembaga = %d (%d koin)\n", tembaga, tembaga*1)
}
