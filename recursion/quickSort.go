package main
import (
    "fmt"
)

func partition(arr []int,left int,right int) int{
    pivot := arr[right]
    i := left - 1;
    
    for j:= left; j<=right ;j++ {
        if arr[j] < pivot {
            i++
            arr[i] , arr[j] = arr[j] , arr[i]
        }
    }
    arr[i+1],arr[right] =  arr[right], arr[i+1]
    return i+1
}


func quickSort(arr []int,left int,right int) {
    if left < right {
        part := partition(arr,left,right)
        quickSort(arr,left,part-1)
        quickSort(arr,part+1,right)
    }
}

func main(){
    nums := []int{5,-1,10,2,7,0}
    fmt.Println("Array:",nums)
    quickSort(nums,0,len(nums)-1)
	fmt.Println("Sorted array:",nums)
}