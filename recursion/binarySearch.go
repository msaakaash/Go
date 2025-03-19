package main
import (
    "fmt"
    "sort"
)
func binarySearch(arr []int,left int,right int,target int) int{
    if left>right {
        return -1
    }
    mid := left + (right - left)/2
    
    if arr[mid]==target {
        return mid
    }else if arr[mid] > target  {
        return binarySearch(arr,left,mid-1,target)
    }else {
        return binarySearch(arr,mid+1,right,target)
    }
}


func main(){
	nums := []int{2,4,5,3,9,-1}
	sort.Ints(nums)
	target := 3
	fmt.Println("Sorted Array:", nums)
	present := binarySearch(nums,0,len(nums)-1,target);
	if present == -1 {
	    fmt.Println("Not found")
	}else {
	    fmt.Println(target,"found at index:",present)
	}
}