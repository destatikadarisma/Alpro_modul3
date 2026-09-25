//Tulislah sebuah algoritma untuk menukar isi tiga variabel, sehingga:

//variabel y berisi nilai dari x
//variabel x berisi nilai dari z
//variabel z berisi nilai dari y
//Masukan berupa tiga bilangan yang disimpan ke variabel x, y, z.

//Keluaran berupa isi variabel x, y, z setelah proses penukaran.

//Petunjuk: Pembacaan harus menggunakan variabel x, y, z 
//dalam urutan ini, dan penulisan juga harus menggunakan variabel 
//yang sama dalam urutan yang sama (x, y, z).
package main

import (
	"fmt"
)

func main() {
	var x, y, z int

	// Membaca masukan ke dalam variabel x, y, z secara berurutan
	_, err := fmt.Scan(&x, &y, &z)
	if err != nil {
		return
	}

	// Menukar nilai variabel secara bersamaan:
	// x baru mengambil nilai z lama
	// y baru mengambil nilai x lama
	// z baru mengambil nilai y lama
	x, y, z = z, x, y

	// Mencetak nilai x, y, z setelah ditukar
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}