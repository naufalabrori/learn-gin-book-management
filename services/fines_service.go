package services

import (
	"errors"
	"learn-go-gin/config"
	"learn-go-gin/dto"
	"learn-go-gin/models"
)

func GetAllFines(page int, limit int, sortBy string, sortOrder string, search string) ([]dto.FinesWithTransaction, int64, error) {
	var fines []dto.FinesWithTransaction
	var total int64

	offset := (page - 1) * limit
	sortQuery := sortBy + " " + sortOrder

	// Query untuk total data (tanpa limit dan offset)
	if err := config.DB.Table("fines").
		Joins("LEFT JOIN transactions ON transactions.id = fines.transaction_id").
		Joins("LEFT JOIN users ON users.id = transactions.user_id").
		Joins("LEFT JOIN books ON books.id = transactions.book_id").
		Where("CAST(fines.amount AS TEXT) ILIKE ? OR users.email ILIKE ? OR users.name ILIKE ? OR books.title ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Query untuk data dengan limit, offset, dan sorting
	if err := config.DB.Table("fines").
		Joins("LEFT JOIN transactions ON transactions.id = fines.transaction_id").
		Joins("LEFT JOIN users ON users.id = transactions.user_id").
		Joins("LEFT JOIN books ON books.id = transactions.book_id").
		Where("CAST(fines.amount AS TEXT) ILIKE ? OR users.email ILIKE ? OR users.name ILIKE ? OR books.title ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Order(sortQuery).
		Limit(limit).
		Offset(offset).
		Select("fines.*, users.email as user_email, users.name as user_name, books.title as book_title, transactions.borrowed_date as borrowed_date, transactions.due_date as due_date, transactions.returned_date as returned_date").
		Find(&fines).Error; err != nil {
		return nil, 0, err
	}

	return fines, total, nil
}

func GetFinesByID(id string) (*models.Fines, error) {
	var fine models.Fines
	if err := config.DB.First(&fine, id).Error; err != nil {
		return nil, errors.New("cannot get fines: " + err.Error())
	}
	return &fine, nil
}

func GetFinesByTransactionID(transactionID string) ([]models.Fines, error) {
	var fines []models.Fines
	if err := config.DB.Where("transaction_id = ?", transactionID).Find(&fines).Error; err != nil {
		return nil, errors.New("cannot get fines: " + err.Error())
	}
	return fines, nil
}

func CreateFines(fine *models.Fines) (*models.Fines, error) {
	var existingFines models.Fines
	if err := config.DB.Where("transaction_id = ?", fine.TransactionID).First(&existingFines).Error; err == nil {
		return nil, errors.New("fines already exists")
	}

	if err := config.DB.Create(fine).Error; err != nil {
		return nil, errors.New("cannot create fines: " + err.Error())
	}
	return fine, nil
}

func UpdateFines(input *models.Fines, id string) (*models.Fines, error) {
	var existingFines models.Fines
	if err := config.DB.First(&existingFines, id).Error; err != nil {
		return nil, errors.New("fines not found")
	}

	if input.Amount != 0 {
		existingFines.Amount = input.Amount
	}
	if !input.PaidDate.IsZero() {
		existingFines.PaidDate = input.PaidDate
	}

	if err := config.DB.Save(&existingFines).Error; err != nil {
		return nil, errors.New("cannot update fines: " + err.Error())
	}
	return &existingFines, nil
}

func DeleteFines(fine *models.Fines, id string) error {
	if err := config.DB.First(fine, id).Error; err != nil {
		return errors.New("fines not found")
	}

	if err := config.DB.Delete(fine).Error; err != nil {
		return errors.New("cannot delete fines: " + err.Error())
	}

	return nil
}
