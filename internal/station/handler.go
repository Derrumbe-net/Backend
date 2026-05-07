package station

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Derrumbe-net/Backend/internal/models"
	"github.com/Derrumbe-net/Backend/internal/utils"
	"github.com/shopspring/decimal"
)

type StationHandler struct {
	Service *StationService
}

func NewStationHandler(service *StationService) *StationHandler {
	return &StationHandler{Service: service}
}

// --- DTOs ---

type CreateStationRequest struct {
	Name                    string     `json:"name"`
	Latitude                float64    `json:"latitude"`
	Longitude               float64    `json:"longitude"`
	IsAvailable             bool       `json:"is_available"`
	StationInstallationDate *time.Time `json:"station_installation_date"`
	LandUnit                *string    `json:"land_unit"`
	GeologicalUnit          *string    `json:"geological_unit"`
	Susceptibility          *string    `json:"susceptibility"`
	Depth                   *string    `json:"depth"`
	LandslideForecast       *float64   `json:"landslide_forecast"`
	ImagePath               *string    `json:"image_url"`
	Elevation               *int       `json:"elevation"`
	Slope                   *float64   `json:"slope"`
	Collaborator            *string    `json:"collaborator"`
	WC1Max                  *float64   `json:"wc1_max"`
	WC2Max                  *float64   `json:"wc2_max"`
	WC3Max                  *float64   `json:"wc3_max"`
	WC4Max                  *float64   `json:"wc4_max"`
}

type UpdateStationRequest struct {
	Name                    *string    `json:"name"`
	Latitude                *float64   `json:"latitude"`
	Longitude               *float64   `json:"longitude"`
	IsAvailable             *bool      `json:"is_available"`
	StationInstallationDate *time.Time `json:"station_installation_date"`
	LandUnit                *string    `json:"land_unit"`
	GeologicalUnit          *string    `json:"geological_unit"`
	Susceptibility          *string    `json:"susceptibility"`
	Depth                   *string    `json:"depth"`
	LandslideForecast       *float64   `json:"landslide_forecast"`
	ImagePath               *string    `json:"image_url"`
	Elevation               *int       `json:"elevation"`
	Slope                   *float64   `json:"slope"`
	Collaborator            *string    `json:"collaborator"`
	WC1Max                  *float64   `json:"wc1_max"`
	WC2Max                  *float64   `json:"wc2_max"`
	WC3Max                  *float64   `json:"wc3_max"`
	WC4Max                  *float64   `json:"wc4_max"`
}

type StationResponse struct {
	StationID               int        `json:"station_id"`
	Name                    string     `json:"name"`
	Latitude                float64    `json:"latitude"`
	Longitude               float64    `json:"longitude"`
	IsAvailable             bool       `json:"is_available"`
	StationInstallationDate *time.Time `json:"station_installation_date"`
	LandUnit                *string    `json:"land_unit"`
	GeologicalUnit          *string    `json:"geological_unit"`
	Susceptibility          *string    `json:"susceptibility"`
	Depth                   *string    `json:"depth"`
	LandslideForecast       *float64   `json:"landslide_forecast"`
	ImagePath               *string    `json:"image_url"`
	Elevation               *int       `json:"elevation"`
	Slope                   *float64   `json:"slope"`
	Collaborator            *string    `json:"collaborator"`
	WC1Max                  *float64   `json:"wc1_max"`
	WC2Max                  *float64   `json:"wc2_max"`
	WC3Max                  *float64   `json:"wc3_max"`
	WC4Max                  *float64   `json:"wc4_max"`
}

func toStationResponse(s *models.Station) StationResponse {
	return StationResponse{
		StationID:               s.StationID,
		Name:                    s.Name,
		Latitude:                s.Latitude,
		Longitude:               s.Longitude,
		IsAvailable:             s.IsAvailable,
		StationInstallationDate: s.StationInstallationDate,
		LandUnit:                s.LandUnit,
		GeologicalUnit:          s.GeologicalUnit,
		Susceptibility:          s.Susceptibility,
		Depth:                   s.Depth,
		LandslideForecast:       s.LandslideForecast,
		ImagePath:               s.ImagePath,
		Elevation:               s.Elevation,
		Slope:                   s.Slope,
		Collaborator:            s.Collaborator,
		WC1Max:                  s.WC1Max,
		WC2Max:                  s.WC2Max,
		WC3Max:                  s.WC3Max,
		WC4Max:                  s.WC4Max,
	}
}

func toStationResponses(stations []models.Station) []StationResponse {
	var responses []StationResponse
	for _, s := range stations {
		responses = append(responses, toStationResponse(&s))
	}
	return responses
}

type CreateReadingRequest struct {
	RecordedAt    time.Time `json:"recorded_at"`
	Precipitation float64   `json:"precipitation"`
	WC1           float64   `json:"wc1"`
	WC2           float64   `json:"wc2"`
	WC3           float64   `json:"wc3"`
	WC4           float64   `json:"wc4"`
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
		Latitude:                req.Latitude,
		Longitude:               req.Longitude,
		IsAvailable:             req.IsAvailable,
		StationInstallationDate: req.StationInstallationDate,
		LandUnit:                req.LandUnit,
		GeologicalUnit:          req.GeologicalUnit,
		Susceptibility:          req.Susceptibility,
		Depth:                   req.Depth,
		LandslideForecast:       req.LandslideForecast,
		ImagePath:               req.ImagePath,
		Elevation:               req.Elevation,
		Slope:                   req.Slope,
		Collaborator:            req.Collaborator,
		WC1Max:                  req.WC1Max,
		WC2Max:                  req.WC2Max,
		WC3Max:                  req.WC3Max,
		WC4Max:                  req.WC4Max,
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

	// 1. Fetch current
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

	// 2. Decode partial request
	var req UpdateStationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// 3. Merge fields only if provided (non-nil)
	if req.Name != nil {
		s.Name = *req.Name
	}
	if req.Latitude != nil {
		s.Latitude = *req.Latitude
	}
	if req.Longitude != nil {
		s.Longitude = *req.Longitude
	}
	if req.IsAvailable != nil {
		s.IsAvailable = *req.IsAvailable
	}
	if req.StationInstallationDate != nil {
		s.StationInstallationDate = req.StationInstallationDate
	}
	if req.LandUnit != nil {
		s.LandUnit = req.LandUnit
	}
	if req.GeologicalUnit != nil {
		s.GeologicalUnit = req.GeologicalUnit
	}
	if req.Susceptibility != nil {
		s.Susceptibility = req.Susceptibility
	}
	if req.Depth != nil {
		s.Depth = req.Depth
	}
	if req.LandslideForecast != nil {
		s.LandslideForecast = req.LandslideForecast
	}
	if req.ImagePath != nil {
		s.ImagePath = req.ImagePath
	}
	if req.Elevation != nil {
		s.Elevation = req.Elevation
	}
	if req.Slope != nil {
		s.Slope = req.Slope
	}
	if req.Collaborator != nil {
		s.Collaborator = req.Collaborator
	}
	if req.WC1Max != nil {
		s.WC1Max = req.WC1Max
	}
	if req.WC2Max != nil {
		s.WC2Max = req.WC2Max
	}
	if req.WC3Max != nil {
		s.WC3Max = req.WC3Max
	}
	if req.WC4Max != nil {
		s.WC4Max = req.WC4Max
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
	// imageType := r.PathValue("type")
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	s, err := h.Service.GetStation(id)

	// Safely check the pointer for ImagePath
	if err != nil || s == nil || s.ImagePath == nil || *s.ImagePath == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Image not found in database"})
		return
	}

	// if imageType == "sensor" || imageType == "" {
	baseDir := os.Getenv("BASE_PATH")

	fullPath := filepath.Join(baseDir, *s.ImagePath)

	// Verify the file physically exists on the server
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Image file missing from disk"})
		return
	}

	http.ServeFile(w, r, fullPath)
	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": "Image type not supported"})
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

func (h *StationHandler) GetStationHistory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	var startDate, endDate *time.Time

	if sdStr := r.URL.Query().Get("start_date"); sdStr != "" {
		t, err := time.Parse(time.RFC3339, sdStr)
		if err != nil {
			// Try just date format if RFC3339 fails
			t, err = time.Parse("2006-01-02", sdStr)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid start_date format. Use RFC3339 (e.g., 2023-01-01T00:00:00Z) or YYYY-MM-DD"})
				return
			}
		}
		startDate = &t
	}

	if edStr := r.URL.Query().Get("end_date"); edStr != "" {
		t, err := time.Parse(time.RFC3339, edStr)
		if err != nil {
			// Try just date format if RFC3339 fails
			t, err = time.Parse("2006-01-02", edStr)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid end_date format. Use RFC3339 (e.g., 2023-01-01T00:00:00Z) or YYYY-MM-DD"})
				return
			}
		}
		endDate = &t
	}

	readings, err := h.Service.GetStationHistory(id, startDate, endDate)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readings)
}

func (h *StationHandler) GetLatestAllStations(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetLatestAllStations()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *StationHandler) GetLatestStation(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid station ID"})
		return
	}

	reading, err := h.Service.GetLatestStation(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reading)
}

func (h *StationHandler) GetStationImages(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	fmt.Printf("[GetStationImages] Handler called with id=%s path=%s\n", idStr, r.URL.Path)

	// Check if a specific image type is being requested
	// e.g. /stations/item/2/images/sensor or /stations/item/2/images/plot
	imageType := ""
	if strings.HasSuffix(r.URL.Path, "/sensor") {
		imageType = "sensor"
	} else if strings.HasSuffix(r.URL.Path, "/plot") {
		imageType = "plot"
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID format"})
		return
	}

	s, err := h.Service.GetStation(id)
	if err != nil || s == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Station not found"})
		return
	}

	baseDir := os.Getenv("BASE_PATH")

	// If a specific type is requested, serve the file directly
	if imageType == "sensor" {
		if s.ImagePath == nil || *s.ImagePath == "" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "No sensor image"})
			return
		}
		fullPath := filepath.Join(baseDir, *s.ImagePath)
		fmt.Printf("[GetStationImages] Serving sensor image at path=%q\n", fullPath)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Sensor image file not found"})
			return
		}
		http.ServeFile(w, r, fullPath)
		return
	}

	if imageType == "plot" {
		latest, err := h.Service.GetLatestStation(id)
		if err != nil || latest == nil || latest.ImagePath == "" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "No plot image"})
			return
		}
		fullPath := filepath.Join(baseDir, "stations", latest.ImagePath)
		fmt.Printf("[GetStationImages] Serving plot image at path=%q\n", fullPath)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Plot image file not found"})
			return
		}
		http.ServeFile(w, r, fullPath)
		return
	}

	// No type specified — return metadata JSON
	type StationImages struct {
		SensorImageURL string `json:"sensor_image_url"`
		PlotImageURL   string `json:"plot_image_url"`
	}
	result := StationImages{}

	if s.ImagePath != nil && *s.ImagePath != "" {
		fullSensorPath := filepath.Join(baseDir, *s.ImagePath)
		fmt.Printf("[GetStationImages] Checking sensor image at path=%q\n", fullSensorPath)
		if _, err := os.Stat(fullSensorPath); err != nil {
			fmt.Printf("[GetStationImages] Sensor image not found: %v\n", err)
		} else {
			result.SensorImageURL = fmt.Sprintf("/stations/item/%d/images/sensor", id)
		}
	}

	latest, err := h.Service.GetLatestStation(id)
	if err != nil {
		fmt.Printf("[GetStationImages] Error fetching latest reading: %v\n", err)
	} else if latest == nil {
		fmt.Printf("[GetStationImages] No latest reading for station_id=%d\n", id)
	} else if latest.ImagePath == "" {
		fmt.Printf("[GetStationImages] Latest reading has empty ImagePath for station_id=%d\n", id)
	} else {
		fullPlotPath := filepath.Join(baseDir, "stations", latest.ImagePath)
		fmt.Printf("[GetStationImages] Checking plot image at path=%q\n", fullPlotPath)
		if _, statErr := os.Stat(fullPlotPath); statErr != nil {
			fmt.Printf("[GetStationImages] Plot image not found: %v\n", statErr)
		} else {
			result.PlotImageURL = fmt.Sprintf("/stations/item/%d/images/plot", id)
		}
	}

	fmt.Printf("[GetStationImages] Returning metadata: %+v\n", result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *StationHandler) ServeStationImageByType(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	imageType := r.PathValue("type")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID format"})
		return
	}

	s, err := h.Service.GetStation(id)
	if err != nil || s == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Station not found"})
		return
	}

	baseDir := os.Getenv("BASE_PATH")
	if baseDir == "" {
		baseDir = "data"
	}

	var fullPath string

	switch imageType {
	case "sensor":
		// Station.ImagePath = "stations/mayaguez.jpg"
		if s.ImagePath == nil || *s.ImagePath == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "No sensor image"})
			return
		}
		fullPath = filepath.Join(baseDir, *s.ImagePath)

	case "plot":
		latest, err := h.Service.GetLatestStation(id)
		if err != nil || latest == nil || latest.ImagePath == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "No plot image"})
			return
		}
		fullPath = filepath.Join(baseDir, "stations", latest.ImagePath)

	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Image type must be 'sensor' or 'plot'"})
		return
	}

	// Prevent directory traversal
	fullPath = filepath.Join(filepath.Dir(fullPath), filepath.Base(fullPath))

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Image file not found"})
		return
	}

	http.ServeFile(w, r, fullPath)
}

func (h *StationHandler) CreateReading(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	stationID, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req CreateReadingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	reading := &models.StationReading{
		StationID:     stationID,
		RecordedAt:    req.RecordedAt,
		Precipitation: decimal.NewFromFloat(req.Precipitation),
		WC1:           decimal.NewFromFloat(req.WC1),
		WC2:           decimal.NewFromFloat(req.WC2),
		WC3:           decimal.NewFromFloat(req.WC3),
		WC4:           decimal.NewFromFloat(req.WC4),
	}

	// Assuming you add CreateReading to your StationService
	if err := h.Service.DAO.CreateReading(reading); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
