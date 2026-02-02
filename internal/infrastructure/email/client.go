package email

// EmailClient interface for email operations
type EmailClient interface {
	Send(to, subject, body string) error
}

// TODO: Implement email client (SMTP, SendGrid, etc.)

