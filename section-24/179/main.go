package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	_, err = os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
		// log.Println("err happened", err)
		// log.Fatalln(err)
		// log.Panicln(err)
	}
}

func foo() {
	fmt.Println("Does this run?")
}
