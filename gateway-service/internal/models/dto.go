package models

type RegisterDTO struct {
	Login    string `json:"login" binding:"required,min=5,max=30"`
	Password string `json:"password" binding:"required,min=5,max=50"`
}
type LoginDTO struct {
	Login    string `json:"login" binding:"required,min=5,max=30"`
	Password string `json:"password" binding:"required,min=5,max=50"`
}

type QuestionDTO struct {
	Question string `json:"question" binding:"required,max=500"`
}

type UpdateDocumentByIDDTO struct {
	DocumentID int64  `json:"document_id" binding:"required"`
	Promt      string `json:"promt"  binding:"required,max=500"`
}

type AddDocumentDTO struct {
	Name string `json:"name" binding:"required,max=255"`
}
