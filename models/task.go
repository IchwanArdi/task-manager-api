package models

import (
	"sync"
	"time"
)

// Task untuk merepresentasikan sebuah tugas
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Description string `json:"description"`
	Priority    string   `json:"priority"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


// TaskStore untuk mengelola semua tugas (menggunakan penyimpanan memori)
type TaskStore struct {
	tasks  map[int]*Task
	nextID int
	mu     sync.RWMutex // Untuk mengelola akses concurrent ke task store
}

// NewTaskStore membuat instance baru dari TaskStore
func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

// CreateTask menambahkan tugas baru ke store
func (ts *TaskStore) CreateTask(title, description, priority string) *Task {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	task := &Task{
		ID:          ts.nextID,
		Title:       title,
		Description: description,
		Priority:    priority,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ts.tasks[ts.nextID] = task
	ts.nextID++

	return task
}

// GetAllTasks mengembalikan semua tugas yang ada di store
func (ts *TaskStore) GetAllTasks() []*Task {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	tasks := make([]*Task, 0, len(ts.tasks))
	for _, task := range ts.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetTaskByID mengembalikan tugas berdasarkan ID
func (ts *TaskStore) GetTaskByID(id int) (*Task, bool) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	task, exists := ts.tasks[id]
	return task, exists
}

// UpdateTask memperbarui tugas yang ada di store
func (ts *TaskStore) UpdateTask(id int, title, description, priority string, completed bool) (*Task, bool) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	task, exists := ts.tasks[id]
	if !exists {
		return nil, false
	}

	task.Title = title
	task.Description = description
	task.Priority = priority
	task.Completed = completed
	task.UpdatedAt = time.Now()

	return task, true
}

// DeleteTask menghapus tugas dari store berdasarkan ID
func (ts *TaskStore) DeleteTask(id int) bool {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	_, exists := ts.tasks[id]
	if !exists {
		return false
	}

	delete(ts.tasks, id)
	return true
}
