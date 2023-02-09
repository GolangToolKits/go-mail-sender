package mailsender

// Mailer Mailer
type Mailer struct {
	SenderAddress string
	Recipients    []string
	Subject       string
	Body          string
}

// Sender Sender
type Sender interface {
	SendMail(mailer *Mailer) bool
}

//go mod init github.com/GolangToolKits/go-mail-sender
