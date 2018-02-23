package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDownloadFile(t *testing.T) {

	testPath := "test.json"
	testUrl := "https://raw.githubusercontent.com/3sky/step-template-go/master/misc/build.json"

	DownloadFile(testPath, testUrl)

	testData, readErr := ioutil.ReadFile(testPath)
	if readErr != nil {
		t.Errorf("File doesn't exist")
	}

	if len(testData) == 0 {
		t.Errorf("File is empty")
		os.RemoveAll(testPath)
	}

	os.RemoveAll(testPath)
}

func TestUnmarshallJson(t *testing.T) {

	testData := UnmarshallJson("misc/build.json")

	if testData.Android.Release.Pasword != "bitrise" ||
		testData.IOS.Debug.UID != 2367 ||
		testData.Android.Release.IsPresent != false {
		t.Errorf("Error with unmarshall")
	}

}
