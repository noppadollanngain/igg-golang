package main

import "fmt"

func main() {
	ages := map[string]int{
		"golf": 25,
		"benz": 25,
		"bee":  25,
	}
	updateMap(ages, "bee", 99)
	printMap(ages)
}

func printMap(ages map[string]int) {
	for name, age := range ages {
		fmt.Println(name, age)
	}
}

func updateMap(ages map[string]int, name string, age int) {
	ages[name] = age
}
