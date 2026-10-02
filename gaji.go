package main

import "fmt"

func main() {
	var gaji_pokok, potongan, bonus_lembur, jam_lembur, gaji_bersih int

	// Menerima input gaji pokok dan lama waktu lembur
	fmt.Scan(&gaji_pokok)
	fmt.Scan(&jam_lembur)

	// Proses hitung komponen gaji
	bonus_lembur = 45000 * jam_lembur
	potongan = ((2 * gaji_pokok) / 100) + ((35 * gaji_pokok) / 1000)
	gaji_bersih = gaji_pokok + bonus_lembur - potongan

	// Menampilkan output angka default (misal: 6300000)
	fmt.Println(gaji_bersih)

	// Menampilkan tambahan teks penjelasan dengan format ribuan rupiah
	fmt.Printf("Gaji bersih sebesar %d rupiah\n", gaji_bersih)
}

