package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() {
	//Connection string dari environment variable
	connStr := os.Getenv("DATABASE_URL")

	// Koneksi ke database
	db, err := sql.Open("postgres", connStr)

	// Cek error saat membuka koneksi
	if err != nil {
		log.Fatal(err)
	}
	// Pastikan koneksi ditutup saat program selesai
	defer db.Close()

	// Cek koneksi dengan ping ke database
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Koneksi ke database berhasil!")
	}
}