package radar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// NOAA Export URL for Puerto Rico Bounding Box (W:-68.0, S:17.5, E:-65.0, N:19.0)
	noaaURL = "https://mapservices.weather.noaa.gov/eventdriven/rest/services/radar/radar_base_reflectivity_time/ImageServer/exportImage?bbox=-68.0,17.5,-65.0,19.0&bboxSR=4326&size=1000,500&imageSR=4326&format=png32&transparent=true&f=image"
	maxAge  = 12 * time.Hour
)

// StartWorker runs in the background to fetch radar images every 5 mins
func StartWorker(cacheDir string) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	// Run immediately on startup
	fetchAndSave(cacheDir)
	cleanup(cacheDir)

	for range ticker.C {
		fetchAndSave(cacheDir)
		cleanup(cacheDir)
	}
}

func fetchAndSave(cacheDir string) {
	resp, err := http.Get(noaaURL)
	if err != nil {
		log.Printf("Radar fetch error: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Radar fetch bad status: %d", resp.StatusCode)
		return
	}

	timestamp := time.Now().UnixMilli()
	filename := fmt.Sprintf("radar_%d.png", timestamp)
	outPath := filepath.Join(cacheDir, filename)

	out, err := os.Create(outPath)
	if err != nil {
		log.Printf("Radar save error: %v", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err == nil {
		log.Printf("Saved new radar frame: %s", filename)
	}
}

func cleanup(cacheDir string) {
	files, _ := os.ReadDir(cacheDir)
	now := time.Now()

	for _, f := range files {
		info, err := f.Info()
		if err != nil {
			continue
		}
		if now.Sub(info.ModTime()) > maxAge {
			os.Remove(filepath.Join(cacheDir, f.Name()))
			log.Printf("Deleted old radar frame: %s", f.Name())
		}
	}
}

// GetFramesHandler returns a JSON list of available radar images
func GetFramesHandler(cacheDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Handle CORS for frontend dev
		w.Header().Set("Content-Type", "application/json")

		files, err := os.ReadDir(cacheDir)
		if err != nil {
			http.Error(w, "[]", http.StatusOK)
			return
		}

		type Frame struct {
			Timestamp int64  `json:"timestamp"`
			URL       string `json:"url"`
		}

		var frames []Frame
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".png") {
				name := strings.TrimPrefix(f.Name(), "radar_")
				name = strings.TrimSuffix(name, ".png")
				ts, _ := strconv.ParseInt(name, 10, 64)

				frames = append(frames, Frame{
					Timestamp: ts,
					URL:       fmt.Sprintf("/radar/images/%s", f.Name()),
				})
			}
		}

		// Sort chronological (oldest to newest)
		sort.Slice(frames, func(i, j int) bool {
			return frames[i].Timestamp < frames[j].Timestamp
		})

		json.NewEncoder(w).Encode(frames)
	}
}
