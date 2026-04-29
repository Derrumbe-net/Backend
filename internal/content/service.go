package content

import (
	"github.com/Derrumbe-net/Backend/internal/models"
)

type ContentService struct {
	DAO *ContentDAO
}

func NewContentService(dao *ContentDAO) *ContentService {
	return &ContentService{DAO: dao}
}

// Projects
func (s *ContentService) GetAllProjects() ([]models.Project, error) {
	return s.DAO.GetAllProjects()
}

func (s *ContentService) GetProject(id int) (*models.Project, error) {
	return s.DAO.GetProjectByID(id)
}

func (s *ContentService) CreateProject(p *models.Project) (int64, error) {
	return s.DAO.CreateProject(p)
}

func (s *ContentService) UpdateProject(p *models.Project) error {
	return s.DAO.UpdateProject(p)
}

func (s *ContentService) UpdateProjectImage(id int, path string) error {
	return s.DAO.UpdateProjectImage(id, path)
}

func (s *ContentService) DeleteProject(id int) error {
	return s.DAO.DeleteProject(id)
}

// Publications
func (s *ContentService) GetAllPublications() ([]models.Publication, error) {
	return s.DAO.GetAllPublications()
}

func (s *ContentService) GetPublication(id int) (*models.Publication, error) {
	return s.DAO.GetPublicationByID(id)
}

func (s *ContentService) CreatePublication(p *models.Publication) (int64, error) {
	return s.DAO.CreatePublication(p)
}

func (s *ContentService) UpdatePublication(p *models.Publication) error {
	return s.DAO.UpdatePublication(p)
}

func (s *ContentService) UpdatePublicationImage(id int, path string) error {
	return s.DAO.UpdatePublicationImage(id, path)
}

func (s *ContentService) DeletePublication(id int) error {
	return s.DAO.DeletePublication(id)
}

// Funding Sources
func (s *ContentService) GetAllFundingSources() ([]models.FundingSource, error) {
	return s.DAO.GetAllFundingSources()
}

func (s *ContentService) GetFundingSource(id int) (*models.FundingSource, error) {
	return s.DAO.GetFundingSourceByID(id)
}

func (s *ContentService) CreateFundingSource(fs *models.FundingSource) (int64, error) {
	return s.DAO.CreateFundingSource(fs)
}

func (s *ContentService) UpdateFundingSource(fs *models.FundingSource) error {
	return s.DAO.UpdateFundingSource(fs)
}

func (s *ContentService) UpdateFundingSourceImage(id int, path string) error {
	return s.DAO.UpdateFundingSourceImage(id, path)
}

func (s *ContentService) DeleteFundingSource(id int) error {
	return s.DAO.DeleteFundingSource(id)
}

// Faculty Members
func (s *ContentService) GetAllFacultyMembers() ([]models.FacultyMember, error) {
	return s.DAO.GetAllFacultyMembers()
}

func (s *ContentService) GetFacultyMember(id int) (*models.FacultyMember, error) {
	return s.DAO.GetFacultyMemberByID(id)
}

func (s *ContentService) CreateFacultyMember(fm *models.FacultyMember) (int64, error) {
	return s.DAO.CreateFacultyMember(fm)
}

func (s *ContentService) UpdateFacultyMember(fm *models.FacultyMember) error {
	return s.DAO.UpdateFacultyMember(fm)
}

func (s *ContentService) UpdateFacultyMemberImage(id int, path string) error {
	return s.DAO.UpdateFacultyMemberImage(id, path)
}

func (s *ContentService) DeleteFacultyMember(id int) error {
	return s.DAO.DeleteFacultyMember(id)
}

// Student Members
func (s *ContentService) GetAllStudentMembers() ([]models.StudentMember, error) {
	return s.DAO.GetAllStudentMembers()
}

func (s *ContentService) GetStudentMember(id int) (*models.StudentMember, error) {
	return s.DAO.GetStudentMemberByID(id)
}

func (s *ContentService) CreateStudentMember(sm *models.StudentMember) (int64, error) {
	return s.DAO.CreateStudentMember(sm)
}

func (s *ContentService) UpdateStudentMember(sm *models.StudentMember) error {
	return s.DAO.UpdateStudentMember(sm)
}

func (s *ContentService) DeleteStudentMember(id int) error {
	return s.DAO.DeleteStudentMember(id)
}

// Office Info
func (s *ContentService) GetOfficeInfo() (*models.OfficeInfo, error) {
	return s.DAO.GetOfficeInfo()
}

func (s *ContentService) UpdateOfficeInfo(oi *models.OfficeInfo) error {
	return s.DAO.UpdateOfficeInfo(oi)
}

func (s *ContentService) UpdateStudentMemberImage(id int, path string) error {
	return s.DAO.UpdateStudentMemberImage(id, path)
}
