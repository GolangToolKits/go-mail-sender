package mailsender

import (
	
	"testing"
)

func TestMockPlainSender_New(t *testing.T) {
	type fields struct {
		User        string
		Password    string
		MailHost    string
		Port        string
		MockSuccess bool
	}
	tests := []struct {
		name   string
		fields fields
		want   Sender
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPlainSender{
				User:        tt.fields.User,
				Password:    tt.fields.Password,
				MailHost:    tt.fields.MailHost,
				Port:        tt.fields.Port,
				MockSuccess: tt.fields.MockSuccess,
			}
			if got := m.New(); got == nil {
				t.Errorf("MockPlainSender.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPlainSender_SendMail(t *testing.T) {
	type fields struct {
		User        string
		Password    string
		MailHost    string
		Port        string
		MockSuccess bool
	}
	type args struct {
		mailer *Mailer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				User:        "test",
				Password:    "test",
				MailHost:    "smtp.test.com",
				Port:        "567",
				MockSuccess: true,
			},
			args: args{
				mailer: &Mailer{
					SenderAddress: "test@test.com",
					Recipients:    []string{"test2@test.com"},
					Subject:       "Test Mail",
					Body:          "This is only a test.",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPlainSender{
				User:        tt.fields.User,
				Password:    tt.fields.Password,
				MailHost:    tt.fields.MailHost,
				Port:        tt.fields.Port,
				MockSuccess: tt.fields.MockSuccess,
			}
			if got := m.SendMail(tt.args.mailer); got != tt.want {
				t.Errorf("MockPlainSender.SendMail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockSecureSender_SendMail(t *testing.T) {
	type fields struct {
		User        string
		Password    string
		MailHost    string
		Port        string
		MockSuccess bool
	}
	type args struct {
		mailer *Mailer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				User:        "test",
				Password:    "test",
				MailHost:    "smtp.test.com",
				Port:        "567",
				MockSuccess: true,
			},
			args: args{
				mailer: &Mailer{
					SenderAddress: "test@test.com",
					Recipients:    []string{"test2@test.com"},
					Subject:       "Test Mail",
					Body:          "This is only a test.",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSecureSender{
				User:        tt.fields.User,
				Password:    tt.fields.Password,
				MailHost:    tt.fields.MailHost,
				Port:        tt.fields.Port,
				MockSuccess: tt.fields.MockSuccess,
			}
			if got := m.SendMail(tt.args.mailer); got != tt.want {
				t.Errorf("MockSecureSender.SendMail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockSecureSender_GetNew(t *testing.T) {
	type fields struct {
		User        string
		Password    string
		MailHost    string
		Port        string
		MockSuccess bool
	}
	tests := []struct {
		name   string
		fields fields
		want   Sender
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSecureSender{
				User:        tt.fields.User,
				Password:    tt.fields.Password,
				MailHost:    tt.fields.MailHost,
				Port:        tt.fields.Port,
				MockSuccess: tt.fields.MockSuccess,
			}
			if got := m.New(); got == nil {
				t.Errorf("MockSecureSender.GetNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
