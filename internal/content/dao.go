package content

import (
	"database/sql"
	"github.com/Derrumbe-net/Backend/internal/models"
)

type ContentDAO struct {
	DB *sql.DB
}

func NewContentDAO(db *sql.DB) *ContentDAO {
	return &ContentDAO{DB: db}
}

// Projects
func (dao *ContentDAO) CreateProject(p *models.Project) (int64, error) {
	query := "INSERT INTO projects (title, description, start_year, end_year, project_status, image_path) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := dao.DB.Exec(query, p.Title, p.Description, p.StartYear, p.EndYear, p.ProjectStatus, p.ImagePath)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (dao *ContentDAO) GetProjectByID(id int) (*models.Project, error) {
	var p models.Project
	query := "SELECT project_id, title, description, start_year, end_year, project_status, image_path FROM projects WHERE project_id = ?"
	err := dao.DB.QueryRow(query, id).Scan(&p.ProjectID, &p.Title, &p.Description, &p.StartYear, &p.EndYear, &p.ProjectStatus, &p.ImagePath)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (dao *ContentDAO) GetAllProjects() ([]models.Project, error) {
	query := "SELECT project_id, title, description, start_year, end_year, project_status, image_path FROM projects"
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(&p.ProjectID, &p.Title, &p.Description, &p.StartYear, &p.EndYear, &p.ProjectStatus, &p.ImagePath); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

func (dao *ContentDAO) UpdateProject(p *models.Project) error {
	query := "UPDATE projects SET title = ?, description = ?, start_year = ?, end_year = ?, project_status = ?, image_path = ? WHERE project_id = ?"
	_, err := dao.DB.Exec(query, p.Title, p.Description, p.StartYear, p.EndYear, p.ProjectStatus, p.ImagePath, p.ProjectID)
	return err
}

func (dao *ContentDAO) UpdateProjectImage(id int, path string) error {
	query := "UPDATE projects SET image_path = ? WHERE project_id = ?"
	_, err := dao.DB.Exec(query, path, id)
	return err
}

func (dao *ContentDAO) DeleteProject(id int) error {
	query := "DELETE FROM projects WHERE project_id = ?"
	_, err := dao.DB.Exec(query, id)
	return err
}

// Publications
func (dao *ContentDAO) CreatePublication(p *models.Publication) (int64, error) {
	query := "INSERT INTO publications (title, description, publication_url, image_path, published_date) VALUES (?, ?, ?, ?, ?)"
	res, err := dao.DB.Exec(query, p.Title, p.Description, p.PublicationURL, p.ImagePath, p.PublishedDate)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (dao *ContentDAO) GetPublicationByID(id int) (*models.Publication, error) {
	var p models.Publication
	query := "SELECT publication_id, title, description, publication_url, image_path, published_date FROM publications WHERE publication_id = ?"
	err := dao.DB.QueryRow(query, id).Scan(&p.PublicationID, &p.Title, &p.Description, &p.PublicationURL, &p.ImagePath, &p.PublishedDate)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (dao *ContentDAO) GetAllPublications() ([]models.Publication, error) {
	query := "SELECT publication_id, title, description, publication_url, image_path, published_date FROM publications"
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication
	for rows.Next() {
		var p models.Publication
		if err := rows.Scan(&p.PublicationID, &p.Title, &p.Description, &p.PublicationURL, &p.ImagePath, &p.PublishedDate); err != nil {
			return nil, err
		}
		publications = append(publications, p)
	}
	return publications, nil
}

func (dao *ContentDAO) UpdatePublication(p *models.Publication) error {
	query := "UPDATE publications SET title = ?, description = ?, publication_url = ?, image_path = ?, published_date = ? WHERE publication_id = ?"
	_, err := dao.DB.Exec(query, p.Title, p.Description, p.PublicationURL, p.ImagePath, p.PublishedDate, p.PublicationID)
	return err
}

func (dao *ContentDAO) UpdatePublicationImage(id int, path string) error {
	query := "UPDATE projects SET image_path = ? WHERE project_id = ?"
	_, err := dao.DB.Exec(query, path, id)
	return err
}

func (dao *ContentDAO) DeletePublication(id int) error {
	query := "DELETE FROM publications WHERE publication_id = ?"
	_, err := dao.DB.Exec(query, id)
	return err
}

// Funding Sources
func (dao *ContentDAO) CreateFundingSource(fs *models.FundingSource) (int64, error) {
	query := "INSERT INTO funding_sources (name, website_url, image_path, display_order) VALUES (?, ?, ?, ?)"
	res, err := dao.DB.Exec(query, fs.Name, fs.WebsiteURL, fs.ImagePath, fs.DisplayOrder)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (dao *ContentDAO) GetFundingSourceByID(id int) (*models.FundingSource) {
	var fs models.FundingSource
	query := "SELECT funding_id, name, website_url, image_path, display_order, created_at FROM funding_sources WHERE funding_id = ?"
	err := dao.DB.QueryRow(query, id).Scan(&fs.FundingID, &fs.Name, &fs.WebsiteURL, &fs.ImagePath, &fs.DisplayOrder, &fs.CreatedAt)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return nil
	}
	return &fs
}

func (dao *ContentDAO) GetAllFundingSources() ([]models.FundingSource, error) {
	query := "SELECT funding_id, name, website_url, image_path, display_order, created_at FROM funding_sources ORDER BY display_order ASC"
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sources []models.FundingSource
	for rows.Next() {
		var fs models.FundingSource
		if err := rows.Scan(&fs.FundingID, &fs.Name, &fs.WebsiteURL, &fs.ImagePath, &fs.DisplayOrder, &fs.CreatedAt); err != nil {
			return nil, err
		}
		sources = append(sources, fs)
	}
	return sources, nil
}

func (dao *ContentDAO) UpdateFundingSource(fs *models.FundingSource) error {
	query := "UPDATE funding_sources SET name = ?, website_url = ?, image_path = ?, display_order = ? WHERE funding_id = ?"
	_, err := dao.DB.Exec(query, fs.Name, fs.WebsiteURL, fs.ImagePath, fs.DisplayOrder, fs.FundingID)
	return err
}

func (dao *ContentDAO) UpdateFundingSourceImage(id int, path string) error {
	query := "UPDATE funding_sources SET image_path = ? WHERE funding_id = ?"
	_, err := dao.DB.Exec(query, path, id)
	return err
}

func (dao *ContentDAO) DeleteFundingSource(id int) error {
	query := "DELETE FROM funding_sources WHERE funding_id = ?"
	_, err := dao.DB.Exec(query, id)
	return err
}

// Faculty Members
func (dao *ContentDAO) CreateFacultyMember(fm *models.FacultyMember) (int64, error) {
	query := "INSERT INTO faculty_members (name, faculty_role, email, phone, extension, linkedin_url, image_path) VALUES (?, ?, ?, ?, ?, ?, ?)"
	res, err := dao.DB.Exec(query, fm.Name, fm.FacultyRole, fm.Email, fm.Phone, fm.Extension, fm.LinkedinURL, fm.ImagePath)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (dao *ContentDAO) GetFacultyMemberByID(id int) (*models.FacultyMember, error) {
	var fm models.FacultyMember
	query := "SELECT faculty_member_id, name, faculty_role, email, phone, extension, linkedin_url, image_path FROM faculty_members WHERE faculty_member_id = ?"
	err := dao.DB.QueryRow(query, id).Scan(&fm.FacultyMemberID, &fm.Name, &fm.FacultyRole, &fm.Email, &fm.Phone, &fm.Extension, &fm.LinkedinURL, &fm.ImagePath)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &fm, nil
}

func (dao *ContentDAO) GetAllFacultyMembers() ([]models.FacultyMember, error) {
	query := "SELECT faculty_member_id, name, faculty_role, email, phone, extension, linkedin_url, image_path FROM faculty_members"
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.FacultyMember
	for rows.Next() {
		var fm models.FacultyMember
		if err := rows.Scan(&fm.FacultyMemberID, &fm.Name, &fm.FacultyRole, &fm.Email, &fm.Phone, &fm.Extension, &fm.LinkedinURL, &fm.ImagePath); err != nil {
			return nil, err
		}
		members = append(members, fm)
	}
	return members, nil
}

func (dao *ContentDAO) UpdateFacultyMember(fm *models.FacultyMember) error {
	query := "UPDATE faculty_members SET name = ?, faculty_role = ?, email = ?, phone = ?, extension = ?, linkedin_url = ?, image_path = ? WHERE faculty_member_id = ?"
	_, err := dao.DB.Exec(query, fm.Name, fm.FacultyRole, fm.Email, fm.Phone, fm.Extension, fm.LinkedinURL, fm.ImagePath, fm.FacultyMemberID)
	return err
}

func (dao *ContentDAO) UpdateFacultyMemberImage(id int, path string) error {
	query := "UPDATE faculty_members SET image_path = ? WHERE faculty_member_id = ?"
	_, err := dao.DB.Exec(query, path, id)
	return err
}

func (dao *ContentDAO) DeleteFacultyMember(id int) error {
	query := "DELETE FROM faculty_members WHERE faculty_member_id = ?"
	_, err := dao.DB.Exec(query, id)
	return err
}

// Student Members
func (dao *ContentDAO) CreateStudentMember(sm *models.StudentMember) (int64, error) {
	query := "INSERT INTO student_members (name, student_type) VALUES (?, ?)"
	res, err := dao.DB.Exec(query, sm.Name, sm.StudentType)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (dao *ContentDAO) GetStudentMemberByID(id int) (*models.StudentMember, error) {
	var sm models.StudentMember
	query := "SELECT student_member_id, name, student_type FROM student_members WHERE student_member_id = ?"
	err := dao.DB.QueryRow(query, id).Scan(&sm.StudentMemberID, &sm.Name, &sm.StudentType)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &sm, nil
}

func (dao *ContentDAO) GetAllStudentMembers() ([]models.StudentMember, error) {
	query := "SELECT student_member_id, name, student_type FROM student_members"
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.StudentMember
	for rows.Next() {
		var sm models.StudentMember
		if err := rows.Scan(&sm.StudentMemberID, &sm.Name, &sm.StudentType); err != nil {
			return nil, err
		}
		members = append(members, sm)
	}
	return members, nil
}

func (dao *ContentDAO) UpdateStudentMember(sm *models.StudentMember) error {
	query := "UPDATE student_members SET name = ?, student_type = ? WHERE student_member_id = ?"
	_, err := dao.DB.Exec(query, sm.Name, sm.StudentType, sm.StudentMemberID)
	return err
}

func (dao *ContentDAO) DeleteStudentMember(id int) error {
	query := "DELETE FROM student_members WHERE student_member_id = ?"
	_, err := dao.DB.Exec(query, id)
	return err
}

// Office Info
func (dao *ContentDAO) GetOfficeInfo() (*models.OfficeInfo, error) {
	var oi models.OfficeInfo
	query := "SELECT id, email, phone, phone_ext, office_location, facebook_url, updated_at FROM office_info ORDER BY updated_at DESC LIMIT 1"
	err := dao.DB.QueryRow(query).Scan(&oi.ID, &oi.Email, &oi.Phone, &oi.PhoneExt, &oi.OfficeLocation, &oi.FacebookURL, &oi.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &oi, nil
}

func (dao *ContentDAO) UpdateOfficeInfo(oi *models.OfficeInfo) error {
	// We assume there's only one row or we update the latest
	query := "UPDATE office_info SET email = ?, phone = ?, phone_ext = ?, office_location = ?, facebook_url = ? WHERE id = ?"
	_, err := dao.DB.Exec(query, oi.Email, oi.Phone, oi.PhoneExt, oi.OfficeLocation, oi.FacebookURL, oi.ID)
	return err
}
