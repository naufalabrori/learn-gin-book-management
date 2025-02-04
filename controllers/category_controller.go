package controllers

import (
	"learn-go-gin/models"
	"learn-go-gin/services"
	"learn-go-gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
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

	categories, total, err := services.GetAllCategories(page, limit, sortBy, sortOrder, search)

	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response := gin.H{
		"data":      categories,
		"totalData": total,
	}

	utils.RespondSuccess(c, "Fetched all categories", response)
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := services.GetCategoryByID(id)

	if category == nil {
		utils.RespondError(c, err.Error(), http.StatusNotFound)
		return
	} else {
		utils.RespondSuccess(c, "Fetched category", category)
	}
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	// Binding JSON to Category struct
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.RespondError(c, "Cannot create category: invalid input", http.StatusBadRequest)
		return
	}

	createdCategory, err := services.CreateCategory(&category)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Category created successfully", createdCategory)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	// Binding JSON to Category struct
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.RespondError(c, "Cannot update category: invalid input", http.StatusBadRequest)
		return
	}

	updatedCategory, err := services.UpdateCategory(&category, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Category updated successfully", updatedCategory)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	err := services.DeleteCategory(&category, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Category deleted successfully", nil)
}
