package user

// User type explains the base of the struct
type User struct {
	id string
	birthday string
	contact Contact
	createdAt string
	firstName string
	name string
	followers []User
	friends []User
	lastName string
	updatedAt string
	username string
}

// Contact for the 
type Contact struct {
	email string
	phone string
}