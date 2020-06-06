package testpkg

import (
	"fmt"
	"testing"
)

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func TestClass(t *testing.T) {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}

func closure_test() func() int {
	i := 1
	return func() int { i+=1; return i}
}

func TestUserLogin_normal(t *testing.T) {
	fmt.Println("I am test login")
	var err error = nil
	if err == nil {
		t.Log("Login test1 success")
	} else {
		t.Error(err)
	}
}

func TestClosure(t *testing.T)  {
	y := 5
	a := func() (func()){
		var i int = 10
		return func() {
			fmt.Printf("i, y: %d, %d\n", i, y)
		}
	}()
	a()
	y = 10
	a()
}

func TestClosure1(t *testing.T)  {
	x  := closure_test()

	fmt.Println(x())
	fmt.Println(x())

	x1  := closure_test()
	fmt.Println(x1())
	fmt.Println(x1())
}

func TestRange(t *testing.T) {
	var arr []int = []int{1,2,3,4,5}
	for i, v := range arr {
		fmt.Printf("arr[%d]=%d\n", i, v)
	}

	kvs := map[string]string{"k1":"value1", "k2":"value2"}
	for k, v := range kvs {
		fmt.Printf("%s=%s\n", k, v)
	}
}

func TestMap(t *testing.T) {
	countryCapitalMap := make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	capital, ok := countryCapitalMap["France"]
	if(ok) {
		fmt.Println("Capital of France is:", capital)
	} else {
		fmt.Println("No capital of China found")
	}
	for key := range countryCapitalMap {
		fmt.Printf("Capital of %s is: %s\n", key, countryCapitalMap[key])
	}
	delete(countryCapitalMap, "India")
	for key := range countryCapitalMap {
		fmt.Println("Capital of", key, "is", countryCapitalMap[key])
	}
}

func Factorial(n int) (result int)  {
	if(n == 0) {
		result = 1
	} else {
		result = n * Factorial(n - 1)
	}
	return
}

func TestFactorial(t *testing.T)  {
	fmt.Println("Factorial of 9 is:", Factorial(4))
}

func Fibonacci(n int) (result int) {
	if(n < 2) {
		result = n
	} else {
		result = Fibonacci(n - 1) + Fibonacci(n - 2)
	}
	return
}
func TestFibonacci(t *testing.T) {
	var i int
	max := 60

	for i = 0; ; i++  {
		num := Fibonacci(i)
		if(num < max) {
			fmt.Println(Fibonacci(i))
		} else {
			break
		}
	}
}

func TestConversion(t *testing.T)  {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)
}