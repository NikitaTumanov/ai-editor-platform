package models

type Document struct {
	ID     uint   `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Data   string `json:"data,omitempty"`
	UserID uint   `json:"user_id,omitempty"`
}
