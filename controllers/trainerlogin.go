package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mansikalra23/Project/Company/models"
)

func Register(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"register.html",
		gin.H{
			"title": "New Trainer Registration",
		},
	)
}

func Registered(c *gin.Context) {
	c.Request.ParseForm()
	trainer := models.Trainer{Tid: c.PostForm("tid"), Name: c.PostForm("name"), Gender: c.PostForm("gender"), Email: c.PostForm("email"), Password: c.PostForm("password")}

	models.DB.Create(&trainer)

	c.HTML(
		http.StatusOK,
		"registered.html",
		gin.H{
			"title": "Registered Successfully!",
		},
	)
}

func Trainer(c *gin.Context) {
	c.Request.ParseForm()
	var trainer models.Trainer
	if err := models.DB.Where("email = ?", c.PostForm("email")).First(&trainer).Error; err != nil {
		c.HTML(
			http.StatusOK,
			"password.html",
			gin.H{
				"title": "Invalid Email. Please try again.",
			},
		)
		return
	}

	if trainer.Email == c.PostForm("email") && trainer.Password == c.PostForm("password") {

		c.HTML(
			http.StatusOK,
			"trainer.html",
			gin.H{
				"title": "Welcome!",
			},
		)
	} else {
		c.HTML(
			http.StatusOK,
			"password.html",
			gin.H{
				"title": "Password incorrect. Please try again.",
			},
		)
	}
}
