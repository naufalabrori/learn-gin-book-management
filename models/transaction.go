package models

import "time"

type Transaction struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	UserID       uint      `json:"userId" gorm:"not null"`
	BookID       uint      `json:"bookId" gorm:"not null"`
	BorrowedDate time.Time `json:"borrowedDate" gorm:"not null"`
	DueDate      time.Time `json:"dueDate" gorm:"not null"`
	ReturnedDate time.Time `json:"returnedDate"`
	Status       string    `json:"status" gorm:"not null"`
	CreatedDate  time.Time `json:"createdDate" gorm:"autoCreateTime"`
	ModifiedDate time.Time `json:"modifiedDate" gorm:"autoUpdateTime"`
}
