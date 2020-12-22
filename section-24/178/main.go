package main

import "fmt"

func main() {
	var answer1, answer2, answer3 string
	read("Name", &answer1)
	read("Favorite Food", &answer2)
	read("Favorite Sport", &answer3)

	fmt.Println(answer1, answer2, answer3)
}

func read(what string, where *string) {
	fmt.Printf("%v:", what)
	_, err := fmt.Scan(where)
	if err != nil {
		fmt.Printf("%v is required!\n", what)
		panic(err)
	}
}
