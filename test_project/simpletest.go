package main

import (
	"fmt"
	"strings"
)

//var counter int = 0
//
//func Count(lock *sync.Mutex) {
//	lock.Lock()
//	counter++
//	fmt.Println(counter)
//	lock.Unlock()
//}
func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	var res map[string]int = make(map[string]int)
	for _, v := range arr {
		if value, ok := res[v]; ok {
			res[v] = value+1
		} else {
			res[v] = 1
		}
	}
	return res
}
//func main(){
	//go func() {
	//	var times int
	//	for {
	//		times++
	//		fmt.Println("tick", times)
	//		time.Sleep(time.Second)
	//	}
	//}()
	//var input string
	//fmt.Scanln(&input)

//	lock := &sync.Mutex{}
//	for i:=0; i < 10; i++ {
//		go Count(lock)
//	}
//	for{
//		lock.Lock()
//		c := counter
//		lock.Unlock()
//		runtime.Gosched()
//		if c >= 10 {
//			break
//		}
//	}
//}
//var (
//	count int64
//	wg    sync.WaitGroup
//)
func main() {
	//wg.Add(2)
	//go incCount()
	//go incCount()
	//wg.Wait()
	//fmt.Println(count)

	//http_proxy.ProxyMain()

	a := WordCount(" this is a test that is a test 11 this a hah   ")
	fmt.Println(a)

	var arr [][]int = make([][]int, 5)
	for i, _ := range arr {
		arr[i] = make([]int, 4)
	}
	fmt.Println(arr, arr[0], arr[0][0])
}
//func incCount() {
//	defer wg.Done()
//	for i := 0; i < 2; i++ {
//		//value := count
//		//runtime.Gosched()
//		//value++
//		//count = value
//
//		atomic.AddInt64(&count, 1)
//		runtime.Gosched()
//	}
//}