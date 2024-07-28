package mail

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/smtp"

	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
)

var auth smtp.Auth

func SendMail(subject string, templateFile string, to []string, data interface{}) error {
	auth = smtp.PlainAuth(
		"",
		helpers.Env("MAIL_USER"),
		helpers.Env("MAIL_PASS"),
		helpers.Env("MAIL_HOST"),
	)

	tParsed, err := parseTemplate(templateFile, data)

	if err != nil {
		return err
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectMsg := "Subject: " + subject + "\n"
	addr := helpers.Env("MAIL_HOST") + ":" + helpers.Env("MAIL_PORT")
	mailsTo, toErr := json.Marshal(to)
	if toErr != nil {
		panic(toErr)
	}
	mailTo := "To: " + string(mailsTo) + "\r\n"
	mailFrom := "From: " + helpers.Env("MAIL_FROM") + "\r\n"
	msg := []byte(mailFrom + mailTo + subjectMsg + mime + "\n" + tParsed.String())

	if err := smtp.SendMail(addr, auth, "from@app.com", to, msg); err != nil {
		panic(err)
	}

	return nil
}

func parseTemplate(templateFile string, data interface{}) (*bytes.Buffer, error) {
	t := template.New("template")
	var err error
	t, err = template.ParseFiles(templateFile)
	if err != nil {
		return nil, err
	}

	var tpl bytes.Buffer
	if err = t.Execute(&tpl, data); err != nil {
		return nil, err
	}

	return &tpl, nil
}
