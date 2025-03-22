package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func request[T any](dto RequestDTO[T]) (*http.Response, error) { // todo ta funkcja ma tylko jednego callera

	// Convert the struct to JSON
	requestBody, err := json.Marshal(dto)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %w", err)
	}

	// Define the URL to send the POST request to
	url := serverAddress + "/api/"

	// Send the POST request with the JSON request body
	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error sending POST request: %w", err)
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code: " + response.Status)
	}

	return response, nil
}
