package main

import (
	"fmt"
	"os"
	//"github.com/bitrise-io/go-utils/log"
)



func main() {

	url := "https://raw.githubusercontent.com/3sky/step-template-go/master/misc/build.json"
	tempPath := "tmp"
	os.MkdirAll("tmp", os.ModePerm)
	DownloadFile("tmp\build.json", url)
	data := UnmarshallJson("build.json")
	fmt.Printf("%+v\n", data) 
	err := os.RemoveAll(tempPath)
	if err != nil {
	 	fmt.Println(err)
	}



}
