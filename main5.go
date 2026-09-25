//D: Tulislah sebuah algoritma seperti aplikasi kalkulator untuk menghitung operasi aritmetika antara dua bilangan pecahan a/b dan c/d.
//Masukan berupa empat bilangan bulat a, b, c, dan d secara berurutan, yang merepresentasikan bilangan pecahan a/b dan c/d.
//Keluaran berupa hasil operasi penjumlahan, pengurangan, perkalian, 
// dan pembagian dalam bentuk pecahan x/y. Untuk setiap operasi, tampilkan label operasi, pembilang x, karakter "/", dan penyebut y secara berurutan

package main

import (
	"fmt"
)

// Fungsi untuk menghitung Pembagi Bersama Terbesar (FPB / GCD)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

// Fungsi untuk menyederhanakan pecahan
func simplify(num, den int) (int, int) {
	if den < 0 {
		num = -num
		den = -den
	}
	g := gcd(num, den)
	if g == 0 {
		return num, den
	}
	return num / g, den / g
}

func main() {
	var a, b, c, d int

	// Membaca masukan 4 bilangan bulat: a, b, c, d
	_, err := fmt.Scan(&a, &b, &c, &d)
	if err != nil || b == 0 || d == 0 {
		return
	}

	// 1. Penjumlahan: (a/b) + (c/d) = (a*d + c*b) / (b*d)
	addNum, addDen := simplify((a*d)+(c*b), b*d)
	fmt.Println("addition:")
	fmt.Println(addNum)
	fmt.Println("/")
	fmt.Println(addDen)

	// 2. Pengurangan: (a/b) - (c/d) = (a*d - c*b) / (b*d)
	subNum, subDen := simplify((a*d)-(c*b), b*d)
	fmt.Println("subtraction:")
	fmt.Println(subNum)
	fmt.Println("/")
	fmt.Println(subDen)

	// 3. Perkalian: (a/b) * (c/d) = (a*c) / (b*d)
	mulNum, mulDen := simplify(a*c, b*d)
	fmt.Println("multiplication:")
	fmt.Println(mulNum)
	fmt.Println("/")
	fmt.Println(mulDen)

	// 4. Pembagian: (a/b) / (c/d) = (a*d) / (b*c)
	if c != 0 {
		divNum, divDen := simplify(a*d, b*c)
		fmt.Println("division:")
		fmt.Println(divNum)
		fmt.Println("/")
		fmt.Println(divDen)
	}
}