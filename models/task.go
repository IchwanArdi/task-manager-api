package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Task untuk merepresentasikan sebuah tugas
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// TaskStore untuk mengelola semua tugas dengan database
type TaskStore struct {
	db *sql.DB
}

// NewTaskStore membuat instance baru dari TaskStore
func NewTaskStore(db *sql.DB) *TaskStore {
	return &TaskStore{
		db: db,
	}
}

// CreateTask menambahkan tugas baru ke database
func (ts *TaskStore) CreateTask(title, description, priority string) *Task {
	task := &Task{
		Title:       title,
		Description: description,
		Priority:    priority,
		Completed:   false,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	query := `INSERT INTO tasks (title, description, priority, completed, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	
	err := ts.db.QueryRow(query, task.Title, task.Description, task.Priority, task.Completed, task.CreatedAt, task.UpdatedAt).Scan(&task.ID)
	if err != nil {
		fmt.Println("Error creating task:", err)
		return nil
	}

	return task
}

// GetAllTasks mengembalikan semua tugas dari database
func (ts *TaskStore) GetAllTasks() []*Task {
	query := `SELECT id, title, description, priority, completed, created_at, updated_at FROM tasks ORDER BY id`
	
	rows, err := ts.db.Query(query)
	if err != nil {
		fmt.Println("Error fetching tasks:", err)
		return []*Task{}
	}
	defer rows.Close()

	tasks := make([]*Task, 0)
	for rows.Next() {
		task := &Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			fmt.Println("Error scanning task:", err)
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks
}

// GetTaskByID mengembalikan tugas berdasarkan ID dari database
func (ts *TaskStore) GetTaskByID(id int) (*Task, bool) {
	query := `SELECT id, title, description, priority, completed, created_at, updated_at FROM tasks WHERE id = $1`
	
	task := &Task{}
	err := ts.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false
		}
		fmt.Println("Error fetching task:", err)
		return nil, false
	}

	return task, true
}

// UpdateTask memperbarui tugas yang ada di database
func (ts *TaskStore) UpdateTask(id int, title, description, priority string, completed bool) (*Task, bool) {
	query := `UPDATE tasks SET title = $1, description = $2, priority = $3, completed = $4, updated_at = $5 
	WHERE id = $6 RETURNING id, title, description, priority, completed, created_at, updated_at`
	
	task := &Task{}
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	
	err := ts.db.QueryRow(query, title, description, priority, completed, updatedAt, id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false
		}
		fmt.Println("Error updating task:", err)
		return nil, false
	}

	return task, true
}

// DeleteTask menghapus tugas dari database berdasarkan ID
func (ts *TaskStore) DeleteTask(id int) bool {
	query := `DELETE FROM tasks WHERE id = $1`
	
	result, err := ts.db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return false
	}

	return rowsAffected > 0
}
