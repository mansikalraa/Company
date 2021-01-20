package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mansikalra23/Project/Company/controllers"
	"github.com/mansikalra23/Project/Company/models"
)

var r *gin.Engine

func main() {
	r = gin.Default()

	// Connect to database
	models.ConnectDatabase()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", controllers.Hello)

	r.GET("/form", controllers.Form)
	r.POST("/submit", controllers.Submit)

	r.GET("/register", controllers.Register)
	r.POST("/registered", controllers.Registered)
	r.GET("/password", controllers.Password)
	r.POST("/trainer", controllers.Trainer)
	r.GET("/update", controllers.Update)
	r.POST("/updated", controllers.Updated)
	r.GET("/delete", controllers.Delete)
	r.POST("/deleted", controllers.Deleted)
	r.GET("/show", controllers.Show)
	r.GET("/email", controllers.Email)
	r.POST("/showtrainee", controllers.ShowTrainee)

	r.Run(":8000")

}
