package mailsender

import (
	"log"
	"net/smtp"
)

// PlainSender PlainSender
type PlainSender struct {
	User     string
	Password string
	MailHost string
	Port     string
}

// SendMail SendMail
func (m *PlainSender) SendMail(mailer *Mailer) bool {
	var rtn bool
	if mailer.Recipients != nil && len(mailer.Recipients) > 0 {
		auth := smtp.PlainAuth("", m.User, m.Password, m.MailHost)
		var recips = "To: "
		for r := range mailer.Recipients {
			recips += mailer.Recipients[r]
			if r < len(mailer.Recipients)-1 {
				recips += ";"
			}
		}
		recips += "\r\n"
		var msg = []byte(
			"From: " + mailer.SenderAddress + "\r\n" +
				recips +
				"Subject: " + mailer.Subject + "\r\n" +
				"\r\n" +
				mailer.Body + "\r\n")
		//fmt.Println("mailer: ", *mailer)
		err := smtp.SendMail(m.MailHost+":"+m.Port, auth, mailer.SenderAddress, mailer.Recipients, msg)
		if err != nil {
			log.Println("Mail send error: ", err)
		} else {
			rtn = true
		}
	}
	return rtn
}

// New New
func (m *PlainSender) New() Sender {
	return m
}
