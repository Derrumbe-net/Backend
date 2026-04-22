package station

import (
	"github.com/Derrumbe-net/Backend/internal/models"
)

type StationService struct {
	DAO *StationDAO
}

func NewStationService(dao *StationDAO) *StationService {
	return &StationService{DAO: dao}
}

func (s *StationService) GetAllStations() ([]models.Station, error) {
	return s.DAO.GetAllStations()
}

func (s *StationService) GetStation(id int) (*models.Station, error) {
	return s.DAO.GetStationByID(id)
}

func (s *StationService) CreateStation(st *models.Station) (int64, error) {
	return s.DAO.CreateStation(st)
}

func (s *StationService) UpdateStation(st *models.Station) error {
	return s.DAO.UpdateStation(st)
}

func (s *StationService) UpdateStationSensorImage(id int, path string) error {
	return s.DAO.UpdateStationSensorImage(id, path)
}

func (s *StationService) DeleteStation(id int) error {
	return s.DAO.DeleteStation(id)
}

func (s *StationService) GetStationWcHistory(stationID int) ([]models.StationReading, error) {
	return s.DAO.GetReadingsHistory(stationID)
}
