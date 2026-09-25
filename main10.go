//Tulislah sebuah algoritma untuk menukar posisi digit dari sebuah bilangan tiga digit. 
//Contoh, untuk masukan 123, hasilnya adalah 321.

//Masukan berupa satu bilangan bulat positif tiga digit.

//Keluaran berupa bilangan bulat tiga digit lain dari bilangan asli,
//menggunakan proses seperti dijelaskan pada contoh.

//Petunjuk: Anda harus membaca masukan sebagai satu bilangan bulat, 
//bukan sebagai 3 karakter atau 3 bilangan satu digit. Demikian pula, 
//Anda harus menulis keluaran sebagai satu bilangan bulat, bukan sebagai 3 karakter atau 3 bilangan satu digit.

package main

import (
	"fmt"
)

func main() {
	var n int

	// Membaca satu bilangan bulat tiga digit
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	// Mengisolasi digit-digitnya
	digit1 := n / 100        // Digit ratusan (misal: 1 dari 123)
	digit2 := (n / 10) % 10  // Digit puluhan (misal: 2 dari 123)
	digit3 := n % 10         // Digit satuan  (misal: 3 dari 123)

	// Menyusun ulang menjadi bilangan terbalik
	reversed := (digit3 * 100) + (digit2 * 10) + digit1

	// Mencetak satu bilangan bulat utuh
	fmt.Println(reversed)
}