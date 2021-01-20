package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mansikalra23/Project/Company/models"
)

func Hello(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"hello.html",
		gin.H{
			"title": "Welcome to Company!",
		},
	)

}

func Form(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"form.html",
		gin.H{
			"title": "Trainee Information Form",
		},
	)
}

func Submit(c *gin.Context) {
	c.Request.ParseForm()

	trainee := models.Trainee{Fname: c.PostForm("firstname"), Lname: c.PostForm("lastname"), Branch: c.PostForm("branch"), Batch: c.PostForm("batch"), Contact: c.PostForm("contact"), Email: c.PostForm("email"), Gender: c.PostForm("gender")}
	models.DB.Create(&trainee)
	fmt.Println("Data successfully entered in database!")

	c.HTML(
		http.StatusOK,
		"submit.html",
		gin.H{
			"title": "Form submitted! Do you want to submit more form?",
		},
	)

}
