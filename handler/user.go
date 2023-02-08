package handler

import (
	"fmt"
	"kakastartup/auth"
	"kakastartup/helper"
	"kakastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterInput
	// struct di atas kita passing seabagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationsError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadGateway, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadGateway, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.Formatter(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// //user memasukan input (email&password)
	// input ditangkap handler
	// mapping dari input user input struct
	// input struct passing service
	// di service mencari dengan bantuan repository user dengan email x
	// mencocokan password

	var input user.LoginInput

	//proses mapping atau penangkapan data
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationsError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	LoggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	token, err := h.authService.GenerateToken(LoggedInUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadGateway, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//mengembalikan dalam format json
	formatter := user.Formatter(LoggedInUser, token)
	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) EmailCheckAbility(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationsError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	iseEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"Is_available": iseEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if iseEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusUnprocessableEntity, "success", data)
	c.JSON(http.StatusUnprocessableEntity, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploded": true}
	response := helper.APIResponse("Upload successfuly", http.StatusOK, "success", data)
	c.JSON(http.StatusBadRequest, response)

}
