package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	fmt.Println(woohoo("Miss", "Moneypenny"))
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprintf("Hello from woo, %v", s)
}

func woohoo(fn, ln string) (string, bool) {
	return fmt.Sprintf("Hello from woohoo, %v %v", fn, ln), true
}
