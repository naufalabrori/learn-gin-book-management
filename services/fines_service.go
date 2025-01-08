package services

import (
	"errors"
	"learn-go-gin/config"
	"learn-go-gin/models"
)

func GetAllFines(page int, limit int, sortBy string, sortOrder string, search string) ([]models.Fines, int64, error) {
	var fines []models.Fines
	var total int64

	offset := (page - 1) * limit
	sortQuery := sortBy + " " + sortOrder

	// Query untuk total data (tanpa limit dan offset)
	if err := config.DB.Model(&models.Fines{}).
		Where("transaction_id ILIKE ? OR amount ILIKE ?", "%"+search+"%", "%"+search+"%").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Query untuk data dengan limit, offset, dan sorting
	if err := config.DB.Where("transaction_id ILIKE ? OR amount ILIKE ?", "%"+search+"%", "%"+search+"%").
		Order(sortQuery).
		Limit(limit).
		Offset(offset).
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
	if input.PaidDate.IsZero() {
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
