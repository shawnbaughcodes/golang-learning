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

	// jimPointer := &jim
	// jimPointer.updateName("Billy")
	// jim.print()
	// fmt.Printf("%+v", jim)

	// var alex person
	// fmt.Printf("%+v", alex)
	// alex.firstName = "Alex"
	// alex.lastName = "Bobby"

	// fmt.Println(alex)
}

func (pointerToperson *person) updateName(newFirstName string) {
	(*pointerToperson).firstName = newFirstName
}
func (p person) print() {
	fmt.Println("%+v", p)
}
