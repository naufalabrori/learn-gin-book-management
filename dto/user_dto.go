package dto

import (
	"learn-go-gin/models"
	"time"
)

// UserResponse adalah DTO untuk menampilkan data user tanpa kolom Password
type UserResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	PhoneNumber  string    `json:"phoneNumber"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
}

// ToUserResponse mengonversi model User ke dalam bentuk UserResponse
func ToUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Role:         user.Role,
		PhoneNumber:  user.PhoneNumber,
		CreatedDate:  user.CreatedDate,
		ModifiedDate: user.ModifiedDate,
	}
}
