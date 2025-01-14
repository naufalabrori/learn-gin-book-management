package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name" gorm:"size:100;not null"`
	Email        string    `json:"email" gorm:"size:100;unique;not null"`
	Password     string    `json:"password" gorm:"size:100;not null"`
	Role         string    `json:"role" gorm:"size:100;not null"`
	PhoneNumber  string    `json:"phone_number" gorm:"size:20"`
	Image        string    `json:"images" gorm:"size:50"`
	CreatedDate  time.Time `json:"created_date" gorm:"autoCreateTime"`
	ModifiedDate time.Time `json:"modified_date" gorm:"autoUpdateTime"`
}
