//Diberikan sebuah bilangan bulat, periksa apakah bilangan tersebut negatif.

//Masukan berupa satu bilangan bulat.

//Keluaran berupa bilangan bulat 1 jika masukan negatif, dan 0 jika bukan.

//Catatan: Gunakan berbagai operasi bilangan bulat.

//Catatan 2: Bilangan bulat disimpan sebagai bilangan 32 bit.
package main

import (
	"fmt"
)

func main() {
	var n int32

	// Membaca masukan satu bilangan bulat 32-bit
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	// Mengambil Sign Bit (bit ke-31):
	// Jika n negatif -> 1
	// Jika n >= 0   -> 0
	hasil := (n >> 31) & 1

	fmt.Println(hasil)
}