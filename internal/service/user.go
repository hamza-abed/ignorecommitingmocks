package service

// UserService defines the interface for user operations
type UserService interface {
	GetUserByID(id int) (*User, error)
}

// User represents a user in the system
type User struct {
	ID       int
	Username string
	Email    string
}
