package controllers

import (
	"learn-go-gin/dto"
	"learn-go-gin/models"
	"learn-go-gin/services"
	"learn-go-gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
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

	// call services to get all users
	users, total, err := services.GetAllUsers(page, limit, sortBy, sortOrder, search)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mapping to dto
	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dto.ToUserResponse(&user))
	}

	// Jika userResponses kosong, set ke array kosong
	if len(userResponses) == 0 {
		userResponses = make([]dto.UserResponse, 0) // Pastikan array kosong, bukan nil
	}

	// Buat respons
	response := gin.H{
		"data":      userResponses,
		"totalData": total,
	}

	utils.RespondSuccess(c, "Fetched users successfully", response)
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

	userResponses := dto.ToUserResponse(createdUser)
	utils.RespondSuccess(c, "User created successfully", userResponses)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUserByID(id)

	if user == nil {
		utils.RespondError(c, err.Error(), http.StatusNotFound)
		return
	} else {
		userResponses := dto.ToUserResponse(user)
		utils.RespondSuccess(c, "Fetched user", userResponses)
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

	userResponses := dto.ToUserResponse(updatedUser)
	utils.RespondSuccess(c, "User updated successfully", userResponses)
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

	user, token, err := services.Login(input.Email, input.Password)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	data := gin.H{
		"token": token,
		"user":  dto.ToUserResponse(user),
	}

	utils.RespondSuccess(c, "Login successful", data)
}

func UploadUserImage(c *gin.Context) {
	id := c.Param("id")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get file"})
		return
	}

	uploadImage, err := services.UserImage(id, *file)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	userResponses := dto.ToUserResponse(uploadImage)
	utils.RespondSuccess(c, "User Image Upload successfully", userResponses)
}
