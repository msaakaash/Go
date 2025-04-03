package main

import "fmt"

func multiply(row [3]int, ma [3][3]int, ch chan [3]int) {
	var ans [3]int
	var sum int

	for i := 0; i < 3; i++ {
		sum = 0
		for j := 0; j < 3; j++ {
			sum = sum + (row[j] * ma[j][i])
		}
		ans[i] = sum
	}
	ch <- ans
}

func main() {
	a := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	b := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	temp := make(chan [3]int)
	temp2 := make(chan [3]int)
	temp3 := make(chan [3]int)

	go multiply(a[0], b, temp)
	go multiply(a[1], b, temp2)
	go multiply(a[2], b, temp3)

	var final [3][3]int
	final[0] = <-temp
	final[1] = <-temp2
	final[2] = <-temp3

	fmt.Println(final)
}
