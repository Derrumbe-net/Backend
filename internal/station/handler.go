package station

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

type StationHandler struct {
	Service *StationService
}

func NewStationHandler(service *StationService) *StationHandler {
	return &StationHandler{Service: service}
}

// --- DTOs ---

type CreateStationRequest struct {
	Name                    string    `json:"name"`
	Depth                   string    `json:"depth"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	IsAvailable             bool      `json:"is_available"`
	StationInstallationDate time.Time `json:"station_installation_date"`
	LandUnit                *string   `json:"land_unit"`
	GeologicalUnit          *string   `json:"geological_unit"`
	Susceptibility          *string   `json:"susceptibility"`
	LandslideForecast       *float64  `json:"landslide_forecast"`
	ImagePath               *string   `json:"image_path"`
	Elevation               *int      `json:"elevation"`
	Slope                   *float64  `json:"slope"`
	Collaborator            *string   `json:"collaborator"`
}

type UpdateStationRequest struct {
	Name                    string    `json:"name"`
	Depth                   string    `json:"depth"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	IsAvailable             bool      `json:"is_available"`
	StationInstallationDate time.Time `json:"station_installation_date"`
	LandUnit                *string   `json:"land_unit"`
	GeologicalUnit          *string   `json:"geological_unit"`
	Susceptibility          *string   `json:"susceptibility"`
	LandslideForecast       *float64  `json:"landslide_forecast"`
	ImagePath               *string   `json:"image_path"`
	Elevation               *int      `json:"elevation"`
	Slope                   *float64  `json:"slope"`
	Collaborator            *string   `json:"collaborator"`
}

type StationResponse struct {
	StationID               int       `json:"station_id"`
	Name                    string    `json:"name"`
	Depth                   string    `json:"depth"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	IsAvailable             bool      `json:"is_available"`
	StationInstallationDate time.Time `json:"station_installation_date"`
	LandUnit                *string   `json:"land_unit"`
	GeologicalUnit          *string   `json:"geological_unit"`
	Susceptibility          *string   `json:"susceptibility"`
	LandslideForecast       *float64  `json:"landslide_forecast"`
	ImagePath               *string   `json:"image_path"`
	Elevation               *int      `json:"elevation"`
	Slope                   *float64  `json:"slope"`
	Collaborator            *string   `json:"collaborator"`
}

func toStationResponse(s *models.Station) StationResponse {
	return StationResponse{
		StationID:               s.StationID,
		Name:                    s.Name,
		Depth:                   s.Depth,
		Latitude:                s.Latitude,
		Longitude:               s.Longitude,
		IsAvailable:             s.IsAvailable,
		StationInstallationDate: s.StationInstallationDate,
		LandUnit:                s.LandUnit,
		GeologicalUnit:          s.GeologicalUnit,
		Susceptibility:          s.Susceptibility,
		LandslideForecast:       s.LandslideForecast,
		ImagePath:               s.ImagePath,
		Elevation:               s.Elevation,
		Slope:                   s.Slope,
		Collaborator:            s.Collaborator,
	}
}

func toStationResponses(stations []models.Station) []StationResponse {
	var responses []StationResponse
	for _, s := range stations {
		responses = append(responses, toStationResponse(&s))
	}
	return responses
}

// --- Handlers ---

func (h *StationHandler) GetAllStations(w http.ResponseWriter, r *http.Request) {
	stations, err := h.Service.GetAllStations()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toStationResponses(stations))
}

func (h *StationHandler) GetStation(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	s, err := h.Service.GetStation(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if s == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Station not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toStationResponse(s))
}

func (h *StationHandler) CreateStation(w http.ResponseWriter, r *http.Request) {
	var req CreateStationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	s := &models.Station{
		Name:                    req.Name,
		Depth:                   req.Depth,
		Latitude:                req.Latitude,
		Longitude:               req.Longitude,
		IsAvailable:             req.IsAvailable,
		StationInstallationDate: req.StationInstallationDate,
		LandUnit:                req.LandUnit,
		GeologicalUnit:          req.GeologicalUnit,
		Susceptibility:          req.Susceptibility,
		LandslideForecast:       req.LandslideForecast,
		ImagePath:               req.ImagePath,
		Elevation:               req.Elevation,
		Slope:                   req.Slope,
		Collaborator:            req.Collaborator,
	}

	id, err := h.Service.CreateStation(s)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	s.StationID = int(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(toStationResponse(s))
}

func (h *StationHandler) UpdateStation(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	var req UpdateStationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	s := &models.Station{
		StationID:               id,
		Name:                    req.Name,
		Depth:                   req.Depth,
		Latitude:                req.Latitude,
		Longitude:               req.Longitude,
		IsAvailable:             req.IsAvailable,
		StationInstallationDate: req.StationInstallationDate,
		LandUnit:                req.LandUnit,
		GeologicalUnit:          req.GeologicalUnit,
		Susceptibility:          req.Susceptibility,
		LandslideForecast:       req.LandslideForecast,
		ImagePath:               req.ImagePath,
		Elevation:               req.Elevation,
		Slope:                   req.Slope,
		Collaborator:            req.Collaborator,
	}

	if err := h.Service.UpdateStation(s); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toStationResponse(s))
}

func (h *StationHandler) DeleteStation(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	if err := h.Service.DeleteStation(id); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *StationHandler) ServeStationImage(w http.ResponseWriter, r *http.Request) {
	imageType := r.PathValue("type")
	filename := r.PathValue("filename")

	if filename == "" {
		idStr := r.PathValue("id")
		id, _ := strconv.Atoi(idStr)
		s, err := h.Service.GetStation(id)

		// Safely check the pointer for ImagePath
		if err != nil || s == nil || s.ImagePath == nil || *s.ImagePath == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Image not found"})
			return
		}

		if imageType == "sensor" || imageType == "" {
			http.ServeFile(w, r, *s.ImagePath) // Serve the dereferenced pointer
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Image type not supported"})
		return
	}

	path := filepath.Join("uploads", "stations", filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Image not found"})
		return
	}
	http.ServeFile(w, r, path)
}

func (h *StationHandler) UploadStationSensorImage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	destDir := filepath.Join("uploads", "stations")
	path, err := utils.UploadFile(r, "image", destDir, "")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateStationSensorImage(id, path); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

func (h *StationHandler) GetStationWcHistory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	readings, err := h.Service.GetStationWcHistory(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readings)
}

func (h *StationHandler) GetAllStationFilesData(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetAllStationFilesData()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *StationHandler) GetStationFileData(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid station ID"})
		return
	}

	readings, err := h.Service.GetStationFileData(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readings)
}
