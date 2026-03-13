package models

type Document struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Data   string `json:"data"`
	UserID int64  `json:"user_id"`
}

type AuthLoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
