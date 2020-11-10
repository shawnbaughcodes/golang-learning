package user

// User type explains the base of the struct
type User struct {
	ID string
	Birthday string
	Contact Contact
	CreatedAt string
	FirstName string
	Name string
	Followers []User
	Friends []User
	LastName string
	UpdatedAt string
	Username string
}

// Contact for the 
type Contact struct {
	Email string
	Phone string
}