package testpkg

import (
	"fmt"
	"testing"
)

func Sum(arr []int, c chan int)  {
	var i int
	sum := 0
	for i = 0; i < len(arr) ;i++  {
		sum += arr[i]
	}
	c <- sum
}

func TestRoutine(t *testing.T) {
	arr := []int{1,2,4,5,-1,4,56,11}
	c1 := make(chan int)
	c2 := make(chan int)

	go Sum(arr[:len(arr)/2], c1)
	go Sum(arr[len(arr)/2:], c2)

	x := <-c1
	y := <-c2

	fmt.Println(x, y, x+y)
}
