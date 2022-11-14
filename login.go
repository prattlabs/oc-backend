package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login holds the login information for a user.
type Login struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Token    string `form:"token" json:"token" xml:"token" binding:"required"`
}

// handleLogin handles the login request.
func handleLogin(c *gin.Context) {
	var login Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if login.Username != "c137@onecause.com" || login.Password != "#th@nH@rm#y#r!$100%D0p#" || !tokenValid(c, login.Token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

// tokenValid checks if the token is valid.
func tokenValid(c *gin.Context, s string) bool {
	// todo implement token validation
	return true
}
