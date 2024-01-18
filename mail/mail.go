// Package mail @Author Zhan 2024/1/18 10:00:00
package mail

import (
	"encoding/base64"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"strings"
)

type Client struct {
	addr   string
	sender string
	auth   sasl.Client
}

func (c *Client) Send(receivers []string, subject string, body string) error {
	subjectBase := base64.StdEncoding.EncodeToString([]byte(subject))

	msg := strings.NewReader(
		"From: " + c.sender + "\r\n" +
			"To: " + strings.Join(receivers, ",") + "\r\n" +
			"Subject: =?UTF-8?B?" + subjectBase + "?=\r\n" +
			"Content-Type: text/html; charset=UTF-8" +
			"\r\n\r\n" +
			body + "\r\n")

	err := smtp.SendMail(c.addr, c.auth, c.sender, receivers, msg)
	if err != nil {
		return err
	}
	return nil
}

func NewMailClient(addr string, sender string, username string, password string) *Client {
	auth := sasl.NewPlainClient("", username, password)
	return &Client{
		addr:   addr,
		sender: sender,
		auth:   auth,
	}
}
