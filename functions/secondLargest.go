// Constraints numbers [1...n] 

package main
import "fmt"

func secondLargest(n []int) (int,int){
    if len(n) < 2 {
        fmt.Println("Array must have at least two elements")
        return -1,-1;
    }
    var largest int = -1
    var secondLargest int = -1
    for i:=0 ; i<len(n) ; i++ {
        if n[i]>largest {
            secondLargest = largest
            largest = n[i]
        }else if n[i]<largest && n[i]>secondLargest {
            secondLargest = n[i]
        }
    }
    return largest,secondLargest
}
func main(){
	var numbers = []int{2,4,5,3,9};
	max_1,max_2 := secondLargest(numbers)
    fmt.Println("Largest:",max_1)
    fmt.Println("Second Largest:",max_2)
}