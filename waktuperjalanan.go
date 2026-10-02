package main

import "fmt"

func main() {
	var kecepatan, menit int 
	var jam, total_jarak int 

	fmt.Scan(&kecepatan) 

	total_jarak = 100 + 60 + 170 
	menit = (total_jarak * 60) / kecepatan 
	jam = menit / 60              
	menit = menit % 60            

	// Menampilkan output angka default (misal: 6 36)
	fmt.Println(jam, menit) 

	// Menampilkan tambahan teks penjelasan (Explanation)
	fmt.Printf("%d jam %d menit\n", jam, menit)
}
