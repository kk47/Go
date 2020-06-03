package main

import "http_proxy"

//var counter int = 0
//
//func Count(lock *sync.Mutex) {
//	lock.Lock()
//	counter++
//	fmt.Println(counter)
//	lock.Unlock()
//}

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

	http_proxy.ProxyMain()
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