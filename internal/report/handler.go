package report

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
)

type ReportHandler struct {
	Service *ReportService
}

func NewReportHandler(service *ReportService) *ReportHandler {
	return &ReportHandler{Service: service}
}

// --- DTOs ---

type CreateReportRequest struct {
	LandslideID     int      `json:"landslide_id"`
	City            string   `json:"city"`
	Latitude        *float64 `json:"latitude"`
	Longitude       *float64 `json:"longitude"`
	PhysicalAddress *string  `json:"physical_address"`
	ReporterName    *string  `json:"reporter_name"`
	ReporterPhone   *string  `json:"reporter_phone"`
	ReporterEmail   *string  `json:"reporter_email"`
	Description     *string  `json:"description"`
	ImagePath       *string  `json:"image_path"`
}

type UpdateReportRequest struct {
	LandslideID     *int       `json:"landslide_id"`
	ReportedAt      *time.Time `json:"reported_at"`
	City            *string    `json:"city"`
	Latitude        *float64   `json:"latitude"`
	Longitude       *float64   `json:"longitude"`
	PhysicalAddress *string    `json:"physical_address"`
	ReporterName    *string    `json:"reporter_name"`
	ReporterPhone   *string    `json:"reporter_phone"`
	ReporterEmail   *string    `json:"reporter_email"`
	Description     *string    `json:"description"`
	ImagePath       *string    `json:"image_path"`
	IsValidated     *bool      `json:"is_validated"`
}

type ReportResponse struct {
	ReportID        int       `json:"report_id"`
	LandslideID     int       `json:"landslide_id"`
	ReportedAt      time.Time `json:"reported_at"`
	City            string    `json:"city"`
	Latitude        *float64  `json:"latitude"`
	Longitude       *float64  `json:"longitude"`
	PhysicalAddress *string   `json:"physical_address"`
	ReporterName    *string   `json:"reporter_name"`
	ReporterPhone   *string   `json:"reporter_phone"`
	ReporterEmail   *string   `json:"reporter_email"`
	Description     *string   `json:"description"`
	ImagePath       *string   `json:"image_path"`
	IsValidated     *bool     `json:"is_validated"`
}

func toReportResponse(r *models.Report) ReportResponse {
	return ReportResponse{
		ReportID:        r.ReportID,
		LandslideID:     r.LandslideID,
		ReportedAt:      r.ReportedAt,
		City:            r.City,
		Latitude:        r.Latitude,
		Longitude:       r.Longitude,
		PhysicalAddress: r.PhysicalAddress,
		ReporterName:    r.ReporterName,
		ReporterPhone:   r.ReporterPhone,
		ReporterEmail:   r.ReporterEmail,
		Description:     r.Description,
		ImagePath:       r.ImagePath,
		IsValidated:     r.IsValidated,
	}
}

func toReportResponses(reports []models.Report) []ReportResponse {
	var responses []ReportResponse
	for _, r := range reports {
		responses = append(responses, toReportResponse(&r))
	}
	return responses
}

// --- Handlers ---

func (h *ReportHandler) GetAllReports(w http.ResponseWriter, r *http.Request) {
	reports, err := h.Service.GetAllReports()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toReportResponses(reports))
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	res, err := h.Service.GetReport(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if res == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Report not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toReportResponse(res))
}

func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var req CreateReportRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	initialValidation := false

	res := &models.Report{
		LandslideID:     req.LandslideID,
		ReportedAt:      time.Now(),
		City:            req.City,
		Latitude:        req.Latitude,
		Longitude:       req.Longitude,
		PhysicalAddress: req.PhysicalAddress,
		ReporterName:    req.ReporterName,
		ReporterPhone:   req.ReporterPhone,
		ReporterEmail:   req.ReporterEmail,
		Description:     req.Description,
		ImagePath:       req.ImagePath,
		IsValidated:     &initialValidation,
	}

	id, err := h.Service.CreateReport(res)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	res.ReportID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(toReportResponse(res))
}

func (h *ReportHandler) UpdateReport(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	// 1. Fetch current
	res, err := h.Service.GetReport(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if res == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Report not found"})
		return
	}

	// 2. Decode partial
	var req UpdateReportRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// 3. Merge
	if req.LandslideID != nil {
		res.LandslideID = *req.LandslideID
	}
	if req.ReportedAt != nil {
		res.ReportedAt = *req.ReportedAt
	}
	if req.City != nil {
		res.City = *req.City
	}
	if req.Latitude != nil {
		res.Latitude = req.Latitude
	}
	if req.Longitude != nil {
		res.Longitude = req.Longitude
	}
	if req.PhysicalAddress != nil {
		res.PhysicalAddress = req.PhysicalAddress
	}
	if req.ReporterName != nil {
		res.ReporterName = req.ReporterName
	}
	if req.ReporterPhone != nil {
		res.ReporterPhone = req.ReporterPhone
	}
	if req.ReporterEmail != nil {
		res.ReporterEmail = req.ReporterEmail
	}
	if req.Description != nil {
		res.Description = req.Description
	}
	if req.ImagePath != nil {
		res.ImagePath = req.ImagePath
	}
	if req.IsValidated != nil {
		res.IsValidated = req.IsValidated
	}

	if err := h.Service.UpdateReport(res); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toReportResponse(res))
}

func (h *ReportHandler) DeleteReport(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	if err := h.Service.DeleteReport(id); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ReportHandler) ServeReportImage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	filename := r.PathValue("filename")
	filename = strings.TrimSuffix(filename, "/") // Clean up any trailing slashes from browser

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	res, err := h.Service.GetReport(id)
	if err != nil || res == nil || res.ImagePath == nil || *res.ImagePath == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Report or image folder not found in database"})
		return
	}

	dbFolder := *res.ImagePath

	baseDir := os.Getenv("BASE_PATH")

	fullPath := filepath.Join(baseDir, "landslides", dbFolder, filename)

	info, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("[X] ERROR: File physically DOES NOT EXIST at the Absolute Path above!\n")
		} else {
			fmt.Printf("[X] ERROR: os.Stat failed with permission/other error: %v\n", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Image file missing from disk"})
		return
	}

	if info.IsDir() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Target is a directory"})
		return
	}
	http.ServeFile(w, r, fullPath)
}

func (h *ReportHandler) GetReportImages(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Printf("[1] Fetching images for Report ID -> %d\n", id)

	res, err := h.Service.GetReport(id)
	var images []string
	baseDir := os.Getenv("BASE_PATH")
	if baseDir == "" {
		baseDir = "data"
	}

	dbFolder := *res.ImagePath
	dirPath := filepath.Join(baseDir, "landslides", dbFolder)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("[X] ERROR: Could not read folder: %v\n", err)
	} else {
		for _, f := range files {
			if !f.IsDir() && len(f.Name()) > 0 && f.Name()[0] != '.' {
				images = append(images, f.Name())
			}
		}

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

func (h *ReportHandler) UploadReportImage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	res, err := h.Service.GetReport(id)
	if err != nil || res == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Report not found"})
		return
	}

	dbFolder := ""
	if res.ImagePath != nil && *res.ImagePath != "" {
		dbFolder = *res.ImagePath
	} else {
		dbFolder = idStr
	}

	baseDir := os.Getenv("BASE_PATH")

	destDir := filepath.Join(baseDir, "landslides", dbFolder)

	path, err := utils.UploadFile(r, "image_file", destDir, "")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	fileName := filepath.Base(path)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Image uploaded successfully",
		"folder":   dbFolder,
		"filename": fileName,
	})
}

func (h *ReportHandler) DeleteReportImage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	filename := r.PathValue("filename")
	filename = strings.TrimSuffix(filename, "/") // Clean up trailing slash

	id, _ := strconv.Atoi(idStr)

	if filename == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Filename is required"})
		return
	}

	res, err := h.Service.GetReport(id)
	if err != nil || res == nil || res.ImagePath == nil || *res.ImagePath == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	baseDir := os.Getenv("BASE_PATH")

	dbFolder := *res.ImagePath

	path := filepath.Join(baseDir, "landslides", dbFolder, filename)

	if err := os.Remove(path); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete image"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
