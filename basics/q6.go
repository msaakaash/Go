SApackage main

import "fmt"

func main() {
var a float64 = 3
var b float64 = 9
var res = div(a,b)

fmt.Printf("Value is :%.2f\n",res)

}

func div(n,n2 float64) float64 {
return n/n2
}
