package testpkg

import (
	"fmt"
	"testing"
	"time"
)

//go GetThingDone(param1, param2)
//
//go func(param1, param2) {
//} (val1, val2)
//
//go {
//	// do something
//}
//
//ci := make(chan int)
//cs := make(chan string)
//cf := make(chan interface{})

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)

		time.Sleep(time.Second)
	}
}

func TestGoroutine(t *testing.T) {
	go running()

	var input string
	fmt.Scanln(&input)
}

func TestGoroutine1(t *testing.T) {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}