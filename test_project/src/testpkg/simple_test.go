package testpkg

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"testing"
	"time"
)


//func TestArr(t *testing.T) {
//	var arr [5]int
//	var i,j int
//
//	for i = 0; i < 5; i++ {
//		arr[i] = 100 + i
//	}
//
//	for j = 0; j < 5; j++ {
//		fmt.Println(arr[j])
//	}
//}

func TestArr2(t *testing.T) {
	var a [3][4] int
	a = [3][4] int{
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11},  /*  第三行索引为 2 */
	}
	var k,l int
	for k = 0; k < 3; k++ {
		for l = 0; l < 4; l++ {
			fmt.Print(a[k][l], " ")
		}
		fmt.Println()
	}


}

func getAvg(arr []int, size int) float32 {
	var sum, i int = 0, 0

	for i = 0; i < size ;i++  {
		sum += arr[i]
	}

	return float32(sum/size)
}

func TestArrWithParam(t *testing.T)  {
	var arr []int = []int{100,2,10,2000,23}

	avg := getAvg(arr, 5)
	fmt.Printf("avarage:%f\n", avg)
}

func TestPtrArr(t *testing.T)  {
	var arr []int = []int{100,2,10,2000,23}
	var size int = 5
	var ptrArr [5]*int

	i := 0
	for i = 0; i < size; i++ {
		ptrArr[i] = &arr[i]
	}
	for i = 0; i < size; i++  {
		fmt.Printf("arr[%d]=%d\n", i, *ptrArr[i])
	}
}

func TestPptr(t *testing.T)  {
	var a int = 1000
	p := &a
	pp := &p

	fmt.Printf("a=%d,p=%d,pp=%d\n", a, *p, **pp)
}

func swap(x *int, y *int)  {
	tmp := 0
	tmp = *x
	*x = *y
	*y = tmp
}
func TestSwap(t *testing.T) {
	a, b := 10,20
	swap(&a, &b)
	fmt.Printf("swap a, b = %d, %d\n", a, b)
}

type Book struct {
	title string
	subject string
	author string
	book_id int
}

func printBook(book *Book)  {
	fmt.Printf("title:%s, author:%s, subject:%s, book_id:%d\n", book.title, book.author, book.subject, book.book_id)
}

func TestStruct(t *testing.T)  {
	var Book1 Book
	var Book2 Book
	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.w3cschool.cn"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.w3cschool.cn"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	printBook(&Book1)
	printBook(&Book2)
}

func TestSwitch(t *testing.T) {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func TestSlice(t *testing.T)  {
	//s := []struct{
	//	i int
	//	j bool
	//} {
	//	{1,true},
	//	{2,false},
	//	{3,true},
	//	{4, false},
	//	{5, false},
	//	{6, false},
	//}
	//
	//fmt.Println(s);
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	str := "  this is a good this test is    a good    "
	fmt.Printf("filds are:%q", strings.Fields(str))
}

type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M()  {
	fmt.Println(t.S)
}

type F float64

func (f F) M()  {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func TestInterface1(t *testing.T) {
	var i I
	i = &T{"test"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var addr [4]byte = [4]byte{1,2,3,4}
	var j int
	for j = 0; j < 4; j++  {
		fmt.Println(addr[j])
	}

	x := make([]byte, 10, 10)
	fmt.Println(x)
}

