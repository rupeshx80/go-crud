package repository

import (
	"errors"
	"fmt"

	"github.com/rupeshx80/go-crud/pkg/db"
	"github.com/rupeshx80/go-crud/pkg/models"
)

func CreateUser(user *models.User) error {
	err := db.RJ.Create(user).Error

	if err != nil{
		return errors.New("failed to create user")
	}
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.RJ.Where("email = ?", email).First(&user).Error 

	if err != nil {
		return nil, err
	}
	fmt.Println("Checking whats in &user", &user)
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.RJ.Where("username = ?", username).First(&user).Error
	
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id uint, newData map[string]interface{}) (*models.User, error) {
    var user models.User
    err := db.RJ.First(&user, id).Error
    if err != nil {
        return nil, err
    }

    err = db.RJ.Model(&user).Updates(newData).Error
	
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func DeleteUser(id uint)(*models.User, error){

	var user models.User

	 err := db.RJ.First(&user, id).Error
    if err != nil {
        return nil, err
    }

	 err = db.RJ.Delete(&user).Error
	
    if err != nil {
        return nil, err
    }

    return &user, nil
}