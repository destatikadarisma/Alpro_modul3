package main

import (
	"fmt"
)

func main () {
	var tahun int
	fmt.Print("masukan tahun :")
	fmt.Scan(&tahun)

	//aturan tahun kabisat 
	isKabisat := (tahun%400 == 0 ) || (tahun%4 == 0 && tahun%100 != 0)
	fmt.Printf ("kabisat : %t\n", isKabisat)
}
