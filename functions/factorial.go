package main
import "fmt"


func factorial(n int) int{
	var f int = 1
	for i:=1 ; i<=n ;i++ {
		f*=i;
	}
	return f
}

func fact_arr(n [5]int){
    for i:=0 ; i< len(n);i++ {
        fmt.Printf("The factorial of %d is: %d\n", n[i], factorial(n[i])); 
    }
}
func main(){
	var numbers = [5]int{2,3,1,5,6}
    fact_arr(numbers)
}