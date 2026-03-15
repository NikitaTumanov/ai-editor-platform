package models

type Document struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Data   []byte `json:"data"`
	UserID int64  `json:"user_id"`
}
