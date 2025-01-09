package dto

import (
	"time"
)

type TransactionWithUserAndBook struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	Email        string    `json:"email"`
	BookID       uint      `json:"book_id"`
	BookTitle    string    `json:"book_title"`
	BorrowedDate time.Time `json:"borrowed_date"`
	DueDate      time.Time `json:"due_date"`
	ReturnedDate time.Time `json:"returned_date"`
	Status       string    `json:"status"`
	CreatedDate  time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
}
