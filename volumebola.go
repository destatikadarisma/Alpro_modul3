package main

import (
	"fmt"
	"math"
)

func main () {
	var r int
	fmt.Print("masukan jari-jari bola :")
	fmt.Scan(&r)

	const pi = 3.1415926535
	rasiusFloat := float64(r)

	volume := (4.0 / 3.0) * pi * math.Pow(rasiusFloat, 3)
	luas := 4 * pi * math.Pow(rasiusFloat, 2)

	fmt.Printf ("bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
