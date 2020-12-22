package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		fmt.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return math.Sqrt(f), nil
}
