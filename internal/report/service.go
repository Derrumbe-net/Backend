package report

import (
	"github.com/Derrumbe-net/Backend/internal/email"
	"github.com/Derrumbe-net/Backend/internal/models"
)

type ReportService struct {
	DAO          *ReportDAO
	EmailService *email.EmailService
}

func NewReportService(dao *ReportDAO, emailService *email.EmailService) *ReportService {
	return &ReportService{
		DAO:          dao,
		EmailService: emailService,
	}
}

func (s *ReportService) GetAllReports() ([]models.Report, error) {
	return s.DAO.GetAllReports()
}

func (s *ReportService) GetReport(id int) (*models.Report, error) {
	return s.DAO.GetReportByID(id)
}

func (s *ReportService) CreateReport(r *models.Report) (int64, error) {
	id, err := s.DAO.CreateReport(r)
	if err != nil {
		return 0, err
	}
	r.ReportID = int(id)

	// Trigger alert to Super Admin
	if s.EmailService != nil {
		_ = s.EmailService.SendReportAlert(r)
	}

	return id, nil
}

func (s *ReportService) UpdateReport(r *models.Report) error {
	return s.DAO.UpdateReport(r)
}

func (s *ReportService) UpdateReportImage(id int, path string) error {
	return s.DAO.UpdateReportImage(id, path)
}

func (s *ReportService) DeleteReport(id int) error {
	return s.DAO.DeleteReport(id)
}
