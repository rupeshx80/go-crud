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

func GetUserByEmailService(email string) (*models.User, error) {
	return repository.GetUserByEmail(email)
}

func GetUserByUsernameService(username string) (*models.User, error) {
	return repository.GetUserByUsername((username))
}

func DeleteUserService(id uint)(*models.User, error) {
	return repository.DeleteUser(id)
}

func UpdateUserService(id uint, newData map[string]interface{}) (*models.User, error) {
	
	if len(newData) == 0 {
		return nil, errors.New("no data provided for update")
	}

		if email, ok := newData["email"].(string); ok {
		existingUser, err := repository.GetUserByEmail(email)

		if err == nil && existingUser.ID != id {
			return nil, errors.New("email already in use by another user")
		}
	}

	if username, ok := newData["username"].(string); ok {
		existingUser, err := repository.GetUserByUsername(username)
		if err == nil && existingUser.ID != id {
			return nil, errors.New("username already in use by another user")
		}
	}

	updatedUser, err := repository.UpdateUser(id, newData)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}