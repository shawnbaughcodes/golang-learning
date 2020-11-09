package main

// Group type for groups in the app
type Group struct {
	id string
	members []User
	name string
	subject string
	createdAt string
	updatedAt string
}