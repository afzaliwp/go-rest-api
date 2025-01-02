package controllers

import (
	"errors"
	"github.com/afzaliwp/go-rest-api/db"
	"github.com/afzaliwp/go-rest-api/helpers"
	"github.com/afzaliwp/go-rest-api/models"
	"github.com/afzaliwp/go-rest-api/services"
	"github.com/afzaliwp/go-rest-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Signup(c *gin.Context) {
	if c.PostForm("email") == "" || c.PostForm("password") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to signup user",
			"error":   errors.New("email or password is required"),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(c.PostForm("password"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to signup user",
			"error":   err.Error(),
		})

		return
	}

	user := models.NewUser(
		c.PostForm("name"),
		c.PostForm("email"),
		hashedPassword,
	)

	err = services.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to signup user",
			"error":   err.Error(),
		})

		return
	}

	user.Password = "" //Hide the user password from the final output
	c.JSON(http.StatusOK, gin.H{
		"message": "signup user successfully",
		"user":    user,
	})
}

func Login(c *gin.Context) {
	if c.PostForm("email") == "" || c.PostForm("password") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to signup user",
			"error":   errors.New("email and password is required"),
		})

		return
	}

	var userModel = models.NewUser("", "", "")
	IsValidCreds := IsValidCredentials(c, &userModel)

	if !IsValidCreds {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed to login user",
			"error":   "email or password is not correct",
		})

		return
	}

	token, err := utils.GenerateToken(userModel.ID, userModel.Email, userModel.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to login user",
			"error":   "failed to generate token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login user successfully",
		"token":   token,
	})
}

func IsValidCredentials(c *gin.Context, userModel *models.User) bool {
	email := c.PostForm("email")
	rawInputPassword := c.PostForm("password")
	query := "SELECT id, name, password FROM users WHERE email=?"

	userModel.Email = email

	row := db.DB.QueryRow(query, email)
	var name string
	var userIdString string
	var retrievedPassword string
	err := row.Scan(&userIdString, &name, &retrievedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to login user",
			"error":   err.Error(),
		})

		return false
	}
	helpers.MyLog("after scan")
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to login user",
			"error":   err.Error(),
		})
	}
	userModel.ID = userId
	userModel.Name = name

	passwordIsValid := utils.ComparePasswords(rawInputPassword, retrievedPassword)

	return passwordIsValid
}
