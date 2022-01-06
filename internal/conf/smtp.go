package conf

import (
	"github.com/GOOFR-Group/store-back-end/internal/utils/env"
)

const (
	SMTPEmail    = "SMTP_EMAIL"
	SMTPPassword = "SMTP_PASSWORD"

	SMTPEmailFile    = "SMTP_EMAIL_FILE"
	SMTPPasswordFile = "SMTP_PASSWORD_FILE"
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

// GetSMTPConfig retrieves the SMTP configuration
func GetSMTPConfig() SMTPConfiguration {
	return smtpConfiguration
}
