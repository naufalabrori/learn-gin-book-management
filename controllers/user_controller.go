package controllers

import (
	"learn-go-gin/models"
	"learn-go-gin/services"
	"learn-go-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()

	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "Fetched all users", users)
}

func CreateUser(c *gin.Context) {
	var user models.User

	// Binding JSON to User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, "Cannot create user: invalid input", http.StatusBadRequest)
		return
	}

	createdUser, err := services.CreateUser(&user)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "User created successfully", createdUser)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUserByID(id)

	if user == nil {
		utils.RespondError(c, err.Error(), http.StatusNotFound)
		return
	} else {
		utils.RespondSuccess(c, "Fetched user", user)
	}

}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Binding JSON to User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, "Cannot update user: invalid input", http.StatusBadRequest)
		return
	}

	updatedUser, err := services.UpdateUser(&user, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "User updated successfully", updatedUser)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	err := services.DeleteUser(&user, id)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "User deleted successfully", nil)
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, "Invalid input", http.StatusBadRequest)
		return
	}

	err := services.Login(input.Email, input.Password)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondSuccess(c, "Login successful", nil)
}
