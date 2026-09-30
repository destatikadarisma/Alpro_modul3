//contoh soal no laprak no 1
// menghitung hasil penjumlahan 5 bilangan bulat

package main
import "fmt"

func main() {
    var a, b, c, d, e int
    var hasil int
    fmt.Scanln(&a, &b, &c, &d, &e)
    hasil = a + b + c + d + e
    fmt.Println("Hasil penjumlahan", a, b, c, d, e, "adalah", hasil)
}
