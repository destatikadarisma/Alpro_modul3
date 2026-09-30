package main

import "fmt"

func main() {
	var sisi int
	_, err := fmt.Scan(&sisi)
	if err != nil || sisi <= 0 {
		return
	}
	sisiFloat := float64(sisi)
	volume := sisiFloat * sisiFloat *sisiFloat

	fmt.Printf("%.1f\n", volume)
}