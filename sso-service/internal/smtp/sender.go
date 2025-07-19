package sender

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Sender struct {
	Addr   string
	Auth   smtp.Auth
	Sender string
}

func New(addr string, host string, sender string, username string, password string) *Sender {

	auth := smtp.PlainAuth("", username, password, host)

	return &Sender{addr, auth, sender}
}

func (s *Sender) Send(to string, access string) error {
	const op = "smpt.sender.Send"

	subject := "Registration"
	body := fmt.Sprintf("<p>Для подтверждения регистрации перейдите по ссылке: <b>%s</b>!</p>", access)
	msg := BuildMessage(s.Sender, []string{to}, subject, body)

	err := smtp.SendMail(s.Addr, s.Auth, s.Sender, []string{to}, []byte(msg))

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func BuildMessage(from string, to []string, subject string, body string) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", from)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(to, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", subject)
	msg += fmt.Sprintf("\r\n%s\r\n", body)

	return msg
}
