package models

var NoteRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}
