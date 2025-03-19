package main
import (
    "fmt"
)

func reverse(n int,rev int) int{
    if n == 0{
        return rev
    }
    return reverse(n/10,rev*10+(n%10))
}

func main(){
	num := 2202
	rev := reverse(num,0)
	fmt.Println(num,"in reverse order is",rev)
}