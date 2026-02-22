package models

import (
	"time"
	"github.com/shopspring/decimal"
)

// StationReading represents the 'station_readings' table
// We use this specific decimal type because it is more accurate
// and WC values will be commonly used for arithmetic
type StationReading struct {
	ReadingID     int64           `json:"reading_id"`
	StationID     int             `json:"station_id"`
	RecordedAt    time.Time       `json:"recorded_at"`
	Precipitation decimal.Decimal `json:"precipitation"`
	WC1           decimal.Decimal `json:"wc1"`
	WC2           decimal.Decimal `json:"wc2"`
	WC3           decimal.Decimal `json:"wc3"`
	WC4           decimal.Decimal `json:"wc4"`
}
