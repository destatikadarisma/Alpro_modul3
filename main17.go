//menghitung persamaan f(x) = (2/x+5) + 5
package main
import "fmt"

func main(){
    var x, fx float64
    fmt.Scan(&x)
    fx = 2/(x+5) + 5
    fmt.Println(fx)
}

