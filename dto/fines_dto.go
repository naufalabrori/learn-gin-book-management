package dto

import (
	"time"
)

type FinesWithTransaction struct {
	ID            uint      `json:"id"`
	TransactionID uint      `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	PaidDate      time.Time `json:"paid_date"`
	UserEmail     string    `json:"user_email"`
	UserName      string    `json:"user_name"`
	BookTitle     string    `json:"book_title"`
	BorrowedDate  time.Time `json:"borrowed_date"`
	DueDate       time.Time `json:"due_date"`
	ReturnedDate  time.Time `json:"returned_date"`
	CreatedDate   time.Time `json:"created_date"`
	ModifiedDate  time.Time `json:"modified_date"`
}
