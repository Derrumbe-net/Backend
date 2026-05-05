package station

import (
	"time"

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

func (s *StationService) GetLatestAllStations() ([]map[string]interface{}, error) {
	stations, err := s.DAO.GetAllStations()
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0, len(stations))
	for _, st := range stations {
		reading, err := s.DAO.GetLatestReading(st.StationID)
		if err != nil {
			result = append(result, map[string]interface{}{
				"station_id": st.StationID,
				"error":      err.Error(),
			})
			continue
		}

		result = append(result, map[string]interface{}{
			"station_id": st.StationID,
			"data":       reading,
		})
	}

	return result, nil
}

func (s *StationService) GetLatestStation(id int) (*models.StationReading, error) {
	return s.DAO.GetLatestReading(id)
}

func (s *StationService) GetStationHistory(id int, startDate, endDate *time.Time) ([]models.StationReading, error) {
	return s.DAO.GetReadingsHistory(id, startDate, endDate)
}
