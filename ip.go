//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goipify.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package goipify

import (
	"encoding/json"
	"net/http"
)

const (
	transferProtocol = "https://"
	baseUrl          = "api.ipify.org"
)

// Return is to decode the json data
type Return struct {
	Ip string `json:"ip"`
}

// Ip is to get your ip address
func Ip() (Return, error) {

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest("GET", transferProtocol+baseUrl, nil)
	if err != nil {
		return Return{}, err
	}

	// Parse url & add attributes
	parse := request.URL.Query()
	parse.Add("format", "json")
	request.URL.RawQuery = parse.Encode()

	// Define header & check if access token is active
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return Return{}, err
	}

	// Decode data
	var decode Return

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return Return{}, err
	}

	// Return data
	return decode, err

}
