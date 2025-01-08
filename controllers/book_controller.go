package controllers

import (
	"learn-go-gin/models"
	"learn-go-gin/services"
	"learn-go-gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	// Get query parameters from request
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) // Default page = 1
	if err != nil || page < 1 {
		utils.RespondError(c, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10")) // Default limit = 10
	if err != nil || limit < 1 {
		utils.RespondError(c, "Invalid limit parameter", http.StatusBadRequest)
		return
	}

	sortBy := c.DefaultQuery("sort_by", "id")        // Default sorting by "id"
	sortOrder := c.DefaultQuery("sort_order", "asc") // Default order is "asc"

	search := c.DefaultQuery("search", "")

	books, total, err := services.GetAllBooks(page, limit, sortBy, sortOrder, search)

	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response := gin.H{
		"data":      books,
		"totalData": total,
	}

	utils.RespondSuccess(c, "Fetched all books", response)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := services.GetBookByID(id)

	if book == nil {
		utils.RespondError(c, err.Error(), http.StatusNotFound)
		return
	} else {
		utils.RespondSuccess(c, "Fetched book", book)
	}
}

func CreateBook(c *gin.Context) {
	var book models.Book

	// Binding JSON to Book struct
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.RespondError(c, "Cannot create book: invalid input", http.StatusBadRequest)
		return
	}

	createdBook, err := services.CreateBook(&book)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Book created successfully", createdBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	// Binding JSON to Book struct
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.RespondError(c, "Cannot update book: invalid input", http.StatusBadRequest)
		return
	}

	updatedBook, err := services.UpdateBook(&book, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Book updated successfully", updatedBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	err := services.DeleteBook(&book, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Book deleted successfully", nil)
}
