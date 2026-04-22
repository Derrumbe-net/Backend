package landslide

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Derrumbe-net/Backend/internal/models"
	"github.com/Derrumbe-net/Backend/internal/utils"
)

type LandslideHandler struct {
	Service *LandslideService
}

func NewLandslideHandler(service *LandslideService) *LandslideHandler {
	return &LandslideHandler{Service: service}
}

// --- DTOs ---

type CreateLandslideRequest struct {
	LandslideDate time.Time `json:"landslide_date"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	ImagePath     string    `json:"image_path"`
}

type UpdateLandslideRequest struct {
	LandslideDate time.Time `json:"landslide_date"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	ImagePath     string    `json:"image_path"`
}

type LandslideResponse struct {
	LandslideID   int       `json:"landslide_id"`
	LandslideDate time.Time `json:"landslide_date"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	ImagePath     string    `json:"image_path"`
}

func toLandslideResponse(l *models.Landslide) LandslideResponse {
	return LandslideResponse{
		LandslideID:   l.LandslideID,
		LandslideDate: l.LandslideDate,
		Latitude:      l.Latitude,
		Longitude:     l.Longitude,
		ImagePath:     l.ImagePath,
	}
}

func toLandslideResponses(landslides []models.Landslide) []LandslideResponse {
	var responses []LandslideResponse
	for _, l := range landslides {
		responses = append(responses, toLandslideResponse(&l))
	}
	return responses
}

// --- Handlers ---

func (h *LandslideHandler) GetAllLandslides(w http.ResponseWriter, r *http.Request) {
	landslides, err := h.Service.GetAllLandslides()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(toLandslideResponses(landslides))
}

func (h *LandslideHandler) GetLandslide(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	l, err := h.Service.GetLandslide(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if l == nil {
		http.Error(w, "Landslide not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(toLandslideResponse(l))
}

func (h *LandslideHandler) CreateLandslide(w http.ResponseWriter, r *http.Request) {
	var req CreateLandslideRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.LandslideDate.IsZero() {
		req.LandslideDate = time.Now()
	}

	l := &models.Landslide{
		LandslideDate: req.LandslideDate,
		Latitude:      req.Latitude,
		Longitude:     req.Longitude,
		ImagePath:     req.ImagePath,
	}

	id, err := h.Service.CreateLandslide(l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	l.LandslideID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(toLandslideResponse(l))
}

func (h *LandslideHandler) UpdateLandslide(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	var req UpdateLandslideRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	l := &models.Landslide{
		LandslideID:   id,
		LandslideDate: req.LandslideDate,
		Latitude:      req.Latitude,
		Longitude:     req.Longitude,
		ImagePath:     req.ImagePath,
	}

	if err := h.Service.UpdateLandslide(l); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(toLandslideResponse(l))
}

func (h *LandslideHandler) DeleteLandslide(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.Service.DeleteLandslide(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *LandslideHandler) GetLandslideImages(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	// If images are stored by ID in subfolders
	dir := filepath.Join("uploads", "landslides", idStr)
	
	files, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			json.NewEncoder(w).Encode([]string{})
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var images []string
	for _, f := range files {
		if !f.IsDir() {
			images = append(images, f.Name())
		}
	}
	json.NewEncoder(w).Encode(images)
}

func (h *LandslideHandler) ServeLandslideImage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	filename := r.PathValue("filename")
	
	var path string
	if idStr != "" && filename != "" {
		path = filepath.Join("uploads", "landslides", idStr, filename)
	} else if filename != "" {
		path = filepath.Join("uploads", "landslides", filename)
	} else {
		// Get by ID (main image)
		id, _ := strconv.Atoi(idStr)
		l, err := h.Service.GetLandslide(id)
		if err != nil || l == nil || l.ImagePath == "" {
			http.Error(w, "Image not found", http.StatusNotFound)
			return
		}
		path = l.ImagePath
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, path)
}

func (h *LandslideHandler) UploadLandslideImage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	destDir := filepath.Join("uploads", "landslides", idStr)
	path, err := utils.UploadFile(r, "image", destDir, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Service.UpdateLandslideImage(id, path); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}
