package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	// 有bug
	var p *MyError = nil

	//var p error = nil
	if bad() {
		p = &ErrBad
	}

	return p
}
func main() {
	err := returnsError()
	//if err != nil {
	//	fmt.Printf("error occur:%+v\n", err)
	//	return
	//}
	if e, ok := err.(*MyError); ok && e != nil {
		fmt.Printf("error occur: %+v\n", e)
		return
	}
	fmt.Println("ok")
}
