package main

import (
	"github.com/Derrumbe-net/Backend/internal/radar"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	cacheDir := "./radar_cache"
	os.MkdirAll(cacheDir, os.ModePerm)

	go radar.StartWorker(cacheDir)

	mux := http.NewServeMux()

	mux.HandleFunc("/radar/frames", radar.GetFramesHandler(cacheDir))

	mux.Handle("/radar/images/", http.StripPrefix("/radar/images/", http.FileServer(http.Dir(cacheDir))))

	log.Println("Backend API running on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
