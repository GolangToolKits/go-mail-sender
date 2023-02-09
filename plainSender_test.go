package mailsender

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type mailFile struct {
	MailHost   string   `json:"mailHost"`
	Port       string   `json:"port"`
	User       string   `json:"user"`
	Password   string   `json:"password"`
	Sender     string   `json:"sender"`
	Recipients []string `json:"recipients"`
}

func TestPlainSender_New(t *testing.T) {
	type fields struct {
		User     string
		Password string
		MailHost string
		Port     string
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
			m := &PlainSender{
				User:     tt.fields.User,
				Password: tt.fields.Password,
				MailHost: tt.fields.MailHost,
				Port:     tt.fields.Port,
			}
			if got := m.New(); got == nil {
				t.Errorf("PlainSender.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlainSender_SendMail(t *testing.T) {
	//ports 587 465 80
	var fileName = "../mailGmail.json"
	var mm mailFile
	file, err2 := ioutil.ReadFile(fileName)
	if err2 == nil {
		err := json.Unmarshal(file, &mm)
		fmt.Println("marshal err: ", err)
	}
	fmt.Println("file: ", mm)

	type fields struct {
		User     string
		Password string
		MailHost string
		Port     string
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
				User:     mm.User,
				Password: mm.Password,
				MailHost: mm.MailHost,
				Port:     mm.Port,
			},
			args: args{
				mailer: &Mailer{
					SenderAddress: mm.Sender,
					Recipients:    mm.Recipients,
					Subject:       "Test Mail",
					Body:          "This is only a test.",
				},
			},
			//----fix---want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &PlainSender{
				User:     tt.fields.User,
				Password: tt.fields.Password,
				MailHost: tt.fields.MailHost,
				Port:     tt.fields.Port,
			}
			m := ms.New()
			if got := m.SendMail(tt.args.mailer); got != tt.want {
				t.Errorf("PlainSender.SendMail() = %v, want %v", got, tt.want)
			}
		})
	}
}
