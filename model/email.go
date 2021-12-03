package model

type Email struct {
	Email   string
	Subject string
	Message string
	Mails   string
}

type EmailService interface {
	SendEmail(to []string, subject string, body string)
}
