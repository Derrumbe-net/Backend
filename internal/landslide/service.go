package landslide

import (
	"github.com/Derrumbe-net/Backend/internal/models"
)

type LandslideService struct {
	DAO *LandslideDAO
}

func NewLandslideService(dao *LandslideDAO) *LandslideService {
	return &LandslideService{DAO: dao}
}

func (s *LandslideService) GetAllLandslides() ([]models.Landslide, error) {
	return s.DAO.GetAllLandslides()
}

func (s *LandslideService) GetLandslide(id int) (*models.Landslide, error) {
	return s.DAO.GetLandslideByID(id)
}

func (s *LandslideService) CreateLandslide(l *models.Landslide) (int64, error) {
	return s.DAO.CreateLandslide(l)
}

func (s *LandslideService) UpdateLandslide(l *models.Landslide) error {
	return s.DAO.UpdateLandslide(l)
}

func (s *LandslideService) UpdateLandslideImage(id int, path string) error {
	return s.DAO.UpdateLandslideImage(id, path)
}

func (s *LandslideService) DeleteLandslide(id int) error {
	return s.DAO.DeleteLandslide(id)
}
