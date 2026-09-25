//Diberikan sebuah bilangan bulat signed 32-bit, cetak nilai mutlak dari bilangan tersebut. Nilai masukan berada pada rentang -2147483647 sampai 2147483647.

//Masukan berupa satu bilangan bulat.

//Keluaran berupa nilai mutlak dari bilangan tersebut.

//Catatan: Gunakan berbagai operasi bilangan bulat. Jangan gunakan struktur percabangan.

//Catatan 2: Bilangan bulat disimpan sebagai bilangan 32 bit.
package main

import (
	"fmt"
)

func main() {
	var x int32

	// Membaca input
	_, err := fmt.Scan(&x)
	if err != nil {
		return
	}

	// Geser ke kanan 31 bit untuk mendapatkan mask
	// Jika x positif/nol: mask = 0
	// Jika x negatif: mask = -1 (semua bit bernilai 1)
	mask := x >> 31

	// Menghitung nilai mutlak tanpa if-else
	abs := (x + mask) ^ mask

	fmt.Println(abs)
}