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
	SETTINGS_FILE   = "./settings.yml"
	MAIL_CONTENT    = "./mail.txt"
	ACCOUNT_FILE    = "./account.txt"
	RECIPIENTS_FILE = "./recipients.txt"
)

type Settings struct {
	From    string `yaml:"From"`
	Addr    string `yaml:"Addr"`
	Host    string `yaml:"Host"`
	Subject string `yaml:"Subject"`
}

type Account struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}

func ReadSettingsFromFile() (string, string, string, string) {
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

	return settings.From, settings.Addr, settings.Host, settings.Subject
}

func ReadAccountFromFile() (string, string) {
	data, err := ioutil.ReadFile(ACCOUNT_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Account file (" + ACCOUNT_FILE + ") does not exist")
		} else {
			fmt.Println("Cannot open settings file")
		}
		os.Exit(1)
	}
	account := Account{}
	err = yaml.Unmarshal(data, &account)
	if err != nil {
		fmt.Println("Your settings file doesn't look right")
		os.Exit(1)
	}

	return account.Username, account.Password
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
	from, addr, host, subject := ReadSettingsFromFile()
	user, password := ReadAccountFromFile()
	recipients := ReadRecipientsAddressesFromFile()

	mailContent, _ := ioutil.ReadFile(MAIL_CONTENT)

	auth := smtp.PlainAuth("", user, password, host)
	for _, to := range recipients {
		msg := BuildMessage(from, to, subject, mailContent)
		err := smtp.SendMail(addr, auth, from, []string{to}, msg)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("Email sent to", to)
		}
	}
}
