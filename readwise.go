package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func sendHighlightsToReadwise(highlights []Highlight) error {
	apiURL := "https://readwise.io/api/v2/highlights/"
	apiKey := "ZK4lUitLvOjVRpecqkAQYjgkSGiIGjtmh6I1CmbrHYGJ9ZINpW"

	bodyData := map[string]interface{}{
		"highlights": highlights,
	}

	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	return nil
}

func addSioyekTag(highlightID int64) {
	highlightIDString := strconv.FormatInt(highlightID, 10)
	apiURL := "https://readwise.io/api/v2/highlights/" + highlightIDString + "/tags/"
	apiKey := os.Getenv("READWISE_API_KEY")

	// Define the request body 'bodyData' as a json object with the key highlights, and with the keys text and title, from the highlight struct paramater of the function

	// Encode the request body as JSON, containing a property 'name' with a value 'sioyek'
	bodyData := map[string]interface{}{
		"name": "sioyek",
	}

	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		panic(err)
	}

	// Create a new HTTP POST request with the JSON body
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Token %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	// Send the request and print the response body and status code on the console
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
