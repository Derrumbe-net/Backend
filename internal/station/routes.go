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
	mux.HandleFunc("GET /stations/item/{id}", handler.GetStation)
	mux.HandleFunc("GET /stations/item/{id}/image/{type}", handler.ServeStationImage)
	mux.HandleFunc("GET /stations/images/{filename}", handler.ServeStationImage)
	mux.HandleFunc("GET /stations/history/{id}/wc", handler.GetStationWcHistory)

	// Station Protected Routes
	mux.Handle("POST /stations", protected(handler.CreateStation))
	mux.Handle("PUT /stations/item/{id}", protected(handler.UpdateStation))
	mux.Handle("DELETE /stations/item/{id}", protected(handler.DeleteStation))
	mux.Handle("POST /stations/item/{id}/image/sensor", protected(handler.UploadStationSensorImage))
}
