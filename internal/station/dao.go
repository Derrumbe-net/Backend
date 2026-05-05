package station

import (
	"database/sql"

	"github.com/Derrumbe-net/Backend/internal/models"
)

type StationDAO struct {
	DB *sql.DB
}

func NewStationDAO(db *sql.DB) *StationDAO {
	return &StationDAO{DB: db}
}

// Station CRUD
func (dao *StationDAO) CreateStation(s *models.Station) (int64, error) {
	query := `INSERT INTO stations (name, land_unit, geological_unit, susceptibility, depth, landslide_forecast, image_path, latitude, longitude, elevation, slope, is_available, collaborator, station_installation_date, wc1_max, wc2_max, wc3_max, wc4_max) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := dao.DB.Exec(query, s.Name, s.LandUnit, s.GeologicalUnit, s.Susceptibility, s.Depth, s.LandslideForecast, s.ImagePath, s.Latitude, s.Longitude, s.Elevation, s.Slope, s.IsAvailable, s.Collaborator, s.StationInstallationDate, s.WC1Max, s.WC2Max, s.WC3Max, s.WC4Max)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (dao *StationDAO) GetStationByID(id int) (*models.Station, error) {
	var s models.Station
	query := "SELECT station_id, name, land_unit, geological_unit, susceptibility, depth, landslide_forecast, image_path, latitude, longitude, elevation, slope, is_available, collaborator, station_installation_date, wc1_max, wc2_max, wc3_max, wc4_max FROM stations WHERE station_id = ?"
	err := dao.DB.QueryRow(query, id).Scan(&s.StationID, &s.Name, &s.LandUnit, &s.GeologicalUnit, &s.Susceptibility, &s.Depth, &s.LandslideForecast, &s.ImagePath, &s.Latitude, &s.Longitude, &s.Elevation, &s.Slope, &s.IsAvailable, &s.Collaborator, &s.StationInstallationDate, &s.WC1Max, &s.WC2Max, &s.WC3Max, &s.WC4Max)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (dao *StationDAO) GetAllStations() ([]models.Station, error) {
	query := "SELECT station_id, name, land_unit, geological_unit, susceptibility, depth, landslide_forecast, image_path, latitude, longitude, elevation, slope, is_available, collaborator, station_installation_date, wc1_max, wc2_max, wc3_max, wc4_max FROM stations"
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stations []models.Station
	for rows.Next() {
		var s models.Station
		if err := rows.Scan(&s.StationID, &s.Name, &s.LandUnit, &s.GeologicalUnit, &s.Susceptibility, &s.Depth, &s.LandslideForecast, &s.ImagePath, &s.Latitude, &s.Longitude, &s.Elevation, &s.Slope, &s.IsAvailable, &s.Collaborator, &s.StationInstallationDate, &s.WC1Max, &s.WC2Max, &s.WC3Max, &s.WC4Max); err != nil {
			return nil, err
		}
		stations = append(stations, s)
	}
	return stations, nil
}

func (dao *StationDAO) UpdateStation(s *models.Station) error {
	query := `UPDATE stations SET name = ?, land_unit = ?, geological_unit = ?, susceptibility = ?, depth = ?, landslide_forecast = ?, image_path = ?, latitude = ?, longitude = ?, elevation = ?, slope = ?, is_available = ?, collaborator = ?, station_installation_date = ?, wc1_max = ?, wc2_max = ?, wc3_max = ?, wc4_max = ? 
	          WHERE station_id = ?`
	_, err := dao.DB.Exec(query, s.Name, s.LandUnit, s.GeologicalUnit, s.Susceptibility, s.Depth, s.LandslideForecast, s.ImagePath, s.Latitude, s.Longitude, s.Elevation, s.Slope, s.IsAvailable, s.Collaborator, s.StationInstallationDate, s.WC1Max, s.WC2Max, s.WC3Max, s.WC4Max, s.StationID)
	return err
}

func (dao *StationDAO) UpdateStationSensorImage(id int, path string) error {
	query := "UPDATE stations SET image_path = ? WHERE station_id = ?"
	_, err := dao.DB.Exec(query, path, id)
	return err
}

func (dao *StationDAO) DeleteStation(id int) error {
	query := "DELETE FROM stations WHERE station_id = ?"
	_, err := dao.DB.Exec(query, id)
	return err
}

// Station Readings
func (dao *StationDAO) CreateReading(r *models.StationReading) error {
	query := "INSERT INTO station_readings (station_id, recorded_at, precipitation, wc1, wc2, wc3, wc4) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := dao.DB.Exec(query, r.StationID, r.RecordedAt, r.Precipitation, r.WC1, r.WC2, r.WC3, r.WC4)
	return err
}

func (dao *StationDAO) GetLatestReading(stationID int) (*models.StationReading, error) {
	var r models.StationReading
	query := "SELECT reading_id, station_id, recorded_at, image_path, precipitation, wc1, wc2, wc3, wc4 FROM station_readings WHERE station_id = ? ORDER BY recorded_at DESC LIMIT 1"
	err := dao.DB.QueryRow(query, stationID).Scan(&r.ReadingID, &r.StationID, &r.RecordedAt, &r.ImagePath, &r.Precipitation, &r.WC1, &r.WC2, &r.WC3, &r.WC4)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (dao *StationDAO) GetReadingsHistory(stationID int) ([]models.StationReading, error) {
	query := "SELECT reading_id, station_id, recorded_at, precipitation, wc1, wc2, wc3, wc4 FROM station_readings WHERE station_id = ? ORDER BY recorded_at ASC"
	rows, err := dao.DB.Query(query, stationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readings []models.StationReading
	for rows.Next() {
		var r models.StationReading
		if err := rows.Scan(&r.ReadingID, &r.StationID, &r.RecordedAt, &r.Precipitation, &r.WC1, &r.WC2, &r.WC3, &r.WC4); err != nil {
			return nil, err
		}
		readings = append(readings, r)
	}
	return readings, nil
}
