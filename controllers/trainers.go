package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mansikalra23/Project/Company/models"
)

func Password(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"password.html",
		gin.H{
			"title": "Enter Username and Password",
		},
	)
}

func Show(c *gin.Context) {
	var trainees []models.Trainee
	models.DB.Find(&trainees)

	c.HTML(
		http.StatusOK,
		"show.html",
		gin.H{
			"title": "All Trainees Information",
		},
	)
	c.JSON(http.StatusOK, gin.H{"Trainees": trainees})

}

func Email(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"email.html",
		gin.H{
			"title": "Enter email of trainee",
		},
	)
}

func ShowTrainee(c *gin.Context) {
	c.Request.ParseForm()
	var trainee models.Trainee
	if err := models.DB.Where("email = ?", c.PostForm("email")).First(&trainee).Error; err != nil {
		c.HTML(
			http.StatusOK,
			"trainer.html",
			gin.H{
				"title": "Record not found",
			},
		)
		return
	}

	c.HTML(
		http.StatusOK,
		"show.html",
		gin.H{
			"title": "Trainee Information",
		},
	)

	c.JSON(http.StatusOK, gin.H{"Trainee": trainee})
}

func Update(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"update.html",
		gin.H{
			"title": "Update Trainee Information",
		},
	)

}

func Updated(c *gin.Context) {
	c.Request.ParseForm()
	var trainee models.Trainee
	if err := models.DB.Where("email = ?", c.PostForm("email")).First(&trainee).Error; err != nil {
		c.HTML(
			http.StatusOK,
			"trainer.html",
			gin.H{
				"title": "Record not found",
			},
		)
		return
	}

	trainee = models.Trainee{Fname: c.PostForm("firstname"), Lname: c.PostForm("lastname"), Branch: c.PostForm("branch"), Batch: c.PostForm("batch"), Contact: c.PostForm("contact"), Email: c.PostForm("email"), Gender: c.PostForm("gender")}
	models.DB.Model(&trainee).Where("email = ?", c.PostForm("email")).Update(trainee)

	fmt.Println("Data updated successfully!")

	c.HTML(
		http.StatusOK,
		"updated.html",
		gin.H{
			"title": "Data updated successfully!",
		},
	)

}

func Delete(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"delete.html",
		gin.H{
			"title": "Enter email of trainee to be deleted :",
		},
	)

}

func Deleted(c *gin.Context) {
	c.Request.ParseForm()
	var trainee models.Trainee
	if err := models.DB.Where("email = ?", c.PostForm("email")).First(&trainee).Error; err != nil {
		c.HTML(
			http.StatusOK,
			"trainer.html",
			gin.H{
				"title": "Record not found",
			},
		)
		return
	}
	models.DB.Where("email = ?", c.PostForm("email")).Delete(&trainee)

	c.HTML(
		http.StatusOK,
		"deleted.html",
		gin.H{
			"title": "Trainee Deleted!",
		},
	)

}
