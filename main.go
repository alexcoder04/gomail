package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	MAIL_CONTENT    = "./mail.txt"
	RECIPIENTS_FILE = "./recipients.txt"
	SETTINGS_FILE   = "./settings.yml"
)

type Settings struct {
	Addr     string `yaml:"Addr"`
	From     string `yaml:"From"`
	Host     string `yaml:"Host"`
	Password string `yaml:"Password"`
	Subject  string `yaml:"Subject"`
	Username string `yaml:"Username"`
}

func ReadSettingsFromFile() Settings {
	data, err := ioutil.ReadFile(SETTINGS_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Settings file (" + SETTINGS_FILE + ") does not exist")
		} else {
			fmt.Println("Cannot open settings file")
		}
		os.Exit(1)
	}
	settings := Settings{}
	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		fmt.Println("Your settings file doesn't look right")
		os.Exit(1)
	}

	return settings
}

func ReadRecipientsAddressesFromFile() []string {
	readFile, err := os.Open(RECIPIENTS_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Settings file (" + RECIPIENTS_FILE + ") does not exist")
		} else {
			fmt.Println("Cannot open settings file")
		}
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	return fileTextLines
}

func BuildMessage(from string, to string, subject string, body []byte) []byte {
	msg := []byte("MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\r\n")
	msg = append(msg, []byte("From: "+from+"\r\n")...)
	msg = append(msg, []byte("To: "+to+"\r\n")...)
	msg = append(msg, []byte("Subject: "+subject+"\r\n\r\n")...)
	msg = append(msg, body...)
	msg = append(msg, []byte("\r\n\r\n")...)
	return msg
}

func main() {
	settings := ReadSettingsFromFile()
	recipients := ReadRecipientsAddressesFromFile()

	mailContent, _ := ioutil.ReadFile(MAIL_CONTENT)

	auth := smtp.PlainAuth("", settings.Username, settings.Password, settings.Host)
	for _, to := range recipients {
		msg := BuildMessage(settings.From, to, settings.Subject, mailContent)
		err := smtp.SendMail(settings.Addr, auth, settings.From, []string{to}, msg)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("Email sent to", to)
		}
	}
}
