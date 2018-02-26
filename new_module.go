package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

//UnmarshallJSON ...
func UnmarshallJSON(jsonPath string) (data OperatingSystem) {

	//open file
	buildFile, readErr := ioutil.ReadFile(jsonPath)
	if readErr != nil {
		panic(readErr)
	}

	//unmarshal file
	Ops := OperatingSystem{}
	json.Unmarshal([]byte(buildFile), &Ops)
	return Ops
}
