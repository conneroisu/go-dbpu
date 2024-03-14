package dbpu

import (
	"fmt"
	"net/http"
)

// ServerClient is a struct that contains the server and client locations.
type ServerClient struct {
	Server string `json:"server"`
	Client string `json:"client"`
}

// ClosestLocation returns the closest location to the given latitude and longitude.
func ClosestLocation() (ServerClient, error) {
	req, reqErr := newClosestLocationRequest()
	done, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[ServerClient](done)
	defer done.Body.Close()
	return resolveApiCall(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// newClosestLocationRequest creates a request for ClosestLocation.
func newClosestLocationRequest() (*http.Request, error) {
	url := fmt.Sprintf("https://region.turso.io/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request. %v", err)
	}
	return req, nil
}
