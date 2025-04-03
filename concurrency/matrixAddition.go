package main

import "fmt"

func addRows(row1 []int, row2 []int, ch chan []int) {
	res := make([]int, len(row1))
	for ind, val := range row1 {
		res[ind] = val + row2[ind]
	}
	ch <- res
}

func main() {
	a := [][]int{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
	}
	b := [][]int{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
	}

	temp := make(chan []int)

	var ans = make([][]int, len(a))

	for i := 0; i < 3; i++ {
		go addRows(a[i], b[i], temp)
		ans[i] = <-temp
	}

	fmt.Println(ans)
}
