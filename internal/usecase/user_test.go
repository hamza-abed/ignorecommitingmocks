package usecase_test

import (
	"blackboxtests/internal/service"
	"blackboxtests/internal/usecase"
	"blackboxtests/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserDetails(t *testing.T) {
	// test cases
	testCases := []struct {
		name          string
		userID        int
		mockUser      *service.User
		mockError     error
		expectedName  string
		expectedError error
	}{
		{
			name:          "successful_get_user",
			userID:        1,
			mockUser:      &service.User{ID: 1, Username: "johndoe", Email: "john@example.com"},
			mockError:     nil,
			expectedName:  "johndoe",
			expectedError: nil,
		},
		{
			name:          "user_not_found",
			userID:        2,
			mockUser:      nil,
			mockError:     errors.New("user not found"),
			expectedName:  "",
			expectedError: errors.New("user not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUserService := new(mocks.UserService)

			mockUserService.On("GetUserByID", tc.userID).Return(tc.mockUser, tc.mockError)

			userUseCase := usecase.NewUserUseCase(mockUserService)

			username, err := userUseCase.GetUserDetails(tc.userID)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedName, username)

			mockUserService.AssertExpectations(t)
		})
	}
}
