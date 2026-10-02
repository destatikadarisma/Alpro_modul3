package main

import "fmt"

func main() {
	var N, hari int
	fmt.Scan(&N)

	// Rumus mencari kode hari berikutnya 
	hari = (4+N-1)%7 + 1

	// Menampilkan output angka default 
	fmt.Println(hari)

	// Membuat nama hari untuk mempermudah cetak teks penjelasan
	namaHari := [8]string{"", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	// Menampilkan tambahan teks penjelasan (Explanation)
	fmt.Printf("%d hari setelah hari Kamis (4) adalah %s (%d)\n", N, namaHari[hari], hari)
}

