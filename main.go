package main

import (
	"fmt"
	"net/http"
	"personal-blog/database"
	"personal-blog/handlers"
	"personal-blog/middleware"
)

func main() {
	database.InitDB()

	http.HandleFunc("GET /api/posts", handlers.GetBlogsHandler)
	http.HandleFunc("GET /api/post/{id}", handlers.GetBlogHandler)

	http.HandleFunc("POST /api/new", middleware.AdminOnly(handlers.CreateBlogHandler))
	http.HandleFunc("PUT /api/update/{id}", middleware.AdminOnly(handlers.UpdateBlogHandler))
	http.HandleFunc("DELETE /api/delete/{id}", middleware.AdminOnly(handlers.DeleteBlogHandler))

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
