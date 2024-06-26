package user

import "os"

var (
	//Production
	EMAIL_FROM    = os.Getenv("EMAIL_FROM")
	SMTP_HOST     = os.Getenv("SMTP_HOST")
	SMTP_PORT     = os.Getenv("SMTP_PORT")
	SMTP_USER     = os.Getenv("SMTP_USER")
	SMTP_PASSWORD = os.Getenv("SMTP_PASS")

// EMAIL_FROM    = "admin@admin.com"
// SMTP_HOST     = "sandbox.smtp.mailtrap.io"
// SMTP_PORT     = 2525
// SMTP_USER     = "b3f8c771a3ac6d"
// SMTP_PASSWORD = "c0faf64735c303"
)
