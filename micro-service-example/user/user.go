package user

// User type explains the base of the struct
type User struct {
	ID string
	Birthday string
	Contact ContactType
	CreatedAt string
	FirstName string
	Name string
	Followers []User
	Friends []User
	LastName string
	UpdatedAt string
	Username string
}

// ContactType for the 
type ContactType struct {
	email string
	phone string
}