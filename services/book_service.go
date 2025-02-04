package services

import (
	"errors"
	"learn-go-gin/config"
	"learn-go-gin/models"
)

func GetAllBooks(page int, limit int, sortBy string, sortOrder string, search string) ([]models.Book, int64, error) {
	var books []models.Book
	var total int64

	offset := (page - 1) * limit
	sortQuery := sortBy + " " + sortOrder

	// Query untuk total data (tanpa limit dan offset)
	if err := config.DB.Model(&models.Book{}).
		Where("title ILIKE ? OR author ILIKE ? OR isbn ILIKE ? OR publisher ILIKE ? OR published_year ILIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Query untuk data dengan limit, offset, dan sorting
	if err := config.DB.Where(
		"title ILIKE ? OR author ILIKE ? OR isbn ILIKE ? OR publisher ILIKE ? OR published_year ILIKE ?",
		"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%",
	).Order(sortQuery).Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		return nil, 0, err
	}

	return books, total, nil
}

func GetBookByID(id string) (*models.Book, error) {
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func CreateBook(book *models.Book) (*models.Book, error) {
	// Check if ISBN already exists in database
	var existingBook models.Book
	if err := config.DB.Where("isbn = ?", book.ISBN).First(&existingBook).Error; err == nil {
		return nil, errors.New("ISBN already exists")
	}

	// Check if category exists in database
	var existingCategory models.Category
	if err := config.DB.First(&existingCategory, book.CategoryID).Error; err != nil {
		return nil, errors.New("category not found")
	}

	book.AvailableQuantity = book.Quantity
	// Create book
	if err := config.DB.Create(book).Error; err != nil {
		return nil, errors.New("cannot create book: " + err.Error())
	}

	return book, nil
}

func UpdateBook(book *models.Book, id string) (*models.Book, error) {
	// Find existing book
	var existingBook models.Book
	if err := config.DB.First(&existingBook, id).Error; err != nil {
		return nil, errors.New("book not found")
	}

	// Update fields from input to existingBook
	if book.Title != "" {
		existingBook.Title = book.Title
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}
	if book.Publisher != "" {
		existingBook.Publisher = book.Publisher
	}
	if book.PublishedYear != "" {
		existingBook.PublishedYear = book.PublishedYear
	}
	if book.ISBN != "" {
		existingBook.ISBN = book.ISBN
	}
	if book.CategoryID != 0 {
		existingBook.CategoryID = book.CategoryID
	}
	if book.Quantity != 0 {
		existingBook.Quantity = book.Quantity
	}
	if book.AvailableQuantity != 0 {
		existingBook.AvailableQuantity = book.AvailableQuantity
	}

	// Save updates
	if err := config.DB.Save(&existingBook).Error; err != nil {
		return nil, errors.New("cannot update book: " + err.Error())
	}

	return &existingBook, nil
}

func DeleteBook(book *models.Book, id string) error {
	// check book exist
	if err := config.DB.First(book, id).Error; err != nil {
		return errors.New("book not found")
	}

	// Delete book
	if err := config.DB.Delete(book).Error; err != nil {
		return errors.New("cannot delete book: " + err.Error())
	}

	return nil
}
