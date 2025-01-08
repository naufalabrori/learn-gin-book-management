package controllers

import (
	"learn-go-gin/models"
	"learn-go-gin/services"
	"learn-go-gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
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

	transactions, total, err := services.GetAllTransactions(page, limit, sortBy, sortOrder, search)

	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response := gin.H{
		"data":      transactions,
		"totalData": total,
	}

	utils.RespondSuccess(c, "Fetched all transactions", response)
}

func GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	transaction, err := services.GetTransactionByID(id)

	if transaction == nil {
		utils.RespondError(c, err.Error(), http.StatusNotFound)
		return
	} else {
		utils.RespondSuccess(c, "Fetched transaction", transaction)
	}
}

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction

	// Binding JSON to Transaction struct
	if err := c.ShouldBindJSON(&transaction); err != nil {
		utils.RespondError(c, "Cannot create transaction: invalid input", http.StatusBadRequest)
		return
	}

	createdTransaction, err := services.CreateTransaction(&transaction)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Transaction created successfully", createdTransaction)
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	// Binding JSON to Transaction struct
	if err := c.ShouldBindJSON(&transaction); err != nil {
		utils.RespondError(c, "Cannot update transaction: invalid input", http.StatusBadRequest)
		return
	}

	updatedTransaction, err := services.UpdateTransaction(&transaction, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Transaction updated successfully", updatedTransaction)
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	err := services.DeleteTransaction(&transaction, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Transaction deleted successfully", nil)
}

func ReturnTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	err := services.ReturnTransaction(&transaction, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Transaction returned successfully", nil)
}
