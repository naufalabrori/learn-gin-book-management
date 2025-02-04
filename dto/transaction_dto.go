package dto

import (
	"time"
)

type TransactionWithUserAndBook struct {
	ID            uint      `json:"id"`
	UserID        uint      `json:"user_id"`
	UserEmail     string    `json:"user_email"`
	UserName      string    `json:"user_name"`
	BookID        uint      `json:"book_id"`
	BookTitle     string    `json:"book_title"`
	BorrowedDate  time.Time `json:"borrowed_date"`
	DueDate       time.Time `json:"due_date"`
	ReturnedDate  time.Time `json:"returned_date"`
	Status        string    `json:"status"`
	FinesPaidDate string    `json:"fines_paid_date"`
	CreatedDate   time.Time `json:"created_date"`
	ModifiedDate  time.Time `json:"modified_date"`
}
