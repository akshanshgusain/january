package mailer

import (
	"errors"
	"testing"
)

func TestMail_SendSMTPMessage(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Joe",
		To:          "you@there.com",
		Subject:     "test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	err := mailer.SendSMTPMessage(msg)
	if err != nil {
		t.Error(err)
	}
}

func TestMail_SendUsingChan(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Joe",
		To:          "you@there.com",
		Subject:     "test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	mailer.Jobs <- msg
	res := <-mailer.Results
	if res.Error != nil {
		t.Error(errors.New("failed to send over channel"))
	}

	msg.To = "not_an_email_address"
	mailer.Jobs <- msg
	res = <-mailer.Results
	if res.Error == nil {
		t.Error(errors.New("no error received with invalid to address"))
	}
}
