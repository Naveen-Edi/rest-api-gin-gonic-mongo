package controllers

import (
	"log"
	"rest-api-gin-gonic-mongo/forms"
	"rest-api-gin-gonic-mongo/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var userModel = new(models.UserModel)

type UserAuthController struct{}

func (user *UserAuthController) Register(c *gin.Context) {
	var data forms.RegisterUserCommand
	if c.ShouldBind(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	profile,err := userModel.GetEmail(data.Email)
	if err == nil {
		c.JSON(406, gin.H{"error": "Already Existing Email Address !", "email" : profile.Email})
		c.Abort()
		return
	}

	userdata := forms.RegisterUserCommand{}
	userdata.Email = data.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return
	}

	userdata.Password = string(hash)
	userdata.Name = data.Name

	err = userModel.Register(userdata)
	if err != nil {
		c.JSON(406, gin.H{"message": "User could not be resgister", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "User Registered Successfully !", "data" : data})
}


func (user *UserAuthController) Login(c *gin.Context) {
	var data forms.LoginUserCommand
	if c.ShouldBind(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	profile,err := userModel.GetEmail(data.Email)
	if err != nil {
		c.JSON(406, gin.H{"error": " Invalid Credentials !"})
		c.Abort()
		return 
	}

	err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(data.Password))
	if err != nil {
		c.JSON(402, gin.H{"error": "Email or password is invalid."})
		return
	}

	token, err1 := userModel.GetJwtToken(data.Email)
	if err1 != nil {
		c.JSON(500, gin.H{"error": err1.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Login  Successfully !", "token" : token})
}

func (auth *UserAuthController) Profile(c *gin.Context) {


	user := c.MustGet("user").(models.User)
	profile,err := userModel.GetEmail(user.Email)
	if err != nil {
		c.JSON(406, gin.H{"error": " Invalid Credentials !"})
		c.Abort()
		return 
	}

	c.JSON(200, gin.H{"message": "Profile Fetched Successfully !", "profile" : profile})
}