package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"personal-blog/database"
	"personal-blog/models"
	"strconv"
)

func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var blog models.CreateBlogRequest
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if blog.Title == "" || blog.Content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO posts (title, content) VALUES (?, ?)`
	_, err = database.DB.Exec(query, blog.Title, blog.Content)
	if err != nil {
		http.Error(w, "Failed to save this blog", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Blog created",
	})
}

func GetBlogsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	query := `SELECT id, title, created_at FROM posts`
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var blogs []models.GetBlogsRequest

	for rows.Next() {
		var b models.GetBlogsRequest
		err := rows.Scan(&b.ID, &b.Title, &b.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to read rows", http.StatusInternalServerError)
			return
		}
		blogs = append(blogs, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func GetBlogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	query := `SELECT title, created_at, content FROM posts WHERE id = ?`

	var blog models.GetBlogRequest

	err = database.DB.QueryRow(query, id).Scan(&blog.Title, &blog.CreatedAt, &blog.Content)
	if err == sql.ErrNoRows {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func UpdateBlogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var blog models.UpdateBlogRequest
	err = json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if blog.Title == "" || blog.Content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	query := `UPDATE posts SET title = ?, content = ? WHERE id = ?`
	result, err := database.DB.Exec(query, blog.Title, blog.Content, id)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Blog updated",
	})
}

func DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM posts WHERE id = ?`
	result, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
