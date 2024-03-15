package golangzerotierapi

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)

// makeGetRequest performs a GET request to the specified URL.
// It returns the response body ([]byte) and an error if any.
func (c *Client) MakeGetRequest(url string) ([]byte, error) {
	var respBody []byte

	// Create an HTTP client
	client := &http.Client{}

	// Create a new GET request with the provided URL
	req, _ := http.NewRequest("GET", c.BaseURL+url, nil)

	// Set the Authorization header with the provided c.Token
	req.Header.Set("Authorization", "token "+c.Token)

	// Perform the HTTP request
	resp, err := client.Do(req)

	// Check if there's an error during the request
	if err != nil {
		// Return empty response body and the error
		return respBody, err
	}

	// Check if the response status code is not 200 OK
	if resp.StatusCode != 200 {
		// Create an error indicating non-200 status code
		err := errors.New("Non 200 response: " + strconv.Itoa(resp.StatusCode))
		// Return empty response body and the error
		return respBody, err
	}

	// Read the response body
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		// Return empty response body and the error if reading fails
		return respBody, err
	}

	// Return the response body and nil error
	return respBody, err
}

// makePostRequest performs a POST request to the specified URL with the provided data.
// It returns the response body ([]byte) and an error if any.
func (c *Client) MakePostRequest(url string, data io.Reader) ([]byte, error) {
	var respBody []byte

	// Create an HTTP client
	client := &http.Client{}

	// Create a new POST request with the provided URL and data
	req, _ := http.NewRequest("POST", c.BaseURL+url, data)

	// Set the Authorization header with the provided c.Token
	req.Header.Set("Authorization", "token "+c.Token)

	// Set the Content-Type header to indicate JSON data
	req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP request
	resp, err := client.Do(req)

	// Check if there's an error during the request
	if err != nil {
		// Return empty response body and the error
		return respBody, err
	}

	// Check if the response status code is not 200 OK
	if resp.StatusCode != 200 {

		// Create an error indicating non-200 status code
		err := errors.New("Non 200 response: " + strconv.Itoa(resp.StatusCode))
		// Return empty response body and the error
		return respBody, err
	}

	// Read the response body
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		// Return empty response body and the error if reading fails
		return respBody, err
	}

	// Return the response body and nil error
	return respBody, err
}

// makeDeleteRequest performs a DELETE request to the specified URL.
// It returns an error if any.
func (c *Client) MakeDeleteRequest(url string) error {
	// Create an HTTP client
	client := &http.Client{}

	// Create a new DELETE request with the provided URL
	req, _ := http.NewRequest("DELETE", c.BaseURL+url, nil)

	// Set the Authorization header with the provided c.Token
	req.Header.Set("Authorization", "token "+c.Token)

	// Perform the HTTP request
	resp, err := client.Do(req)

	// Check if there's an error during the request
	if err != nil {
		// Return the error
		return err
	}

	// Check if the response status code is not 200 OK
	if resp.StatusCode != 200 {
		// Create an error indicating non-200 status code
		err := errors.New("Non 200 response: " + strconv.Itoa(resp.StatusCode))
		// Return the error
		return err
	}

	// Return nil error if everything is successful
	return err
}
