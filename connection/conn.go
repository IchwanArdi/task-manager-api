package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() *sql.DB {
	// Connection string dari environment variable
	connStr := os.Getenv("DATABASE_URL")

	// Koneksi ke database
	db, err := sql.Open("postgres", connStr)

	// Cek error saat membuka koneksi
	if err != nil {
		log.Fatal(err)
	}

	// Cek koneksi dengan ping ke database
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Koneksi ke database berhasil!")
	}

	// Buat tabel tasks jika belum ada
	createTableTasks(db)

	// Set global DB variable
	DB = db

	return db
}

func createTableTasks(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	priority VARCHAR(50) NOT NULL,
	completed BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Tabel 'tasks' berhasil dibuat.")
	}
}