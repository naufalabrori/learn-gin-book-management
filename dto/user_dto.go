package dto

import (
	"learn-go-gin/models"
	"time"
)

type UserResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	PhoneNumber  string    `json:"phone_number"`
	CreatedDate  time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
}

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
