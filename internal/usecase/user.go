package usecase

import (
	"blackboxtests/internal/service"
)

// UserUseCase is the structure that uses the UserService
type UserUseCase struct {
	userService service.UserService
}

// NewUserUseCase creates a new instance of UserUseCase
func NewUserUseCase(userService service.UserService) *UserUseCase {
	return &UserUseCase{
		userService: userService,
	}
}

// GetUserDetails retrieves user details by their ID
func (uc *UserUseCase) GetUserDetails(id int) (string, error) {
	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}
