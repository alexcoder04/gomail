package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	settingsFile   = flag.String("s", "settings.yml", "settings file")
	recipientsFile = flag.String("r", "recipients.txt", "recipients file")
	mailBodyFile   = flag.String("b", "mail.txt", "mail body file")
)

type Settings struct {
	Addr     string `yaml:"Addr"`
	From     string `yaml:"From"`
	Host     string `yaml:"Host"`
	Password string `yaml:"Password"`
	Subject  string `yaml:"Subject"`
	Username string `yaml:"Username"`
}

func ReadSettingsFromFile(file string) Settings {
	fmt.Println("Reading settings from " + file + "...")
	data, err := ioutil.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Settings file (" + file + ") does not exist")
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

func ReadRecipientsAddressesFromFile(file string) []string {
	fmt.Println("Reading recipients from " + file + "...")
	readFile, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Recipients file (" + file + ") does not exist")
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

	fmt.Println("Reading mail content from " + *mailBodyFile + "...")
	mailContent, err := ioutil.ReadFile(*mailBodyFile)
	if err != nil {
		fmt.Println("Error reading " + *mailBodyFile)
		os.Exit(1)
	}

	fmt.Println("Performing login on " + settings.Host + "...")
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
