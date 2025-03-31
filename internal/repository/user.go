package repository

import (
	"blackboxtests/internal/service"
	"errors"
)

// UserRepository implements UserService interface
type UserRepository struct {
	users map[int]*service.User
}

// NewUserRepository create a new instance of UserRepository
func NewUserRepository() *UserRepository {
	repo := &UserRepository{
		users: make(map[int]*service.User),
	}

	// add some dummy data
	repo.users[1] = &service.User{
		ID:       1,
		Username: "johndoe",
		Email:    "john@example.com",
	}

	return repo
}

// GetUserByID returns a user by its ID
func (r *UserRepository) GetUserByID(id int) (*service.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
