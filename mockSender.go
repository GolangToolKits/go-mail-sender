package mailsender

// MockPlainSender MockPlainSender
type MockPlainSender struct {
	User        string
	Password    string
	MailHost    string
	Port        string
	MockSuccess bool
}

// SendMail SendMail
func (m *MockPlainSender) SendMail(mailer *Mailer) bool {
	return m.MockSuccess
}

// New GetNew
func (m *MockPlainSender) New() Sender {
	return m
}

// MockSecureSender SecureSender
type MockSecureSender struct {
	User        string
	Password    string
	MailHost    string
	Port        string
	MockSuccess bool
}

// SendMail SendMail
func (m *MockSecureSender) SendMail(mailer *Mailer) bool {
	return m.MockSuccess
}

// New GetNew
func (m *MockSecureSender) New() Sender {
	return m
}
