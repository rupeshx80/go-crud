

package service

import (
	"errors"

	"github.com/rupeshx80/go-crud/pkg/models"
	"github.com/rupeshx80/go-crud/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(user *models.User) (*models.User, error) {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return nil, errors.New("username, email and password are required")
	}

	_, err := repository.GetUserByEmail(user.Email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	_, err = repository.GetUserByUsername(user.Username)
	if err == nil {
		return nil, errors.New("username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)

	if err := repository.CreateUser(user); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}
