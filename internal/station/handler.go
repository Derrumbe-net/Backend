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
	LandUnit                string    `json:"land_unit"`
	GeologicalUnit          string    `json:"geological_unit"`
	Susceptibility          string    `json:"susceptibility"`
	Depth                   string    `json:"depth"`
	LandslideForecast       float64   `json:"landslide_forecast"`
	ImagePath               string    `json:"image_path"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	Elevation               int       `json:"elevation"`
	Slope                   float64   `json:"slope"`
	IsAvailable             bool      `json:"is_available"`
	Collaborator            string    `json:"collaborator"`
	StationInstallationDate time.Time `json:"station_installation_date"`
}

type UpdateStationRequest struct {
	Name                    string    `json:"name"`
	LandUnit                string    `json:"land_unit"`
	GeologicalUnit          string    `json:"geological_unit"`
	Susceptibility          string    `json:"susceptibility"`
	Depth                   string    `json:"depth"`
	LandslideForecast       float64   `json:"landslide_forecast"`
	ImagePath               string    `json:"image_path"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	Elevation               int       `json:"elevation"`
	Slope                   float64   `json:"slope"`
	IsAvailable             bool      `json:"is_available"`
	Collaborator            string    `json:"collaborator"`
	StationInstallationDate time.Time `json:"station_installation_date"`
}

type StationResponse struct {
	StationID               int       `json:"station_id"`
	Name                    string    `json:"name"`
	LandUnit                string    `json:"land_unit"`
	GeologicalUnit          string    `json:"geological_unit"`
	Susceptibility          string    `json:"susceptibility"`
	Depth                   string    `json:"depth"`
	LandslideForecast       float64   `json:"landslide_forecast"`
	ImagePath               string    `json:"image_path"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	Elevation               int       `json:"elevation"`
	Slope                   float64   `json:"slope"`
	IsAvailable             bool      `json:"is_available"`
	Collaborator            string    `json:"collaborator"`
	StationInstallationDate time.Time `json:"station_installation_date"`
}

func toStationResponse(s *models.Station) StationResponse {
	return StationResponse{
		StationID:               s.StationID,
		Name:                    s.Name,
		LandUnit:                s.LandUnit,
		GeologicalUnit:          s.GeologicalUnit,
		Susceptibility:          s.Susceptibility,
		Depth:                   s.Depth,
		LandslideForecast:       s.LandslideForecast,
		ImagePath:               s.ImagePath,
		Latitude:                s.Latitude,
		Longitude:               s.Longitude,
		Elevation:               s.Elevation,
		Slope:                   s.Slope,
		IsAvailable:             s.IsAvailable,
		Collaborator:            s.Collaborator,
		StationInstallationDate: s.StationInstallationDate,
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(toStationResponses(stations))
}

func (h *StationHandler) GetStation(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	s, err := h.Service.GetStation(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if s == nil {
		http.Error(w, "Station not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(toStationResponse(s))
}

func (h *StationHandler) CreateStation(w http.ResponseWriter, r *http.Request) {
	var req CreateStationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	s := &models.Station{
		Name:                    req.Name,
		LandUnit:                req.LandUnit,
		GeologicalUnit:          req.GeologicalUnit,
		Susceptibility:          req.Susceptibility,
		Depth:                   req.Depth,
		LandslideForecast:       req.LandslideForecast,
		ImagePath:               req.ImagePath,
		Latitude:                req.Latitude,
		Longitude:               req.Longitude,
		Elevation:               req.Elevation,
		Slope:                   req.Slope,
		IsAvailable:             req.IsAvailable,
		Collaborator:            req.Collaborator,
		StationInstallationDate: req.StationInstallationDate,
	}

	id, err := h.Service.CreateStation(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.StationID = int(id)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(toStationResponse(s))
}

func (h *StationHandler) UpdateStation(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	var req UpdateStationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	s := &models.Station{
		StationID:               id,
		Name:                    req.Name,
		LandUnit:                req.LandUnit,
		GeologicalUnit:          req.GeologicalUnit,
		Susceptibility:          req.Susceptibility,
		Depth:                   req.Depth,
		LandslideForecast:       req.LandslideForecast,
		ImagePath:               req.ImagePath,
		Latitude:                req.Latitude,
		Longitude:               req.Longitude,
		Elevation:               req.Elevation,
		Slope:                   req.Slope,
		IsAvailable:             req.IsAvailable,
		Collaborator:            req.Collaborator,
		StationInstallationDate: req.StationInstallationDate,
	}

	if err := h.Service.UpdateStation(s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(toStationResponse(s))
}

func (h *StationHandler) DeleteStation(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	if err := h.Service.DeleteStation(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		if err != nil || s == nil || s.ImagePath == "" {
			http.Error(w, "Image not found", http.StatusNotFound)
			return
		}

		if imageType == "sensor" || imageType == "" {
			http.ServeFile(w, r, s.ImagePath)
			return
		}
		// If type is specified but doesn't match, we could handle other types here
		http.Error(w, "Image type not supported", http.StatusBadRequest)
		return
	}

	path := filepath.Join("uploads", "stations", filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(w, "Image not found", http.StatusNotFound)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Service.UpdateStationSensorImage(id, path); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

func (h *StationHandler) GetStationWcHistory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	readings, err := h.Service.GetStationWcHistory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(readings)
}
