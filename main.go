package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/bitrise-io/go-utils/log"
)

// ConfigsModel ...
type ConfigsModel struct {
	ExampleInput string
}

type configuration struct {
	Enabled bool
	Path    string
}

func createConfigsModelFromEnvs() ConfigsModel {
	return ConfigsModel{
		ExampleInput: os.Getenv("example_step_input"),
	}
}

func (configs ConfigsModel) print() {
	log.Infof("Configs:")
	log.Printf("- ExampleInput: %s", configs.ExampleInput)
}

func (configs ConfigsModel) validate() error {
	if configs.ExampleInput == "" {
		return errors.New("no ExampleInput parameter specified")
	}

	return nil
}

func main() {
	// Input validation
	configs := createConfigsModelFromEnvs()

	fmt.Println()
	configs.print()

	if err := configs.validate(); err != nil {
		log.Errorf("Issue with input: %s", err)
		os.Exit(1)
	}

	fmt.Println()

	/*
		"https://raw.githubusercontent.com/3sky/step-template-go/master/misc/build.json"
	*/
	url := os.Getenv("download_url")
	tempPath := "tmp"
	os.MkdirAll("tmp", os.ModePerm)
	DownloadFile("tmp/build.json", url)
	data := UnmarshallJSON("tmp/build.json")
	log.Printf("%+v\n", data)
	err := os.RemoveAll(tempPath)
	if err != nil {
		log.Printf("Error while remove tmp directory")
	}
	log.Printf("Hello world!")

}
