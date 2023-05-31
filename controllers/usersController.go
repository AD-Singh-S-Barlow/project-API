package controllers

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context){
	//get email/password of req body
var body struct{
	Email string
	Password string
}
if c.Bind(&body) != nil{
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "failed to read body",
	})
	return
}
	// hash password
bcrypt.GenerateFromPassword()
	//create user
	

}