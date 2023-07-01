package mail

import (
	"context"
	"net/smtp"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
)

type (
	SendMailRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Message  string `json:"message"`
		Subject  string `json:"subject"`
	}

	SendMailResponse struct {
		Status string `json:"status"`
		SendAt string `json:"send_at"`
	}
)

var (
	Host     = "smtp.mailtrap.io"
	Port     = 2525
	User     = "53c73caf7553e1"
	Password = "ccbeaed9132daf"
)

func SendMail(ctx context.Context, req *SendMailRequest) (*SendMailResponse, error) {
	smtpAuth := smtp.PlainAuth("", User, Password, Host)
	em := email.NewEmail()
	em.From = "ElloApp@merehead.com"
	em.To = []string{req.Email}
	em.Subject = req.Subject
	em.Text = []byte(req.Message)
	err := em.Send(Host+":"+strconv.Itoa(Port), smtpAuth)
	var status string
	if err != nil {
		status = err.Error()
	} else {
		status = "OK"
	}

	return &SendMailResponse{
		Status: status,
		SendAt: time.Now().Format("2006-01-02 15:04:05"),
	}, err
}
