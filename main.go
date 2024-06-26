package main

import {
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/go-resty/resty/v2"
}

const (
	apiKey = "MY KEY" // how do I obfuscate this?
	endpoint = "https://api.fortnitetracker.com/v1/profile/{platform}/{epic-name}" // get correct endpoint
)

func main(){
	client := resty.New()

	// 
	platform := "pc"
	epicName := "pegasusfly_"

	url := endpoint
	url = strings.Replace(url, "{platform}", platform, 1)
	url = strings.Replace(url, "{epic-name}", epicName, 1)

	resp, err := client.R().
		SetHeader("TRN-Api-Key", apiKey).
		Get(url)

	if err != nil {
		log.Fatal("Error fetching data from Fortnite API: %v", err)
	}

	// Writes the raw (wink wink) JSON response to the notepad
	data := resp.String()

	err = ioutil.WriteFile("fortnite_data.txt", []byte(data), 0644)
	if err != nil {
		log.Fatal("Error writing file: %v", err)
	}

	fmt.Println("Data successfully written to fortnite_data.txt")

}