package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func getConfigDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalln("Cannot determine user's home directory")
	}
	return dir
}

func GetSettingsFile() string {
	return path.Join(getConfigDir(), "gomail", "settings.yml")
}

func GetRecepientsFile() string {
	return path.Join(getConfigDir(), "gomail", "recipients.txt")
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
	log.Printf("Loading recipients from %s...\n", file)
	readFile, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Recipients file (%s) does not exist", file)
		}
		log.Fatalln("Cannot open settings file")
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
