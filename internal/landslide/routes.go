package landslide

import (
	"net/http"
	"github.com/Derrumbe-net/Backend/internal/auth"
)

func RegisterRoutes(mux *http.ServeMux, handler *LandslideHandler) {
	protected := func(h http.HandlerFunc) http.Handler {
		return auth.JWTMiddleware(h)
	}

	// Landslide Public Routes
	mux.HandleFunc("GET /landslides", handler.GetAllLandslides)
	mux.HandleFunc("GET /landslides/item/{id}", handler.GetLandslide)
	mux.HandleFunc("GET /landslides/item/{id}/images", handler.GetLandslideImages)
	mux.HandleFunc("GET /landslides/item/{id}/images/{filename}", handler.ServeLandslideImage)
	mux.HandleFunc("GET /landslides/images/{filename}", handler.ServeLandslideImage)

	// Landslide Protected Routes
	mux.Handle("POST /landslides", protected(handler.CreateLandslide))
	mux.Handle("PUT /landslides/item/{id}", protected(handler.UpdateLandslide))
	mux.Handle("DELETE /landslides/item/{id}", protected(handler.DeleteLandslide))
	mux.Handle("POST /landslides/item/{id}/image", protected(handler.UploadLandslideImage))
}
