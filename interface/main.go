package main

import "fmt"

type bot interface {
	getGeeting() string
}

type thBot struct{}
type enBot struct{}

func main() {
	botTH := thBot{}
	botEN := enBot{}

	print(botTH, "golf")
	print(botEN, "golf")
}

func print(b bot, name string) {
	fmt.Println(b.getGeeting(), name)
}

func (thBot) getGeeting() string {
	return "Sawanddee"
}

func (enBot) getGeeting() string {
	return "Hello"
}
