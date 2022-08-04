package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/smtp"
)

const PROGRAM_NAME = "gomail"

var (
	settingsFile   = flag.String("s", GetSettingsFile(), "settings file")
	recipientsFile = flag.String("r", GetRecepientsFile(), "recipients file")
	mailBodyFile   = flag.String("b", "./mail.txt", "mail body file")
)

type Settings struct {
	Addr     string `yaml:"Addr"`
	From     string `yaml:"From"`
	Host     string `yaml:"Host"`
	Password string `yaml:"Password"`
	Subject  string `yaml:"Subject"`
	Username string `yaml:"Username"`
}

func BuildMessage(from string, to string, subject string, body []byte) []byte {
	msg := append([]byte("MIME-version: 1.0;\r\n"+
		"Content-Type: text/plain; charset=utf-8;\r\n"+
		"From: "+from+"\r\n"+
		"To: "+to+"\r\n"+
		"Subject: "+subject+"\r\n\r\n"), body...)
	return msg
}

func main() {
	flag.Parse()

	settings := ReadSettingsFromFile(*settingsFile)
	recipients := ReadRecipientsAddressesFromFile(*recipientsFile)

	log.Printf("Reading mail content from %s...\n", *mailBodyFile)
	mailContent, err := ioutil.ReadFile(*mailBodyFile)
	if err != nil {
		log.Fatalf("Error reading %s\n", *mailBodyFile)
	}

	log.Printf("Performing login on %s...\n", settings.Host)
	auth := smtp.PlainAuth("", settings.Username, settings.Password, settings.Host)
	for _, to := range recipients {
		msg := BuildMessage(settings.From, to, settings.Subject, mailContent)
		err := smtp.SendMail(settings.Addr, auth, settings.From, []string{to}, msg)
		if err != nil {
			log.Fatalf("Cannot send mail: %s\n", err.Error())
		} else {
			log.Printf("Email sent to %s\n", to)
		}
	}
}
