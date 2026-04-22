package landslide

import (
	"database/sql"
	"github.com/Derrumbe-net/Backend/internal/models"
)

type LandslideDAO struct {
	DB *sql.DB
}

func NewLandslideDAO(db *sql.DB) *LandslideDAO {
	return &LandslideDAO{DB: db}
}

func (dao *LandslideDAO) CreateLandslide(l *models.Landslide) (int64, error) {
	query := "INSERT INTO landslides (landslide_date, latitude, longitude, image_path) VALUES (?, ?, ?, ?)"
	res, err := dao.DB.Exec(query, l.LandslideDate, l.Latitude, l.Longitude, l.ImagePath)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (dao *LandslideDAO) GetLandslideByID(id int) (*models.Landslide, error) {
	var l models.Landslide
	query := "SELECT landslide_id, landslide_date, latitude, longitude, image_path FROM landslides WHERE landslide_id = ?"
	err := dao.DB.QueryRow(query, id).Scan(&l.LandslideID, &l.LandslideDate, &l.Latitude, &l.Longitude, &l.ImagePath)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func (dao *LandslideDAO) GetAllLandslides() ([]models.Landslide, error) {
	query := "SELECT landslide_id, landslide_date, latitude, longitude, image_path FROM landslides"
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var landslides []models.Landslide
	for rows.Next() {
		var l models.Landslide
		if err := rows.Scan(&l.LandslideID, &l.LandslideDate, &l.Latitude, &l.Longitude, &l.ImagePath); err != nil {
			return nil, err
		}
		landslides = append(landslides, l)
	}
	return landslides, nil
}

func (dao *LandslideDAO) UpdateLandslide(l *models.Landslide) error {
	query := "UPDATE landslides SET landslide_date = ?, latitude = ?, longitude = ?, image_path = ? WHERE landslide_id = ?"
	_, err := dao.DB.Exec(query, l.LandslideDate, l.Latitude, l.Longitude, l.ImagePath, l.LandslideID)
	return err
}

func (dao *LandslideDAO) UpdateLandslideImage(id int, path string) error {
	query := "UPDATE landslides SET image_path = ? WHERE landslide_id = ?"
	_, err := dao.DB.Exec(query, path, id)
	return err
}

func (dao *LandslideDAO) DeleteLandslide(id int) error {
	query := "DELETE FROM landslides WHERE landslide_id = ?"
	_, err := dao.DB.Exec(query, id)
	return err
}
