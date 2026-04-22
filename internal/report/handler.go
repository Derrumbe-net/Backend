package report

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
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(toReportResponses(reports))
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    res, err := h.Service.GetReport(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if res == nil {
        http.Error(w, "Report not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(toReportResponse(res))
}

func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request) {
    var req CreateReportRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
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
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    res.ReportID = int(id)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(toReportResponse(res))
}

func (h *ReportHandler) UpdateReport(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    var req UpdateReportRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    res := &models.Report{
        ReportID:        id,
        LandslideID:     req.LandslideID,
        ReportedAt:      req.ReportedAt,
        City:            req.City,
        Latitude:        req.Latitude,
        Longitude:       req.Longitude,
        PhysicalAddress: req.PhysicalAddress,
        ReporterName:    req.ReporterName,
        ReporterPhone:   req.ReporterPhone,
        ReporterEmail:   req.ReporterEmail,
        Description:     req.Description,
        ImagePath:       req.ImagePath,
        IsValidated:     req.IsValidated,
    }

    if err := h.Service.UpdateReport(res); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(toReportResponse(res))
}

func (h *ReportHandler) DeleteReport(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeleteReport(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *ReportHandler) ServeReportImage(w http.ResponseWriter, r *http.Request) {
    filename := r.PathValue("filename")
    if filename == "" {
        idStr := r.PathValue("id")
        id, _ := strconv.Atoi(idStr)
        res, err := h.Service.GetReport(id)
        
        // Safely check pointer value for ImagePath
        if err != nil || res == nil || res.ImagePath == nil || *res.ImagePath == "" {
            http.Error(w, "Image not found", http.StatusNotFound)
            return
        }
        http.ServeFile(w, r, *res.ImagePath)
        return
    }

    idStr := r.PathValue("id")
    path := filepath.Join("uploads", "reports", idStr, filename)
    if _, err := os.Stat(path); os.IsNotExist(err) {
        path = filepath.Join("uploads", "reports", filename)
        if _, err := os.Stat(path); os.IsNotExist(err) {
            http.Error(w, "Image not found", http.StatusNotFound)
            return
        }
    }
    http.ServeFile(w, r, path)
}

func (h *ReportHandler) UploadReportImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    destDir := filepath.Join("uploads", "reports", idStr)
    path, err := utils.UploadFile(r, "image", destDir, "")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := h.Service.UpdateReportImage(id, path); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

func (h *ReportHandler) GetReportImages(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    dir := filepath.Join("uploads", "reports", idStr)

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

func (h *ReportHandler) DeleteReportImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    filename := r.PathValue("filename")
    path := filepath.Join("uploads", "reports", idStr, filename)

    if err := os.Remove(path); err != nil {
        http.Error(w, "Failed to delete image", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}