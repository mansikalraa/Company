package models

type Trainer struct {
	Tid      string `json:"tid" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}
