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
	mux.HandleFunc("GET /projects/{id}", handler.GetProject)
	mux.HandleFunc("GET /projects/{id}/image", handler.ServeProjectImage)

	// Project Protected Routes
	mux.Handle("POST /projects", protected(handler.CreateProject))
	mux.Handle("PUT /projects/{id}", protected(handler.UpdateProject))
	mux.Handle("DELETE /projects/{id}", protected(handler.DeleteProject))
	mux.Handle("POST /projects/{id}/image", protected(handler.UploadProjectImage))

	// Publication Public Routes
	mux.HandleFunc("GET /publications", handler.GetAllPublications)
	mux.HandleFunc("GET /publications/{id}", handler.GetPublication)
	mux.HandleFunc("GET /publications/{id}/image", handler.ServePublicationImage)

	// Publication Protected Routes
	mux.Handle("POST /publications", protected(handler.CreatePublication))
	mux.Handle("PUT /publications/{id}", protected(handler.UpdatePublication))
	mux.Handle("DELETE /publications/{id}", protected(handler.DeletePublication))
	mux.Handle("POST /publications/{id}/image", protected(handler.UploadPublicationImage))

	// Funding Sources
	mux.HandleFunc("GET /funding-sources", handler.GetAllFundingSources)
	mux.HandleFunc("GET /funding-sources/item/{id}", handler.GetFundingSource)
	mux.HandleFunc("GET /funding-sources/item/{id}/image", handler.ServeFundingSourceImage) // NEW: Serve Funding Image

	mux.Handle("POST /funding-sources", protected(handler.CreateFundingSource))
	mux.Handle("PUT /funding-sources/{id}", protected(handler.UpdateFundingSource))
	mux.Handle("DELETE /funding-sources/{id}", protected(handler.DeleteFundingSource))
	mux.Handle("POST /funding-sources/{id}/image", protected(handler.UploadFundingSourceImage))

	// Faculty Members
	mux.HandleFunc("GET /faculty-members", handler.GetAllFacultyMembers)
	mux.HandleFunc("GET /faculty-members/item/{id}", handler.GetFacultyMember)
	mux.HandleFunc("GET /faculty-members/item/{id}/image", handler.ServeFacultyMemberImage) // NEW: Serve Faculty Image

	mux.Handle("POST /faculty-members", protected(handler.CreateFacultyMember))
	mux.Handle("PUT /faculty-members/{id}", protected(handler.UpdateFacultyMember))
	mux.Handle("DELETE /faculty-members/{id}", protected(handler.DeleteFacultyMember))
	mux.Handle("POST /faculty-members/{id}/image", protected(handler.UploadFacultyMemberImage))

	// Student Members
	mux.HandleFunc("GET /student-members", handler.GetAllStudentMembers)
	mux.HandleFunc("GET /student-members/item/{id}", handler.GetStudentMember)
	mux.HandleFunc("GET /student-members/item/{id}/image", handler.ServeStudentMemberImage) // NEW: Serve Student Image

	mux.Handle("POST /student-members", protected(handler.CreateStudentMember))
	mux.Handle("PUT /student-members/item/{id}", protected(handler.UpdateStudentMember))
	mux.Handle("DELETE /student-members/item/{id}", protected(handler.DeleteStudentMember))
	mux.Handle("POST /student-members/item/{id}/image", protected(handler.UploadStudentMemberImage)) // NEW: Upload Student Image

	// Office Info
	mux.HandleFunc("GET /office-info", handler.GetOfficeInfo)
	mux.Handle("PUT /office-info", protected(handler.UpdateOfficeInfo))
}
