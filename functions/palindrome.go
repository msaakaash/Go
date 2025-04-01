package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	name := "madama"

	valid := isPalindrome(name)

	fmt.Println(name, "is palindrome??:", valid)
}
