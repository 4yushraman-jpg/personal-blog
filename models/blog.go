package models

import "time"

type CreateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetBlogsRequest struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type GetBlogRequest struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

type UpdateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
