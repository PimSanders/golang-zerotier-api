package golangzerotierapi

import (
	"bytes"
	"encoding/json"
	"log"
)

// GetNetworkList retrieves a list of networks.
// It returns a slice of Network structs and an error if any.
func (c *Client) GetNetworkList() ([]Network, error) {
	var respMap []Network

	// Make a GET request to retrieve network list
	respBody, err := c.MakeGetRequest("/network")
	if err != nil {
		return respMap, err
	}

	// Unmarshal the JSON response into the respMap slice
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		log.Panic(err)
	}

	return respMap, err
}

// GetNetwork retrieves details of a specific network identified by its ID.
// It returns a Network struct and an error if any.
func (c *Client) GetNetwork(networkID string) (Network, error) {
	var respMap Network

	// Make a GET request to retrieve details of the specified network
	respBody, err := c.MakeGetRequest("/network/" + networkID)
	if err != nil {
		return respMap, err
	}

	// Unmarshal the JSON response into the respMap struct
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		log.Panic(err)
	}

	return respMap, err
}

// GetNetworkMemberList retrieves a list of members belonging to a specific network identified by its ID.
// It returns a slice of NetworkMember structs and an error if any.
func (c *Client) GetNetworkMemberList(networkID string) ([]NetworkMember, error) {
	var respMap []NetworkMember

	// Make a GET request to retrieve the member list of the specified network
	respBody, err := c.MakeGetRequest("/network/" + networkID + "/member")
	if err != nil {
		return respMap, err
	}

	// Unmarshal the JSON response into the respMap slice
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		log.Panic(err)
	}

	return respMap, err
}

// GetNetworkMember retrieves details of a specific network member identified by its ID and the network's ID.
// It returns a NetworkMember struct and an error if any.
func (c *Client) GetNetworkMember(networkID string, memberID string) (NetworkMember, error) {
	var respMap NetworkMember

	// Make a GET request to retrieve details of the specified network member
	respBody, err := c.MakeGetRequest("/network/" + networkID + "/member/" + memberID)
	if err != nil {
		return respMap, err
	}

	// Unmarshal the JSON response into the respMap struct
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		log.Panic(err)
	}

	return respMap, err
}

// CreateNetwork creates a new network.
// It returns the created Network struct and an error if any.
func (c *Client) CreateNetwork() (Network, error) {
	var respMap Network

	// Create an empty JSON object as data
	data := []byte("{}")
	reader := bytes.NewReader(data)

	// Make a POST request to create a new network
	respBody, err := c.MakePostRequest("/network", reader)

	if err != nil {
		log.Panic(err)
		return respMap, err
	}

	// Unmarshal the JSON response into the respMap struct
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		log.Panic(err)
		return respMap, err
	}

	return respMap, err
}

// UpdateNetwork updates an existing network identified by its ID with the provided data.
// It returns the updated Network struct and an error if any.
func (c *Client) UpdateNetwork(networkID string, network UpdateNetwork) (Network, error) {
	var respMap Network

	// Marshal the provided network data into JSON
	data, err := json.Marshal(network)
	if err != nil {
		log.Panic(err)
		return respMap, err
	}
	reader := bytes.NewReader(data)

	// Make a POST request to update the network
	respBody, err := c.MakePostRequest("/network/"+networkID, reader)

	if err != nil {
		log.Panic(err)
		return respMap, err
	}

	// Unmarshal the JSON response into the respMap struct
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		log.Panic(err)
		return respMap, err
	}

	return respMap, err
}

// DeleteNetwork deletes a network identified by its ID.
// It returns an error if any.
func (c *Client) DeleteNetwork(networkID string) error {
	// Make a DELETE request to delete the network
	err := c.MakeDeleteRequest("/network/" + networkID)
	if err != nil {
		log.Panic(err)
	}

	return err
}

// UpdateNetworkMember updates an existing member in a network identified by their IDs.
// It takes an UpdateNetworkMember struct containing the updated member information.
// It returns a NetworkMember struct and an error if any.
// STILL BROKEN!
func (c *Client) UpdateNetworkMember(networkID string, memberID string, member UpdateNetworkMember) (NetworkMember, error) {
	var respMap NetworkMember

	// Convert the UpdateNetworkMember struct to JSON data
	data, err := json.Marshal(member)
	if err != nil {
		log.Panic(err)
		return respMap, err
	}

	reader := bytes.NewReader(data)

	// Make a POST request to update the network member
	respBody, err := c.MakePostRequest("/network/"+networkID+"/member/"+memberID, reader)
	if err != nil {
		return respMap, err
	}

	// Unmarshal the JSON response into the respMap struct
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		log.Panic(err)
	}

	return respMap, err
}

// DeleteNetworkMember deletes a member from a network identified by their IDs.
// It returns an error if any.
func (c *Client) DeleteNetworkMember(networkID string, memberID string) error {
	// Make a DELETE request to delete the network member
	err := c.MakeDeleteRequest("/network/" + networkID + "/member/" + memberID)
	if err != nil {
		log.Panic(err)
	}

	return err
}
