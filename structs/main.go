package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Boy",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 75838,
		},
	}
	jim.updateName("Billy")
	jim.print()

	jim.updateName("bob")
	fmt.Printf("%+v", jim)
}

func (pointerToperson *person) updateName(newFirstName string) {
	(*pointerToperson).firstName = newFirstName
}
func (p person) print() {
	fmt.Println("%+v", p)
}
