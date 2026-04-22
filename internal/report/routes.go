package report

import (
	"net/http"
	"github.com/Derrumbe-net/Backend/internal/auth"
)

func RegisterRoutes(mux *http.ServeMux, handler *ReportHandler) {
	protected := func(h http.HandlerFunc) http.Handler {
		return auth.JWTMiddleware(h)
	}

	// Report Public Routes
	mux.HandleFunc("GET /reports", handler.GetAllReports)
	mux.HandleFunc("GET /reports/item/{id}", handler.GetReport)
	mux.HandleFunc("POST /reports", handler.CreateReport)
	mux.HandleFunc("POST /reports/item/{id}/upload", handler.UploadReportImage)
	mux.HandleFunc("GET /reports/item/{id}/images", handler.GetReportImages)
	mux.HandleFunc("GET /reports/item/{id}/images/{filename}", handler.ServeReportImage)

	// Report Protected Routes
	mux.Handle("PUT /reports/item/{id}", protected(handler.UpdateReport))
	mux.Handle("DELETE /reports/item/{id}", protected(handler.DeleteReport))
	mux.Handle("DELETE /reports/item/{id}/images/{filename}", protected(handler.DeleteReportImage))
}
