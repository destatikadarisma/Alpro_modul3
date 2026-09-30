//cd pkg
// filename : ascii.go
// program ASCII yang akan membaca 5 data intenger dan mencetak dalam pormat karakter
//dan mencetak 3 buah karakter 
package main
import "fmt"

func main(){
    var c1, c2, c3, c4, c5 byte
    var b1, b2, b3 int

    fmt.Scan(&c1, &c2, &c3, &c4, &c5)
    fmt.Scanf("%c", &b1)
    fmt.Scanf("%c", &b2)
    fmt.Scanf("%c", &b3)

    fmt.Printf("%c%c%c%c%c\n", c1, c2, c3, c4, c5)
    fmt.Printf("%c%c%c\n", b1+1, b2+1, b3+1)
}

