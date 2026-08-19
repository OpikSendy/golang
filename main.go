package main

import (
	"log"
	"os"

	"mini-order-api/config"
	"mini-order-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Muat environment variable dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ File .env tidak ditemukan, menggunakan konfigurasi default")
	}

	// 2. Inisialisasi koneksi Database PostgreSQL & Auto-Migrate GORM
	config.ConnectDatabase()

	// 3. Inisialisasi router Gin
	r := routes.SetupRouter()

	// 4. Tentukan port & jalankan server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server aktif di http://localhost:%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
