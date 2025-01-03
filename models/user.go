package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name" gorm:"size:100;not null"`
	Email        string    `json:"email" gorm:"size:100;unique"`
	Password     string    `json:"password" gorm:"size:100;not null"`
	Role         string    `json:"role" gorm:"size:100;not null"`
	PhoneNumber  string    `json:"phoneNumber" gorm:"size:20"`
	CreatedDate  time.Time `json:"createdDate" gorm:"autoCreateTime"`
	ModifiedDate time.Time `json:"modifiedDate" gorm:"autoUpdateTime"`
}
