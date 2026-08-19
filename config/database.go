package config

import (
	"fmt"
	"log"
	"os"

	"mini-order-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB adalah instance global GORM yang bisa diakses oleh handlers/services
var DB *gorm.DB

// ConnectDatabase menginisialisasi koneksi PostgreSQL dan menjalankan Auto-Migration
func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "mini_order_db"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5433"
	}
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}
	timezone := os.Getenv("DB_TIMEZONE")
	if timezone == "" {
		timezone = "Asia/Jakarta"
	}

	// Data Source Name (DSN) format untuk PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("❌ Gagal terhubung ke Database PostgreSQL: %v", err)
	}

	log.Println("✅ Berhasil terhubung ke Database PostgreSQL!")

	// Menjalankan Auto-Migration untuk membuat/memperbarui tabel orders
	err = DB.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("❌ Gagal menjalankan AutoMigrate: %v", err)
	}

	log.Println("✅ Auto-Migration tabel 'orders' berhasil!")
}
