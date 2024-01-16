package models

import (
	"time"
)

type Post struct {
	ID          string
	Title       string
	Content     string
	Category    string
	Status      string
	CreatedDate time.Time
	UpdatedDate time.Time
}

type CreateArticle struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}

type ArticleResponse struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status" `
}
