package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func askServer(userName string) (ServerState, error) {

	requestDTO := RequestDTO[interface{}]{
		Command:  "get-instructions-for-watcher",
		UserName: userName,
		Data:     nil,
	}

	requestJsonStr, err := json.Marshal(requestDTO)

	if err != nil {
		return ServerState{}, fmt.Errorf("error marshalling JSON: %w", err)
	}

	// Send the POST request with the JSON request body
	resp, err := http.Post(serverAddress+"/api/", "application/json", bytes.NewBuffer(requestJsonStr))
	if err != nil {
		return ServerState{}, fmt.Errorf("error sending POST request: %w", err)
	}

	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("--Response body reading err", err)
		return ServerState{}, err
	}

	var responseDTO ResponseDTO[ServerState]
	if err := json.Unmarshal(body, &responseDTO); err != nil {
		fmt.Println("--Response unmarshaling err", err)
		return ServerState{}, err
	}

	return responseDTO.Data, nil

}
