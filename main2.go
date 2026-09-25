package main

import "fmt"

func main() {
	var nama, nim, kelas string

	// Menerima masukan dari pengguna
	fmt.Print(" Nama  : ")
	fmt.Scanln(&nama)

	fmt.Print("NIM   : ")
	fmt.Scanln(&nim)

	fmt.Print("Kelas : ")
	fmt.Scanln(&kelas)

	// Menampilkan resume biodata
	fmt.Println()
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n", nama, kelas, nim)
}