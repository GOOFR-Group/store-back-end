package conf

import (
	"net/smtp"

	"github.com/GOOFR-Group/store-back-end/internal/utils/env"
)

const (
	SMTPEmail    = "SMTP_EMAIL"
	SMTPPassword = "SMTP_PASSWORD"

	SMTPEmailFile    = "SMTP_EMAIL_FILE"
	SMTPPasswordFile = "SMTP_PASSWORD_FILE"
)

const (
	smtpHost    = "smtp.gmail.com"
	smtpPort    = "587"
	smtpAddress = smtpHost + ":" + smtpPort
)

type SMTPConfiguration struct {
	Email    string
	Password string
}

var smtpConfiguration SMTPConfiguration

// InitSMTP starts the environment variables required for the SMTP
func InitSMTP() {
	env.CreateEnvValueFromEnvFile(SMTPEmail, SMTPEmailFile, true)
	env.CreateEnvValueFromEnvFile(SMTPPassword, SMTPPasswordFile, true)

	smtpConfiguration = SMTPConfiguration{
		Email:    env.GetEnvOrPanic(SMTPEmail),
		Password: env.GetEnvOrPanic(SMTPPassword),
	}

}

// SMTPAuthentication retrieves the SMTP authentication
func SMTPAuthentication() smtp.Auth {
	return smtp.PlainAuth("", smtpConfiguration.Email, smtpConfiguration.Password, smtpHost)
}

// SMTPEmailAddress retrieves the SMTP email
func SMTPEmailAddress() string {
	return smtpConfiguration.Email
}

// SMTPAddress retrieves the SMTP address
func SMTPAddress() string {
	return smtpAddress
}
