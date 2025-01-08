package models

import "time"

type Transaction struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	BookID       uint      `json:"book_id" gorm:"not null"`
	BorrowedDate time.Time `json:"borrowed_date" gorm:"not null"`
	DueDate      time.Time `json:"due_date" gorm:"not null"`
	ReturnedDate time.Time `json:"returned_date"`
	Status       string    `json:"status" gorm:"not null"`
	CreatedDate  time.Time `json:"created_date" gorm:"autoCreateTime"`
	ModifiedDate time.Time `json:"modified_date" gorm:"autoUpdateTime"`
}
