package services

import (
	"errors"
	"learn-go-gin/config"
	"learn-go-gin/models"
	"learn-go-gin/utils"
)

func GetAllUsers(page int, limit int, sortBy string, sortOrder string) ([]models.User, int64, error) {
	var users []models.User

	var total int64

	offset := (page - 1) * limit
	sortQuery := sortBy + " " + sortOrder

	if err := config.DB.Order(sortQuery).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	total = int64(len(users))
	return users, total, nil
}

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	// Check if email already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already exists")
	}

	// Hash password before save into database
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	// Create user
	if err := config.DB.Create(user).Error; err != nil {
		return nil, errors.New("cannot create user: " + err.Error())
	}

	return user, nil
}

func UpdateUser(input *models.User, id string) (*models.User, error) {
	// Find existing user
	var existingUser models.User
	if err := config.DB.First(&existingUser, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Update fields from input to existingUser
	if input.Name != "" {
		existingUser.Name = input.Name
	}
	if input.Email != "" {
		existingUser.Email = input.Email
	}
	if input.Password != "" {
		hashedPassword, err := utils.HashPassword(input.Password)
		if err != nil {
			return nil, errors.New("failed to hash password")
		}
		existingUser.Password = hashedPassword
	}

	// Save updates
	if err := config.DB.Save(&existingUser).Error; err != nil {
		return nil, errors.New("cannot update user: " + err.Error())
	}

	return &existingUser, nil
}

func DeleteUser(user *models.User, id string) error {
	// check user exist
	if err := config.DB.First(user, id).Error; err != nil {
		return errors.New("user not found")
	}

	// Delete user
	if err := config.DB.Delete(user).Error; err != nil {
		return errors.New("cannot delete user: " + err.Error())
	}

	return nil
}

func Login(email, password string) (*string, error) {
	// find user by email
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// check password
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	token, _ := utils.GenerateJWTToken(user.ID)

	return &token, nil
}
