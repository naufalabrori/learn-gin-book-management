package services

import (
	"errors"
	"fmt"
	"io"
	"learn-go-gin/config"
	"learn-go-gin/models"
	"learn-go-gin/utils"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func GetAllUsers(page int, limit int, sortBy string, sortOrder string, search string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	offset := (page - 1) * limit
	sortQuery := sortBy + " " + sortOrder

	// Query untuk total data (tanpa limit dan offset)
	if err := config.DB.Model(&models.User{}).
		Where("name ILIKE ? OR email ILIKE ? OR role ILIKE ? OR phone_number ILIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Query untuk data dengan limit, offset, dan sorting
	if err := config.DB.Where(
		"name ILIKE ? OR email ILIKE ? OR role ILIKE ? OR phone_number ILIKE ?",
		"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%",
	).Order(sortQuery).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

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
	if input.Role != "" {
		existingUser.Role = input.Role
	}
	if input.PhoneNumber != "" {
		existingUser.PhoneNumber = input.PhoneNumber
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

func Login(email, password string) (*models.User, *string, error) {
	// find user by email
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, nil, errors.New("user not found")
	}

	// check password
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, nil, errors.New("invalid password")
	}

	token, _ := utils.GenerateJWTToken(user.ID)

	return &user, &token, nil
}

func UserImage(id string, file multipart.FileHeader) (*models.User, error) {
	// find user by id
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Validasi ukuran file (misalnya maksimum 5 MB)
	maxFileSize := int64(5 << 20) // 5 MB
	if file.Size > maxFileSize {
		return nil, errors.New("file size exceeds limit of 5 mb")
	}

	// Validasi tipe file (misalnya hanya menerima jpg, png)
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	fileExtension := filepath.Ext(file.Filename)
	valid := false
	for _, ext := range allowedExtensions {
		if ext == fileExtension {
			valid = true
			break
		}
	}
	if !valid {
		return nil, errors.New("invalid file type")
	}

	// Tentukan lokasi penyimpanan
	uploadDir := "./uploads/images"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm) // Buat folder jika belum ada
	}

	// Buat nama file yang unik
	timestamp := time.Now().Unix()
	newFileName := fmt.Sprintf("%d%s", timestamp, fileExtension)
	filePath := filepath.Join(uploadDir, newFileName)

	// Buka file untuk disalin
	srcFile, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer srcFile.Close()

	// Simpan file ke server
	destFile, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file on server: %v", err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return nil, fmt.Errorf("failed to save file: %v", err)
	}

	// Perbarui kolom image pada user
	user.Image = newFileName
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, fmt.Errorf("cannot update user: %v", err)
	}

	return &user, nil
}
