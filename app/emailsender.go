package app

import (
	"net/smtp"
)

type EmailWidget struct {
	from     string
	password string
	smtpHost string
	smtpPort string
	auth     smtp.Auth
}

func NewEmailWidget(from, password, smtpHost, smtpPort string) *EmailWidget {
	auth := smtp.PlainAuth("", from, password, smtpHost)
	return &EmailWidget{
		from:     from,
		password: password,
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		auth:     auth,
	}
}

func (ew *EmailWidget) Send(msg, destinationEmail string) error {
	to := []string{
		destinationEmail,
	}
	return smtp.SendMail(ew.smtpHost+":"+ew.smtpPort, ew.auth, ew.from, to, []byte(msg))
}
