package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	golf := person{
		firstName: "Noppadol",
		lastName:  "La",
		contact: contactInfo{
			email: "noppaodllanngain@gmail.com",
			zip:   50180,
		},
	}
	golf.print()
}
