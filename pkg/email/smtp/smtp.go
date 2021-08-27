package smtp

import (
	"errors"
	"fmt"
	"github.com/SerjLeo/mlf_backend/pkg/email"
	goSMTP "net/smtp"
)

type SMTPSender struct {
	host   string
	pass   string
	port   string
	sender string
}

func NewSMTPSender(host string, pass string, port string, sender string) (*SMTPSender, error) {
	if !email.IsEmailValid(sender) {
		return nil, errors.New("invalid sender's email")
	}
	return &SMTPSender{host: host, pass: pass, port: port, sender: sender}, nil
}

func (s *SMTPSender) SendEmail(input email.SendInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	auth := goSMTP.PlainAuth("", s.sender, s.pass, s.host)
	addr := s.host + ":" + s.port
	subject := "Subject:" + input.Subject + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte( subject + mime + input.Body)

	if err := goSMTP.SendMail(addr, auth, s.sender, []string{input.To}, msg); err != nil {
		fmt.Println("error while sending email", err.Error())
		return err
	}

	return nil
}
