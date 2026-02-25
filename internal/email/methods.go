package email

// ---------------------------------------------

type AdminWelcomeData struct {
	Email   string
	CMSLink string
}

func (s *EmailService) SendAdminWelcome(to string, data AdminWelcomeData) error {
	return s.send(to, "Welcome to DerrumbeNet", "admin_welcome.html", data)
}

// ---------------------------------------------

type EmailVerificationData struct {
	Email   string
	VerificationLink string
}

func (s *EmailService) SendEmailVerification(to string, data EmailVerificationData) error {
	return s.send(to, "Verify Your Email", "email_verification.html", data)
}

// ---------------------------------------------

type NewAdminRequestNotificationData struct {
	Email   string
	CMSLink string
}

func (s *EmailService) SendNewAdminRequestNotification(to string, data NewAdminRequestNotificationData) error {
	return s.send(to, "New Admin Signup Request", "new_admin.html", data)
}

// ---------------------------------------------

type ReportSubmittedData struct {
	ID              int
	ReportedAt      string
	City            string
	Latitude        float64
	Longitude       float64
	PhysicalAddress string
	Description     string
	ReporterName    string
	ReporterPhone   string
	ReporterEmail   string
	CMSLink         string
}

func (s *EmailService) SendReportNotification(to string, data ReportSubmittedData) error {
	return s.send(to, "New Report Submitted", "report_submitted.html", data)
}
