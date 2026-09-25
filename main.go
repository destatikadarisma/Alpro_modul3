package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("expresso: ")
	fmt.Scan(&satu)
	fmt.Print("americano: ")
	fmt.Scan(&dua)
	fmt.Print("capuccino: ")
	fmt.Scan(&tiga)

	// Cetak Output Awal (Wajib ada)
	fmt.Println("Output awal = ", satu, dua, tiga)
	// Proses Penggeseran
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	// Cetak Output Akhir
	fmt.Println("Output akhir = ", satu, dua, tiga)

}
