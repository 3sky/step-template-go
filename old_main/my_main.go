package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bitrise-io/go-utils/log"
)

//OperatingSystem ...
type OperatingSystem struct {
	Android AndoridData `json:"android"`
	IOS     iOSData     `json:"ios"`
}

//AndoridData ...
type AndoridData struct {
	Release Release `json:"release"`
}

//Release ...
type Release struct {
	Keystore      string `json:"keystore"`
	StorePassword string `json:"storePassword"`
	Alias         string `json:"alias"`
	Pasword       string `json:"password"`
	IsPresent     bool   `json:"isPresent"`
}

//iOSData ...
type iOSData struct {
	Debug Debug `json:"debug"`
}

//Debug ...
type Debug struct {
	UID              int    `json:"UID"`
	CodeSignIdentity string `json:"codeSignIdentity"`
	DevelopmentTeam  string `json:"developmentTeam"`
	PackageType      string `json:"packageType"`
}

//DownloadFile ...
func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func myMain() {

	tempPath := "tmp"

	url := flag.String("url", "None", "URL to download json, with https://")

	flag.Parse()
	/*
	   https://github.com/3sky/step-template-go/blob/master/misc/build.json
	*/
	//fmt.Println(*url)

	os.MkdirAll(tempPath, os.ModePerm)
	DownloadFile("tmp/build.json", *url)

	buildFile, readErr := ioutil.ReadFile("tmp/build.json")
	if readErr != nil {
		panic(readErr)
	}

	Ops := OperatingSystem{}
	json.Unmarshal([]byte(buildFile), &Ops)
	//You can take whaever you want
	log.Printf("%+v\n", Ops.Android.Release.Pasword)
	log.Printf("%+v\n", Ops.IOS.Debug.UID)
	log.Printf("%+v\n", Ops)

	err := os.RemoveAll(tempPath)
	if err != nil {
		fmt.Println(err)
	}
}
