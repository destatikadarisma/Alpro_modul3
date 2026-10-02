package main

import "fmt"

func main() {
	var Waktu, Sisa, Jam, Menit, Detik int
// Menerima input total detik
	fmt.Scan(&Waktu)
// Proses hitung konversi waktu
	Jam = Waktu / 3600
	Sisa = Waktu % 3600
	Menit = Sisa / 60
	Detik = Sisa % 60

	fmt.Println(Jam, Menit, Detik)
// Menampilkan tambahan teks penjelasan (Explanation)
	fmt.Printf("%d = %d jam %d menit %d detik\n", Waktu, Jam, Menit, Detik)
}

