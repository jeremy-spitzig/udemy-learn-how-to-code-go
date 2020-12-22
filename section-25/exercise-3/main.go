package main

import "fmt"

type customErr struct {
	message string
}

func (e customErr) Error() string {
	return e.message
}

func main() {
	foo(customErr{"This is a custom error"})
}

func foo(e error) {
	fmt.Println(e.Error())
}
