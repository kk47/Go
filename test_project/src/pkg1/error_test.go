package pkg1

import (
	"fmt"
	"testing"
)

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(var_e int, var_r int)(result int, err_msg string) {
	if var_r == 0 {
		dData := DivideError{
			dividee: var_e,
			divider: 0,
		}
		err_msg = dData.Error()
		return
	} else {
		return var_e / var_r, ""
	}
}

func TestError(t *testing.T) {
	if r, msg := Divide(100, 10); msg == "" {
		fmt.Println("100/10=", r)
	}
	if _, msg := Divide(100, 0); msg != "" {
		fmt.Println("100/0:", msg)
	}
}