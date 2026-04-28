package content

import (
    "encoding/json"
    "net/http"
    "path/filepath"
    "strconv"
    "time"

    "github.com/Derrumbe-net/Backend/internal/models"
    "github.com/Derrumbe-net/Backend/internal/utils"
)

type ContentHandler struct {
    Service *ContentService
}

func NewContentHandler(service *ContentService) *ContentHandler {
    return &ContentHandler{Service: service}
}

// --- DTOs ---

type CreateProjectRequest struct {
    Title         string `json:"title"`
    Description   string `json:"description"`
    StartYear     any    `json:"start_year"`
    EndYear       any    `json:"end_year"`
    ProjectStatus string `json:"project_status"`
    ImagePath     string `json:"image_url"`
}

type UpdateProjectRequest struct {
    Title         *string `json:"title"`
    Description   *string `json:"description"`
    StartYear     *any    `json:"start_year"`
    EndYear       *any    `json:"end_year"`
    ProjectStatus *string `json:"project_status"`
    ImagePath     *string `json:"image_url"`
}

type ProjectResponse struct {
    ProjectID     int    `json:"project_id"`
    Title         string `json:"title"`
    Description   string `json:"description"`
    StartYear     uint16 `json:"start_year"`
    EndYear       uint16 `json:"end_year"`
    ProjectStatus string `json:"project_status"`
    ImagePath     string `json:"image_url"`
}

func toProjectResponse(p *models.Project) ProjectResponse {
    return ProjectResponse{
        ProjectID:     p.ProjectID,
        Title:         p.Title,
        Description:   p.Description,
        StartYear:     p.StartYear,
        EndYear:       p.EndYear,
        ProjectStatus: p.ProjectStatus,
        ImagePath:     p.ImagePath,
    }
}

func toProjectResponses(projects []models.Project) []ProjectResponse {
    var responses []ProjectResponse
    for _, p := range projects {
        responses = append(responses, toProjectResponse(&p))
    }
    return responses
}

type CreatePublicationRequest struct {
    Title          string     `json:"title"`
    Description    string     `json:"description"`
    PublicationURL string     `json:"publication_url"`
    ImagePath      string     `json:"image_path"`
    PublishedDate  *time.Time `json:"published_date"`
}

type UpdatePublicationRequest struct {
    Title          *string    `json:"title"`
    Description    *string    `json:"description"`
    PublicationURL *string    `json:"publication_url"`
    ImagePath      *string    `json:"image_path"`
    PublishedDate  *time.Time `json:"published_date"`
}

type PublicationResponse struct {
    PublicationID  int        `json:"publication_id"`
    Title          string     `json:"title"`
    Description    string     `json:"description"`
    PublicationURL string     `json:"publication_url"`
    ImagePath      string     `json:"image_path"`
    PublishedDate  *time.Time `json:"published_date"` 
}

func toPublicationResponse(p *models.Publication) PublicationResponse {
    return PublicationResponse{
        PublicationID:  p.PublicationID,
        Title:          p.Title,
        Description:    p.Description,
        PublicationURL: p.PublicationURL,
        ImagePath:      p.ImagePath,
        PublishedDate:  p.PublishedDate,
    }
}

func toPublicationResponses(publications []models.Publication) []PublicationResponse {
    var responses []PublicationResponse
    for _, p := range publications {
        responses = append(responses, toPublicationResponse(&p))
    }
    return responses
}

type CreateFundingSourceRequest struct {
    Name         string `json:"name"`
    WebsiteURL   string `json:"website_url"`
    ImagePath    string `json:"image_path"`
    DisplayOrder int    `json:"display_order"`
}

type UpdateFundingSourceRequest struct {
    Name         *string `json:"name"`
    WebsiteURL   *string `json:"website_url"`
    ImagePath    *string `json:"image_path"`
    DisplayOrder *int    `json:"display_order"`
}

type CreateFacultyMemberRequest struct {
    Name        string `json:"name"`
    FacultyRole string `json:"faculty_role"`
    Email       string `json:"email"`
    Phone       string `json:"phone"`
    Extension   string `json:"extension"`
    LinkedinURL string `json:"linkedin_url"`
    ImagePath   string `json:"image_path"`
}

type UpdateFacultyMemberRequest struct {
    Name        *string `json:"name"`
    FacultyRole *string `json:"faculty_role"`
    Email       *string `json:"email"`
    Phone       *string `json:"phone"`
    Extension   *string `json:"extension"`
    LinkedinURL *string `json:"linkedin_url"`
    ImagePath   *string `json:"image_path"`
}

type CreateStudentMemberRequest struct {
    Name        string `json:"name"`
    StudentType string `json:"student_type"`
}

type UpdateStudentMemberRequest struct {
    Name        *string `json:"name"`
    StudentType *string `json:"student_type"`
}

type UpdateOfficeInfoRequest struct {
    Email          *string `json:"email"`
    Phone          *string `json:"phone"`
    PhoneExt       *string `json:"phone_ext"`
    OfficeLocation *string `json:"office_location"`
    FacebookURL    *string `json:"facebook_url"`
}

// --- Handlers ---

// Projects
func (h *ContentHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
    projects, err := h.Service.GetAllProjects()
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(toProjectResponses(projects))
}

func (h *ContentHandler) GetProject(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    p, err := h.Service.GetProject(id)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    if p == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Project not found"})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(toProjectResponse(p))
}

func parseAnyUint16(v any) uint16 {
    if v == nil {
        return 0
    }
    switch val := v.(type) {
    case float64:
        return uint16(val)
    case int:
        return uint16(val)
    case uint16:
        return val
    case string:
        if val == "" {
            return 0
        }
        f, err := strconv.ParseUint(val, 10, 16)
        if err == nil {
            return uint16(f)
        }
    }
    return 0
}

func (h *ContentHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
    var req CreateProjectRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }

    p := &models.Project{
        Title:         req.Title,
        Description:   req.Description,
        StartYear:     parseAnyUint16(req.StartYear),
        EndYear:       parseAnyUint16(req.EndYear),
        ProjectStatus: req.ProjectStatus,
        ImagePath:     req.ImagePath,
    }

    id, err := h.Service.CreateProject(p)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    p.ProjectID = int(id)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(toProjectResponse(p))
}

func (h *ContentHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    // 1. Fetch existing project
    p, err := h.Service.GetProject(id)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    if p == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Project not found"})
        return
    }

    // 2. Decode partial request
    var req UpdateProjectRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }

    // 3. Merge fields only if provided (non-nil)
    if req.Title != nil {
        p.Title = *req.Title
    }
    if req.Description != nil {
        p.Description = *req.Description
    }
    if req.StartYear != nil {
        p.StartYear = parseAnyUint16(*req.StartYear)
    }
    if req.EndYear != nil {
        p.EndYear = parseAnyUint16(*req.EndYear)
    }
    if req.ProjectStatus != nil {
        p.ProjectStatus = *req.ProjectStatus
    }
    if req.ImagePath != nil {
        p.ImagePath = *req.ImagePath
    }

    if err := h.Service.UpdateProject(p); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(toProjectResponse(p))
}

func (h *ContentHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeleteProject(id); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *ContentHandler) ServeProjectImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    p, err := h.Service.GetProject(id)
    if err != nil || p == nil || p.ImagePath == "" {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Image not found"})
        return
    }
    http.ServeFile(w, r, p.ImagePath)
}

func (h *ContentHandler) UploadProjectImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    destDir := filepath.Join("uploads", "projects")
    path, err := utils.UploadFile(r, "image", destDir, "")
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    if err := h.Service.UpdateProjectImage(id, path); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

// Publications
func (h *ContentHandler) GetAllPublications(w http.ResponseWriter, r *http.Request) {
    publications, err := h.Service.GetAllPublications()
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(toPublicationResponses(publications))
}

func (h *ContentHandler) GetPublication(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    p, err := h.Service.GetPublication(id)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    if p == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Publication not found"})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(toPublicationResponse(p))
}

func (h *ContentHandler) CreatePublication(w http.ResponseWriter, r *http.Request) {
    var req CreatePublicationRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }

    p := &models.Publication{
        Title:          req.Title,
        Description:    req.Description,
        PublicationURL: req.PublicationURL,
        ImagePath:      req.ImagePath,
        PublishedDate:  req.PublishedDate, 
    }

    id, err := h.Service.CreatePublication(p)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    p.PublicationID = int(id)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(toPublicationResponse(p))
}

func (h *ContentHandler) UpdatePublication(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    // 1. Fetch current
    p, err := h.Service.GetPublication(id)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    if p == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Publication not found"})
        return
    }

    // 2. Decode partial request
    var req UpdatePublicationRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }

    // 3. Merge
    if req.Title != nil {
        p.Title = *req.Title
    }
    if req.Description != nil {
        p.Description = *req.Description
    }
    if req.PublicationURL != nil {
        p.PublicationURL = *req.PublicationURL
    }
    if req.ImagePath != nil {
        p.ImagePath = *req.ImagePath
    }
    if req.PublishedDate != nil {
        p.PublishedDate = req.PublishedDate
    }

    if err := h.Service.UpdatePublication(p); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(toPublicationResponse(p))
}

func (h *ContentHandler) DeletePublication(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeletePublication(id); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *ContentHandler) ServePublicationImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    p, err := h.Service.GetPublication(id)
    if err != nil || p == nil || p.ImagePath == "" {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Image not found"})
        return
    }
    http.ServeFile(w, r, p.ImagePath)
}

func (h *ContentHandler) UploadPublicationImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    destDir := filepath.Join("uploads", "publications")
    path, err := utils.UploadFile(r, "image", destDir, "")
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    if err := h.Service.UpdatePublicationImage(id, path); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

// Funding Sources
func (h *ContentHandler) GetAllFundingSources(w http.ResponseWriter, r *http.Request) {
    sources, err := h.Service.GetAllFundingSources()
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sources)
}

func (h *ContentHandler) GetFundingSource(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    fs, err := h.Service.GetFundingSource(id)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    if fs == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Funding source not found"})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(fs)
}

func (h *ContentHandler) CreateFundingSource(w http.ResponseWriter, r *http.Request) {
    var req CreateFundingSourceRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }

    fs := &models.FundingSource{
        Name:         req.Name,
        WebsiteURL:   req.WebsiteURL,
        ImagePath:    req.ImagePath,
        DisplayOrder: req.DisplayOrder,
    }

    id, err := h.Service.CreateFundingSource(fs)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    fs.FundingID = int(id)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(fs)
}

func (h *ContentHandler) UpdateFundingSource(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    // 1. Fetch current
    fs, err := h.Service.GetFundingSource(id)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    if fs == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Funding source not found"})
        return
    }

    // 2. Decode partial
    var req UpdateFundingSourceRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }

    // 3. Merge
    if req.Name != nil {
        fs.Name = *req.Name
    }
    if req.WebsiteURL != nil {
        fs.WebsiteURL = *req.WebsiteURL
    }
    if req.ImagePath != nil {
        fs.ImagePath = *req.ImagePath
    }
    if req.DisplayOrder != nil {
        fs.DisplayOrder = *req.DisplayOrder
    }

    if err := h.Service.UpdateFundingSource(fs); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(fs)
}

func (h *ContentHandler) DeleteFundingSource(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeleteFundingSource(id); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *ContentHandler) UploadFundingSourceImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    destDir := filepath.Join("uploads", "funding")
    path, err := utils.UploadFile(r, "image", destDir, "")
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    if err := h.Service.UpdateFundingSourceImage(id, path); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

func (h *ContentHandler) ServeFundingSourceImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    p, err := h.Service.GetFundingSource(id)
    if err != nil || p == nil || p.ImagePath == "" {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Image not found"})
        return
    }
    http.ServeFile(w, r, p.ImagePath)
}

// Faculty Members
func (h *ContentHandler) GetAllFacultyMembers(w http.ResponseWriter, r *http.Request) {
    members, err := h.Service.GetAllFacultyMembers()
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(members)
}

func (h *ContentHandler) GetFacultyMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    fm, err := h.Service.GetFacultyMember(id)
    if err != nil || fm == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Faculty member not found"})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(fm)
}

func (h *ContentHandler) CreateFacultyMember(w http.ResponseWriter, r *http.Request) {
    var req CreateFacultyMemberRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }
    fm := &models.FacultyMember{
        Name:        req.Name,
        FacultyRole: req.FacultyRole,
        Email:       req.Email,
        Phone:       req.Phone,
        Extension:   req.Extension,
        LinkedinURL: req.LinkedinURL,
        ImagePath:   req.ImagePath,
    }
    id, err := h.Service.CreateFacultyMember(fm)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    fm.FacultyMemberID = int(id)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(fm)
}

func (h *ContentHandler) UpdateFacultyMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    // 1. Fetch current
    fm, err := h.Service.GetFacultyMember(id)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    if fm == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Faculty member not found"})
        return
    }

    // 2. Decode partial
    var req UpdateFacultyMemberRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }

    // 3. Merge
    if req.Name != nil {
        fm.Name = *req.Name
    }
    if req.FacultyRole != nil {
        fm.FacultyRole = *req.FacultyRole
    }
    if req.Email != nil {
        fm.Email = *req.Email
    }
    if req.Phone != nil {
        fm.Phone = *req.Phone
    }
    if req.Extension != nil {
        fm.Extension = *req.Extension
    }
    if req.LinkedinURL != nil {
        fm.LinkedinURL = *req.LinkedinURL
    }
    if req.ImagePath != nil {
        fm.ImagePath = *req.ImagePath
    }

    if err := h.Service.UpdateFacultyMember(fm); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(fm)
}

func (h *ContentHandler) DeleteFacultyMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeleteFacultyMember(id); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *ContentHandler) UploadFacultyMemberImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    destDir := filepath.Join("uploads", "faculty")
    path, err := utils.UploadFile(r, "image", destDir, "")
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    if err := h.Service.UpdateFacultyMemberImage(id, path); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

func (h *ContentHandler) ServeFacultyMemberImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    p, err := h.Service.GetFacultyMember(id)
    if err != nil || p == nil || p.ImagePath == "" {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Image not found"})
        return
    }
    http.ServeFile(w, r, p.ImagePath)
}

// Student Members
func (h *ContentHandler) GetAllStudentMembers(w http.ResponseWriter, r *http.Request) {
    members, err := h.Service.GetAllStudentMembers()
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(members)
}

func (h *ContentHandler) GetStudentMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    sm, err := h.Service.GetStudentMember(id)
    if err != nil || sm == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Student member not found"})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sm)
}

func (h *ContentHandler) CreateStudentMember(w http.ResponseWriter, r *http.Request) {
    var req CreateStudentMemberRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }
    sm := &models.StudentMember{
        Name:        req.Name,
        StudentType: req.StudentType,
    }
    id, err := h.Service.CreateStudentMember(sm)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    sm.StudentMemberID = int(id)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(sm)
}

func (h *ContentHandler) UpdateStudentMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    // 1. Fetch current
    sm, err := h.Service.GetStudentMember(id)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    if sm == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Student member not found"})
        return
    }

    // 2. Decode partial
    var req UpdateStudentMemberRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }

    // 3. Merge
    if req.Name != nil {
        sm.Name = *req.Name
    }
    if req.StudentType != nil && *req.StudentType != "" { // Special check for ENUM
        sm.StudentType = *req.StudentType
    }

    if err := h.Service.UpdateStudentMember(sm); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sm)
}

func (h *ContentHandler) DeleteStudentMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeleteStudentMember(id); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// Office Info
func (h *ContentHandler) GetOfficeInfo(w http.ResponseWriter, r *http.Request) {
    oi, err := h.Service.GetOfficeInfo()
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    if oi == nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Office info not found"})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(oi)
}

func (h *ContentHandler) UpdateOfficeInfo(w http.ResponseWriter, r *http.Request) {
    // 1. Fetch current
    oi, err := h.Service.GetOfficeInfo()
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    if oi == nil {
        oi = &models.OfficeInfo{ID: 1} // Default if none exists
    }

    // 2. Decode partial
    var req UpdateOfficeInfoRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
        return
    }

    // 3. Merge
    if req.Email != nil {
        oi.Email = *req.Email
    }
    if req.Phone != nil {
        oi.Phone = *req.Phone
    }
    if req.PhoneExt != nil {
        oi.PhoneExt = *req.PhoneExt
    }
    if req.OfficeLocation != nil {
        oi.OfficeLocation = *req.OfficeLocation
    }
    if req.FacebookURL != nil {
        oi.FacebookURL = *req.FacebookURL
    }

    if err := h.Service.UpdateOfficeInfo(oi); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(oi)
}
