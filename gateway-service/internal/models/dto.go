package models

type RegisterDTO struct {
	Login    string `json:"login" binding:"required,min=5,max=30"`
	Password string `json:"password" binding:"required,min=5,max=50"`
}
