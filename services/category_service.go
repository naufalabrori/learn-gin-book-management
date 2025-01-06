package services

import (
	"errors"
	"learn-go-gin/config"
	"learn-go-gin/models"
)

func GetAllCategories(page int, limit int, sortBy string, sortOrder string) ([]models.Category, int64, error) {
	var categories []models.Category

	var total int64

	offset := (page - 1) * limit
	sortQuery := sortBy + " " + sortOrder

	if err := config.DB.Order(sortQuery).Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	total = int64(len(categories))
	return categories, total, nil
}

func GetCategoryByID(id string) (*models.Category, error) {
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func CreateCategory(category *models.Category) (*models.Category, error) {
	// Check if category name already exists
	var existingCategory models.Category
	if err := config.DB.Where("category_name = ?", category.CategoryName).First(&existingCategory).Error; err == nil {
		return nil, errors.New("category name already exists")
	}

	if err := config.DB.Create(category).Error; err != nil {
		return nil, errors.New("cannot create category: " + err.Error())
	}
	return category, nil
}

func UpdateCategory(category *models.Category, id string) (*models.Category, error) {
	// Find existing category
	var existingCategory models.Category
	if err := config.DB.First(&existingCategory, id).Error; err != nil {
		return nil, errors.New("category not found")
	}

	// Update fields from input to existingCategory
	if category.CategoryName != "" {
		existingCategory.CategoryName = category.CategoryName
	}

	// Save updates
	if err := config.DB.Save(&existingCategory).Error; err != nil {
		return nil, errors.New("cannot update category: " + err.Error())
	}

	return &existingCategory, nil
}

func DeleteCategory(category *models.Category, id string) error {
	// check category exist
	if err := config.DB.First(category, id).Error; err != nil {
		return errors.New("category not found")
	}

	// Delete category
	if err := config.DB.Delete(category).Error; err != nil {
		return errors.New("cannot delete category: " + err.Error())
	}

	return nil
}
