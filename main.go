package main

import (
	"log"

	"mini-order-api/routes"
)

func main() {
	// Inisialisasi router
	r := routes.SetupRouter()

	// Jalankan server di port 8080
	port := ":8080"
	log.Printf("🚀 Server aktif di http://localhost%s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
