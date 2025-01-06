package services

import (
	"errors"
	"learn-go-gin/config"
	"learn-go-gin/models"
)

func GetAllFines() ([]models.Fines, error) {
	var fines []models.Fines

	if err := config.DB.Find(&fines).Error; err != nil {
		return nil, errors.New("cannot get finess: " + err.Error())
	}

	return fines, nil
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
