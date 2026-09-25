//Kasir Algomart menyimpan uang kertas sepuluh-ribuan, lima-ribuan, dan seribuan di mesin kasir.
//Ketika pelanggan membayar lebih dari total belanjaannya, kembalian diambil dari mesin kasir. 
//Tulislah sebuah algoritma untuk menghitung jumlah lembar uang kertas dari tiap pecahan yang dikembalikan kasir.
//Masukan berupa satu bilangan bulat, jumlah uang kembalian yang harus diterima pelanggan.
//Keluaran berupa jumlah lembar uang sepuluh-ribuan, lima-ribuan, dan seribuan.
package main

import (
	"fmt"
)

func main() {
	var kembalian int

	// Membaca masukan jumlah uang kembalian
	_, err := fmt.Scan(&kembalian)
	if err != nil {
		return
	}

	// Menghitung jumlah lembar 10.000
	lembar10k := kembalian / 10000
	sisa := kembalian % 10000

	// Menghitung jumlah lembar 5.000
	lembar5k := sisa / 5000
	sisa = sisa % 5000

	// Menghitung jumlah lembar 1.000
	lembar1k := sisa / 1000

	// Mencetak hasil secara berurutan
	fmt.Println(lembar10k)
	fmt.Println(lembar5k)
	fmt.Println(lembar1k)
}