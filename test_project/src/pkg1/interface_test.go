package pkg1

import (
	"fmt"
	"testing"
)

type Phone interface {
	call()
}

type Iphone struct {
}

type NokiaPhone struct {

}

func (ip Iphone) call() {
	fmt.Println("Iphone calling")
}

func (np NokiaPhone) call() {
	fmt.Println("Nokia Phone calling")
}

func TestInterface(t *testing.T) {
	var p Phone

	p = new(Iphone)
	p.call()

	p = new(NokiaPhone)
	p.call()
}