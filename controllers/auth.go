package controllers

import (
	"errors"
	"github.com/afzaliwp/go-rest-api/models"
	"github.com/afzaliwp/go-rest-api/services"
	"github.com/afzaliwp/go-rest-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
