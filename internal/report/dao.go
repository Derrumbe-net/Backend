package report

import (
	"database/sql"
	"github.com/Derrumbe-net/Backend/internal/models"
)

type ReportDAO struct {
	DB *sql.DB
}

func NewReportDAO(db *sql.DB) *ReportDAO {
	return &ReportDAO{DB: db}
}

func (dao *ReportDAO) CreateReport(r *models.Report) (int64, error) {
	query := `INSERT INTO reports (landslide_id, reported_at, latitude, longitude, city, physical_address, reporter_name, reporter_phone, reporter_email, description, image_path, is_validated) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	
	// Handle zero LandslideID (null in DB)
	var landslideID interface{} = r.LandslideID
	if r.LandslideID == 0 {
		landslideID = nil
	}

	res, err := dao.DB.Exec(query, landslideID, r.ReportedAt, r.Latitude, r.Longitude, r.City, r.PhysicalAddress, r.ReporterName, r.ReporterPhone, r.ReporterEmail, r.Description, r.ImagePath, r.IsValidated)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (dao *ReportDAO) GetReportByID(id int) (*models.Report, error) {
	var r models.Report
	var landslideID sql.NullInt64
	query := "SELECT * FROM reports WHERE report_id = ?"
	err := dao.DB.QueryRow(query, id).Scan(&r.ReportID, &landslideID, &r.ReportedAt, &r.Latitude, &r.Longitude, &r.City, &r.PhysicalAddress, &r.ReporterName, &r.ReporterPhone, &r.ReporterEmail, &r.Description, &r.ImagePath, &r.IsValidated)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if landslideID.Valid {
		r.LandslideID = int(landslideID.Int64)
	}
	return &r, nil
}

func (dao *ReportDAO) GetAllReports() ([]models.Report, error) {
	query := "SELECT * FROM reports"
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []models.Report
	for rows.Next() {
		var r models.Report
		var landslideID sql.NullInt64
		if err := rows.Scan(&r.ReportID, &landslideID, &r.ReportedAt, &r.Latitude, &r.Longitude, &r.City, &r.PhysicalAddress, &r.ReporterName, &r.ReporterPhone, &r.ReporterEmail, &r.Description, &r.ImagePath, &r.IsValidated); err != nil {
			return nil, err
		}
		if landslideID.Valid {
			r.LandslideID = int(landslideID.Int64)
		}
		reports = append(reports, r)
	}
	return reports, nil
}

func (dao *ReportDAO) UpdateReport(r *models.Report) error {
	query := `UPDATE reports SET landslide_id = ?, reported_at = ?, latitude = ?, longitude = ?, city = ?, physical_address = ?, reporter_name = ?, reporter_phone = ?, reporter_email = ?, description = ?, image_path = ?, is_validated = ? 
	          WHERE report_id = ?`
	
	var landslideID interface{} = r.LandslideID
	if r.LandslideID == 0 {
		landslideID = nil
	}

	_, err := dao.DB.Exec(query, landslideID, r.ReportedAt, r.Latitude, r.Longitude, r.City, r.PhysicalAddress, r.ReporterName, r.ReporterPhone, r.ReporterEmail, r.Description, r.ImagePath, r.IsValidated, r.ReportID)
	return err
}

func (dao *ReportDAO) UpdateReportImage(id int, path string) error {
	query := "UPDATE reports SET image_path = ? WHERE report_id = ?"
	_, err := dao.DB.Exec(query, path, id)
	return err
}

func (dao *ReportDAO) DeleteReport(id int) error {
	query := "DELETE FROM reports WHERE report_id = ?"
	_, err := dao.DB.Exec(query, id)
	return err
}
