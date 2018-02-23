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

	// Main

	url := "https://raw.githubusercontent.com/3sky/step-template-go/master/misc/build.json"
	tempPath := "tmp"
	os.MkdirAll("tmp", os.ModePerm)
	DownloadFile("tmp/build.json", url)
	data := UnmarshallJson("tmp/build.json")
	fmt.Printf("%+v\n", data)
	err := os.RemoveAll(tempPath)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Hello world!")

}
