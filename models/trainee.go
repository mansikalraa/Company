package models

type Trainee struct {
	Fname   string `json:"fname" binding:"required"`
	Lname   string `json:"lname" binding:"required"`
	Branch  string `json:"branch" binding:"required"`
	Batch   string `json:"batch" binding:"required"`
	Contact string `json:"contact" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Gender  string `json:"gender" binding:"required"`
}
