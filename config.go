package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/alexcoder04/friendly"
	"gopkg.in/yaml.v3"
)

func GetSettingsFile() string {
	confDir, err := friendly.GetConfigDir(PROGRAM_NAME)
	if err != nil {
		log.Fatalln("Cannot get config directory")
	}
	return path.Join(confDir, "settings.yml")
}

func GetRecepientsFile() string {
	confDir, err := friendly.GetConfigDir(PROGRAM_NAME)
	if err != nil {
		log.Fatalln("Cannot get config directory")
	}
	return path.Join(confDir, "recipients.txt")
}

func ReadSettingsFromFile(file string) Settings {
	log.Printf("Loading settings from %s...\n", file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Settings file (%s) does not exist\n", file)
		}
		log.Fatalln("Cannot open settings file")
	}
	settings := Settings{}
	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		log.Fatalln("Your settings file doesn't look right")
	}
	return settings
}

func ReadRecipientsAddressesFromFile(file string) []string {
	lines, err := friendly.ReadLines(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Recipients file (%s) does not exist\n", file)
		}
		log.Fatalln("Cannot read recipients")
	}
	recipients := make([]string, 0)
	for _, l := range lines {
		recipients = append(recipients, strings.TrimSpace(l))
	}
	return recipients
}
