package services

import (
	"errors"
	"learn-go-gin/config"
	"learn-go-gin/models"
	"time"
)

func GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := config.DB.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetTransactionByID(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := config.DB.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	var existingTransaction models.Transaction
	if err := config.DB.Where("user_id = ? AND book_id = ? AND status = 'Borrowed'", transaction.UserID, transaction.BookID).First(&existingTransaction).Error; err == nil {
		return nil, errors.New("transaction already exists")
	}

	if err := config.DB.Create(transaction).Error; err != nil {
		return nil, errors.New("cannot create transaction: " + err.Error())
	}
	return transaction, nil
}

func UpdateTransaction(input *models.Transaction, id string) (*models.Transaction, error) {
	var existingTransaction models.Transaction
	if err := config.DB.First(&existingTransaction, id).Error; err != nil {
		return nil, errors.New("transaction not found")
	}

	if input.UserID != 0 {
		existingTransaction.UserID = input.UserID
	}
	if input.BookID != 0 {
		existingTransaction.BookID = input.BookID
	}
	if !input.BorrowedDate.IsZero() {
		existingTransaction.BorrowedDate = input.BorrowedDate
	}
	if !input.DueDate.IsZero() {
		existingTransaction.DueDate = input.DueDate
	}
	if !input.ReturnedDate.IsZero() {
		existingTransaction.ReturnedDate = input.ReturnedDate
	}
	if input.Status != "" {
		existingTransaction.Status = input.Status
	}

	if err := config.DB.Save(&existingTransaction).Error; err != nil {
		return nil, errors.New("cannot update transaction: " + err.Error())
	}

	return &existingTransaction, nil
}

func DeleteTransaction(transaction *models.Transaction, id string) error {
	if err := config.DB.First(transaction, id).Error; err != nil {
		return errors.New("transaction not found")
	}

	if err := config.DB.Delete(transaction).Error; err != nil {
		return errors.New("cannot delete transaction: " + err.Error())
	}

	return nil
}

func ReturnTransaction(transaction *models.Transaction, id string) error {
	if err := config.DB.First(transaction, id).Error; err != nil {
		return errors.New("transaction not found")
	}

	transaction.ReturnedDate = time.Now()

	dueDate := time.Date(transaction.DueDate.Year(), transaction.DueDate.Month(), transaction.DueDate.Day(), 0, 0, 0, 0, transaction.DueDate.Location())
	returnedDate := time.Date(transaction.ReturnedDate.Year(), transaction.ReturnedDate.Month(), transaction.ReturnedDate.Day(), 0, 0, 0, 0, transaction.ReturnedDate.Location())

	if returnedDate.After(dueDate) {
		transaction.Status = "Overdue"
	} else {
		transaction.Status = "Returned"
	}

	if err := config.DB.Save(transaction).Error; err != nil {
		return errors.New("cannot return transaction: " + err.Error())
	}

	return nil
}
