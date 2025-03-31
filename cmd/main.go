package main

import (
	"blackboxtests/internal/repository"
	"fmt"
	"log"

	"blackboxtests/internal/usecase"
)

func main() {
	// Init repo
	userRepo := repository.NewUserRepository()

	// Init usecase with repo
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Example of use case usage
	// In a real case, the ID would probably come from an HTTP request
	userID := 1
	username, err := userUseCase.GetUserDetails(userID)
	if err != nil {
		log.Fatalf("Error retrieving user details: %v", err)
	}

	fmt.Printf("user name: %s\n", username)
}
