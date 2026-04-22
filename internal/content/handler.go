package content

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
    StartYear     uint16 `json:"start_year"`
    EndYear       uint16 `json:"end_year"`
    ProjectStatus string `json:"project_status"`
    ImagePath     string `json:"image_path"`
}

type UpdateProjectRequest struct {
    Title         string `json:"title"`
    Description   string `json:"description"`
    StartYear     uint16 `json:"start_year"`
    EndYear       uint16 `json:"end_year"`
    ProjectStatus string `json:"project_status"`
    ImagePath     string `json:"image_path"`
}

type ProjectResponse struct {
    ProjectID     int    `json:"project_id"`
    Title         string `json:"title"`
    Description   string `json:"description"`
    StartYear     uint16 `json:"start_year"`
    EndYear       uint16 `json:"end_year"`
    ProjectStatus string `json:"project_status"`
    ImagePath     string `json:"image_path"`
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
    Title          string     `json:"title"`
    Description    string     `json:"description"`
    PublicationURL string     `json:"publication_url"`
    ImagePath      string     `json:"image_path"`
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
    Name         string `json:"name"`
    WebsiteURL   string `json:"website_url"`
    ImagePath    string `json:"image_path"`
    DisplayOrder int    `json:"display_order"`
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
    Name        string `json:"name"`
    FacultyRole string `json:"faculty_role"`
    Email       string `json:"email"`
    Phone       string `json:"phone"`
    Extension   string `json:"extension"`
    LinkedinURL string `json:"linkedin_url"`
    ImagePath   string `json:"image_path"`
}

type CreateStudentMemberRequest struct {
    Name        string `json:"name"`
    StudentType string `json:"student_type"`
}

type UpdateStudentMemberRequest struct {
    Name        string `json:"name"`
    StudentType string `json:"student_type"`
}

type UpdateOfficeInfoRequest struct {
    Email          string `json:"email"`
    Phone          string `json:"phone"`
    PhoneExt       string `json:"phone_ext"`
    OfficeLocation string `json:"office_location"`
    FacebookURL    string `json:"facebook_url"`
}

// --- Handlers ---

// Projects
func (h *ContentHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
    projects, err := h.Service.GetAllProjects()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(toProjectResponses(projects))
}

func (h *ContentHandler) GetProject(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    p, err := h.Service.GetProject(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if p == nil {
        http.Error(w, "Project not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(toProjectResponse(p))
}

func (h *ContentHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
    var req CreateProjectRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    p := &models.Project{
        Title:         req.Title,
        Description:   req.Description,
        StartYear:     req.StartYear,
        EndYear:       req.EndYear,
        ProjectStatus: req.ProjectStatus,
        ImagePath:     req.ImagePath,
    }

    id, err := h.Service.CreateProject(p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    p.ProjectID = int(id)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(toProjectResponse(p))
}

func (h *ContentHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    var req UpdateProjectRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    p := &models.Project{
        ProjectID:     id,
        Title:         req.Title,
        Description:   req.Description,
        StartYear:     req.StartYear,
        EndYear:       req.EndYear,
        ProjectStatus: req.ProjectStatus,
        ImagePath:     req.ImagePath,
    }

    if err := h.Service.UpdateProject(p); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(toProjectResponse(p))
}

func (h *ContentHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeleteProject(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *ContentHandler) ServeProjectImage(w http.ResponseWriter, r *http.Request) {
    filename := r.PathValue("filename")
    if filename == "" {
        idStr := r.PathValue("id")
        id, _ := strconv.Atoi(idStr)
        p, err := h.Service.GetProject(id)
        if err != nil || p == nil || p.ImagePath == "" {
            http.Error(w, "Image not found", http.StatusNotFound)
            return
        }
        http.ServeFile(w, r, p.ImagePath)
        return
    }

    path := filepath.Join("uploads", "projects", filename)
    if _, err := os.Stat(path); os.IsNotExist(err) {
        http.Error(w, "Image not found", http.StatusNotFound)
        return
    }
    http.ServeFile(w, r, path)
}

func (h *ContentHandler) UploadProjectImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    destDir := filepath.Join("uploads", "projects")
    path, err := utils.UploadFile(r, "image", destDir, "")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := h.Service.UpdateProjectImage(id, path); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

// Publications
func (h *ContentHandler) GetAllPublications(w http.ResponseWriter, r *http.Request) {
    publications, err := h.Service.GetAllPublications()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(toPublicationResponses(publications))
}

func (h *ContentHandler) GetPublication(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    p, err := h.Service.GetPublication(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if p == nil {
        http.Error(w, "Publication not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(toPublicationResponse(p))
}

func (h *ContentHandler) CreatePublication(w http.ResponseWriter, r *http.Request) {
    var req CreatePublicationRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
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
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    p.PublicationID = int(id)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(toPublicationResponse(p))
}

func (h *ContentHandler) UpdatePublication(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    var req UpdatePublicationRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    p := &models.Publication{
        PublicationID:  id,
        Title:          req.Title,
        Description:    req.Description,
        PublicationURL: req.PublicationURL,
        ImagePath:      req.ImagePath,
        PublishedDate:  req.PublishedDate,
    }

    if err := h.Service.UpdatePublication(p); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(toPublicationResponse(p))
}

func (h *ContentHandler) DeletePublication(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeletePublication(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *ContentHandler) ServePublicationImage(w http.ResponseWriter, r *http.Request) {
    filename := r.PathValue("filename")
    if filename == "" {
        idStr := r.PathValue("id")
        id, _ := strconv.Atoi(idStr)
        p, err := h.Service.GetPublication(id)
        if err != nil || p == nil || p.ImagePath == "" {
            http.Error(w, "Image not found", http.StatusNotFound)
            return
        }
        http.ServeFile(w, r, p.ImagePath)
        return
    }

    path := filepath.Join("uploads", "publications", filename)
    if _, err := os.Stat(path); os.IsNotExist(err) {
        http.Error(w, "Image not found", http.StatusNotFound)
        return
    }
    http.ServeFile(w, r, path)
}

func (h *ContentHandler) UploadPublicationImage(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)

    destDir := filepath.Join("uploads", "publications")
    path, err := utils.UploadFile(r, "image", destDir, "")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := h.Service.UpdatePublicationImage(id, path); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

// Funding Sources
func (h *ContentHandler) GetAllFundingSources(w http.ResponseWriter, r *http.Request) {
    sources, err := h.Service.GetAllFundingSources()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(sources)
}

func (h *ContentHandler) GetFundingSource(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    fs := h.Service.GetFundingSource(id)
    if fs == nil {
        http.Error(w, "Funding source not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(fs)
}

func (h *ContentHandler) CreateFundingSource(w http.ResponseWriter, r *http.Request) {
    var req CreateFundingSourceRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
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
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fs.FundingID = int(id)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(fs)
}

func (h *ContentHandler) UpdateFundingSource(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    var req UpdateFundingSourceRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    fs := &models.FundingSource{
        FundingID:    id,
        Name:         req.Name,
        WebsiteURL:   req.WebsiteURL,
        ImagePath:    req.ImagePath,
        DisplayOrder: req.DisplayOrder,
    }
    if err := h.Service.UpdateFundingSource(fs); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(fs)
}

func (h *ContentHandler) DeleteFundingSource(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeleteFundingSource(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
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
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := h.Service.UpdateFundingSourceImage(id, path); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

// Faculty Members
func (h *ContentHandler) GetAllFacultyMembers(w http.ResponseWriter, r *http.Request) {
    members, err := h.Service.GetAllFacultyMembers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(members)
}

func (h *ContentHandler) GetFacultyMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    fm, err := h.Service.GetFacultyMember(id)
    if err != nil || fm == nil {
        http.Error(w, "Faculty member not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(fm)
}

func (h *ContentHandler) CreateFacultyMember(w http.ResponseWriter, r *http.Request) {
    var req CreateFacultyMemberRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
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
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fm.FacultyMemberID = int(id)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(fm)
}

func (h *ContentHandler) UpdateFacultyMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    var req UpdateFacultyMemberRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    fm := &models.FacultyMember{
        FacultyMemberID: id,
        Name:            req.Name,
        FacultyRole:     req.FacultyRole,
        Email:           req.Email,
        Phone:           req.Phone,
        Extension:       req.Extension,
        LinkedinURL:     req.LinkedinURL,
        ImagePath:       req.ImagePath,
    }
    if err := h.Service.UpdateFacultyMember(fm); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(fm)
}

func (h *ContentHandler) DeleteFacultyMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeleteFacultyMember(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
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
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := h.Service.UpdateFacultyMemberImage(id, path); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"image_path": path})
}

// Student Members
func (h *ContentHandler) GetAllStudentMembers(w http.ResponseWriter, r *http.Request) {
    members, err := h.Service.GetAllStudentMembers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(members)
}

func (h *ContentHandler) GetStudentMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    sm, err := h.Service.GetStudentMember(id)
    if err != nil || sm == nil {
        http.Error(w, "Student member not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(sm)
}

func (h *ContentHandler) CreateStudentMember(w http.ResponseWriter, r *http.Request) {
    var req CreateStudentMemberRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    sm := &models.StudentMember{
        Name:        req.Name,
        StudentType: req.StudentType,
    }
    id, err := h.Service.CreateStudentMember(sm)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    sm.StudentMemberID = int(id)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(sm)
}

func (h *ContentHandler) UpdateStudentMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    var req UpdateStudentMemberRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    sm := &models.StudentMember{
        StudentMemberID: id,
        Name:            req.Name,
        StudentType:     req.StudentType,
    }
    if err := h.Service.UpdateStudentMember(sm); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(sm)
}

func (h *ContentHandler) DeleteStudentMember(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Service.DeleteStudentMember(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// Office Info
func (h *ContentHandler) GetOfficeInfo(w http.ResponseWriter, r *http.Request) {
    oi, err := h.Service.GetOfficeInfo()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if oi == nil {
        http.Error(w, "Office info not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(oi)
}

func (h *ContentHandler) UpdateOfficeInfo(w http.ResponseWriter, r *http.Request) {
    var req UpdateOfficeInfoRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    // We assume we update the only row (ID 1 for simplicity or from current latest)
    current, _ := h.Service.GetOfficeInfo()
    id := 1
    if current != nil {
        id = current.ID
    }

    oi := &models.OfficeInfo{
        ID:             id,
        Email:          req.Email,
        Phone:          req.Phone,
        PhoneExt:       req.PhoneExt,
        OfficeLocation: req.OfficeLocation,
        FacebookURL:    req.FacebookURL,
    }
    if err := h.Service.UpdateOfficeInfo(oi); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(oi)
}