package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"taskmanager/connection"
	"taskmanager/handlers"
	"taskmanager/models"

	"github.com/joho/godotenv"
)

func main() {
	// Load dari file .env
	godotenv.Load()

	// Inisialisasi koneksi database
	db := connection.InitDB()
	defer db.Close()

	// Inisialisasi TaskStore dengan database connection
	taskStore := models.NewTaskStore(db)

	// Create HTTP mux (router)
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"ok"}`)
	})

	// Tasks endpoints
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/tasks" {
			// Handle /tasks endpoint
			switch r.Method {
			case http.MethodGet:
				handlers.GetAllTasks(w, r, taskStore)
			case http.MethodPost:
				handlers.CreateTask(w, r, taskStore)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		} else {
			http.NotFound(w, r)
		}
	})

	// Handle /tasks/{id} - untuk GET, PUT, DELETE
	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		// Extract ID dari path
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) != 3 || pathParts[2] == "" {
			http.NotFound(w, r)
			return
		}
		switch r.Method {
		case http.MethodGet:
			handlers.GetTaskByID(w, r, taskStore)
		case http.MethodPut:
			handlers.UpdateTask(w, r, taskStore)
		case http.MethodDelete:
			handlers.DeleteTask(w, r, taskStore)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})


	// Start server
	port := ":8081"
	log.Printf("Server running on http://localhost%s\n", port)
	log.Printf("Endpoints available:\n")
	log.Printf("  GET    /health           - Health check\n")
	log.Printf("  GET    /tasks            - Get all tasks\n")
	log.Printf("  GET    /tasks/{id}       - Get specific task\n")
	log.Printf("  POST   /tasks            - Create new task\n")
	log.Printf("  PUT    /tasks/{id}       - Update task\n")
	log.Printf("  DELETE /tasks/{id}       - Delete task\n")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
