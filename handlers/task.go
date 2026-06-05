package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"taskmanager/models"
)

// Request payload untuk create/update task
type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Error Response struct
type ErrorResponse struct {
	Error string `json:"error"`
}

// GetAllTasks untuk menangani GET /tasks
func GetAllTasks(w http.ResponseWriter, r *http.Request, store *models.TaskStore) {
	w.Header().Set("Content-Type", "application/json")

	tasks := store.GetAllTasks()
	if len(tasks) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `[]`)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID untuk menangani GET /tasks/{id}
func GetTaskByID(w http.ResponseWriter, r *http.Request, store *models.TaskStore) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID dari URL
	id, err := extractIDFromPath(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Task ID tidak valid"})
		return
	}

	task, exists := store.GetTaskByID(id)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Task tidak ditemukan"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// CreateTask untuk POST /tasks
func CreateTask(w http.ResponseWriter, r *http.Request, store *models.TaskStore) {
	w.Header().Set("Content-Type", "application/json")

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Gagal membaca body permintaan"})
		return
	}
	defer r.Body.Close()

	// Parse JSON
	var req TaskRequest
	if err := json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Format JSON tidak valid"})
		return
	}

	// Validate input
	if strings.TrimSpace(req.Title) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Title diperlukan"})
		return
	}

	// Membuat task
	task := store.CreateTask(req.Title, req.Description)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// UpdateTask handles PUT /tasks/{id}
func UpdateTask(w http.ResponseWriter, r *http.Request, store *models.TaskStore) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID dari URL
	id, err := extractIDFromPath(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Task ID tidak valid"})
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Gagal membaca body permintaan"})
		return
	}
	defer r.Body.Close()

	// Parse JSON
	var req TaskRequest
	if err := json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Format JSON tidak valid"})
		return
	}

	// Validate input
	if strings.TrimSpace(req.Title) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Title diperlukan"})
		return
	}

	// Update task
	task, exists := store.UpdateTask(id, req.Title, req.Description, req.Completed)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Task tidak ditemukan"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// DeleteTask handles DELETE /tasks/{id}
func DeleteTask(w http.ResponseWriter, r *http.Request, store *models.TaskStore) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL
	id, err := extractIDFromPath(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Task ID tidak valid"})
		return
	}

	// Delete task
	exists := store.DeleteTask(id)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Task tidak ditemukan"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Helper function untuk mengekstrak ID dari path URL
func extractIDFromPath(path string) (int, error) {
	// Path format: /tasks/{id}
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("invalid path format")
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format")
	}

	return id, nil
}
