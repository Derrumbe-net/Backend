package station

import (
	"net/http"
	"github.com/Derrumbe-net/Backend/internal/auth"
)

func RegisterRoutes(mux *http.ServeMux, handler *StationHandler) {
	protected := func(h http.HandlerFunc) http.Handler {
		return auth.JWTMiddleware(h)
	}

	// Station Public Routes
	mux.HandleFunc("GET /stations", handler.GetAllStations)
	mux.HandleFunc("GET /stations/{id}", handler.GetStation)
	mux.HandleFunc("GET /stations/{id}/image/{type}", handler.ServeStationImage)
	mux.HandleFunc("GET /stations/historical/{id}", handler.GetStationHistory)
	mux.HandleFunc("GET /stations/latest", handler.GetLatestAllStations)
	mux.HandleFunc("GET /stations/latest/{id}", handler.GetLatestStation)

	// Station Protected Routes
	mux.Handle("POST /stations", protected(handler.CreateStation))
	mux.Handle("PUT /stations/{id}", protected(handler.UpdateStation))
	mux.Handle("DELETE /stations/{id}", protected(handler.DeleteStation))
	mux.Handle("POST /stations/{id}/image/sensor", protected(handler.UploadStationSensorImage))
	mux.Handle("POST /stations/{id}/readings", protected(handler.CreateReading))
}
