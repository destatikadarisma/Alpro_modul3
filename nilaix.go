package main

import "fmt"

func main () {
	var fx float64
	fmt.Print("masukan nilai f(x) :")
	fmt.Scan(&fx)

	x := (2 / (fx + 5)) + 5

	fmt.Printf("nilai x: %g\n", x)
}