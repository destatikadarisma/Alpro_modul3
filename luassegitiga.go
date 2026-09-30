package main

import "fmt"

func main(){
	var alas, tinggi int
	fmt.Print("masukan panjang alas yang merupakan bilangan bukat positif")
	fmt.Scan(&alas)

	fmt.Print("masukan tinggi segitiga yang merupakan bilangan bukat positif")
	fmt.Scan(&tinggi)

	if alas <=0 || tinggi <=0 {
		fmt.Println("error: alas dan tinggi harus bilangan bulat positif")
		return
	}
	luas := 0.5 * float64(alas) * float64(tinggi)
	fmt.Println("luas segitiga adalah", luas)
}
