package main

import (
	"log"
	"os"

	"mini-order-api/config"
	"mini-order-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Muat environment variable dari file .env jika di lokal
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ File .env tidak ditemukan, membaca konfigurasi dari sistem/Railway environment")
	}

	// 2. Inisialisasi koneksi Database PostgreSQL & Auto-Migrate GORM
	config.ConnectDatabase()

	// 3. Inisialisasi router Gin
	r := routes.SetupRouter()

	// 4. Railway & Cloud provider otomatis inject env PORT. Fallback ke APP_PORT atau 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
		if port == "" {
			port = "8080"
		}
	}

	log.Printf("🚀 Server aktif di port :%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
