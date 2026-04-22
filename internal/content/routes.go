package content

import (
	"net/http"
	"github.com/Derrumbe-net/Backend/internal/auth"
)

func RegisterRoutes(mux *http.ServeMux, handler *ContentHandler) {
	protected := func(h http.HandlerFunc) http.Handler {
		return auth.JWTMiddleware(h)
	}

	// Project Public Routes
	mux.HandleFunc("GET /projects", handler.GetAllProjects)
	mux.HandleFunc("GET /projects/item/{id}", handler.GetProject)
	mux.HandleFunc("GET /projects/item/{id}/image", handler.ServeProjectImage)
	mux.HandleFunc("GET /projects/images/{filename}", handler.ServeProjectImage)

	// Project Protected Routes
	mux.Handle("POST /projects", protected(handler.CreateProject))
	mux.Handle("PUT /projects/item/{id}", protected(handler.UpdateProject))
	mux.Handle("DELETE /projects/item/{id}", protected(handler.DeleteProject))
	mux.Handle("POST /projects/item/{id}/image", protected(handler.UploadProjectImage))

	// Publication Public Routes
	mux.HandleFunc("GET /publications", handler.GetAllPublications)
	mux.HandleFunc("GET /publications/item/{id}", handler.GetPublication)
	mux.HandleFunc("GET /publications/item/{id}/image", handler.ServePublicationImage)
	mux.HandleFunc("GET /publications/images/{filename}", handler.ServePublicationImage)

	// Publication Protected Routes
	mux.Handle("POST /publications", protected(handler.CreatePublication))
	mux.Handle("PUT /publications/item/{id}", protected(handler.UpdatePublication))
	mux.Handle("DELETE /publications/item/{id}", protected(handler.DeletePublication))
	mux.Handle("POST /publications/item/{id}/image", protected(handler.UploadPublicationImage))

	// Funding Sources
	mux.HandleFunc("GET /funding-sources", handler.GetAllFundingSources)
	mux.HandleFunc("GET /funding-sources/item/{id}", handler.GetFundingSource)
	mux.Handle("POST /funding-sources", protected(handler.CreateFundingSource))
	mux.Handle("PUT /funding-sources/item/{id}", protected(handler.UpdateFundingSource))
	mux.Handle("DELETE /funding-sources/item/{id}", protected(handler.DeleteFundingSource))
	mux.Handle("POST /funding-sources/item/{id}/image", protected(handler.UploadFundingSourceImage))

	// Faculty Members
	mux.HandleFunc("GET /faculty-members", handler.GetAllFacultyMembers)
	mux.HandleFunc("GET /faculty-members/item/{id}", handler.GetFacultyMember)
	mux.Handle("POST /faculty-members", protected(handler.CreateFacultyMember))
	mux.Handle("PUT /faculty-members/item/{id}", protected(handler.UpdateFacultyMember))
	mux.Handle("DELETE /faculty-members/item/{id}", protected(handler.DeleteFacultyMember))
	mux.Handle("POST /faculty-members/item/{id}/image", protected(handler.UploadFacultyMemberImage))

	// Student Members
	mux.HandleFunc("GET /student-members", handler.GetAllStudentMembers)
	mux.HandleFunc("GET /student-members/item/{id}", handler.GetStudentMember)
	mux.Handle("POST /student-members", protected(handler.CreateStudentMember))
	mux.Handle("PUT /student-members/item/{id}", protected(handler.UpdateStudentMember))
	mux.Handle("DELETE /student-members/item/{id}", protected(handler.DeleteStudentMember))

	// Office Info
	mux.HandleFunc("GET /office-info", handler.GetOfficeInfo)
	mux.Handle("PUT /office-info", protected(handler.UpdateOfficeInfo))
}
