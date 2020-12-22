package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(string(bs))
	loginPword1 := `password123`
	if err2 := bcrypt.CompareHashAndPassword(bs, []byte(loginPword1)); err2 != nil {
		fmt.Println("Not fine")
	} else {
		fmt.Println("Just fine")
	}
	loginPword2 := `asdf`
	if err2 := bcrypt.CompareHashAndPassword(bs, []byte(loginPword2)); err2 != nil {
		fmt.Println("Not fine")
	} else {
		fmt.Println("Just fine")
	}
}
