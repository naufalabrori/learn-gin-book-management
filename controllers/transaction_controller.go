package controllers

import (
	"learn-go-gin/models"
	"learn-go-gin/services"
	"learn-go-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	transactions, err := services.GetAllTransactions()

	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "Fetched all transactions", transactions)
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
