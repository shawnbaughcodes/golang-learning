package group

import "../user"

// Group type for groups in the app
type Group struct {
	id string
	members []user.User
	name string
	subject string
	createdAt string
	updatedAt string
}