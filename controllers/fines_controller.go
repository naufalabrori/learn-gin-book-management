package controllers

import (
	"learn-go-gin/models"
	"learn-go-gin/services"
	"learn-go-gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFines(c *gin.Context) {
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

	fines, total, err := services.GetAllFines(page, limit, sortBy, sortOrder, search)

	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response := gin.H{
		"data":      fines,
		"totalData": total,
	}

	utils.RespondSuccess(c, "Fetched all fines", response)
}

func GetFinesByTransactionID(c *gin.Context) {
	transactionID := c.Param("transactionId")
	fines, err := services.GetFinesByTransactionID(transactionID)

	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "Fetched all fines", fines)
}

func GetFinesByID(c *gin.Context) {
	id := c.Param("id")
	fine, err := services.GetFinesByID(id)

	if fine == nil {
		utils.RespondError(c, err.Error(), http.StatusNotFound)
		return
	} else {
		utils.RespondSuccess(c, "Fetched fines", fine)
	}
}

func CreateFines(c *gin.Context) {
	var fine models.Fines
	if err := c.ShouldBindJSON(&fine); err != nil {
		utils.RespondError(c, "Cannot create fines: invalid input", http.StatusBadRequest)
		return
	}
	createdFine, err := services.CreateFines(&fine)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RespondSuccess(c, "Fines created successfully", createdFine)
}

func UpdateFines(c *gin.Context) {
	id := c.Param("id")
	var fine models.Fines
	if err := c.ShouldBindJSON(&fine); err != nil {
		utils.RespondError(c, "Cannot update fines: invalid input", http.StatusBadRequest)
		return
	}
	updatedFine, err := services.UpdateFines(&fine, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RespondSuccess(c, "Fines updated successfully", updatedFine)
}

func DeleteFines(c *gin.Context) {
	id := c.Param("id")
	var fine models.Fines
	err := services.DeleteFines(&fine, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RespondSuccess(c, "Fines deleted successfully", nil)
}
