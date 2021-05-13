package main

import (
	"fmt"
)

func main() {
	numbers := []int{4, 3, 1, 8, 11, 1, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		out := "odd"
		if number%2 == 0 {
			out = "even"
		}
		fmt.Println(number, "is", out)
	}
}