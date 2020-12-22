package main

import (
	"fmt"
)

func main() {
	favSport := "Swimming"
	switch favSport {
	case "Baseball":
		fmt.Println("Your favorite sport is baseball")
	case "Basketball":
		fmt.Println("Your favorite sport is basketball")
	case "Football":
		fmt.Println("Your favorite sport is football")
	case "American Football":
		fmt.Println("Your favorite sport is American football")
	default:
		fmt.Println("I don't know your favorite sport")
	}
}
