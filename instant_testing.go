package neustar

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// ToolsURI is the endpoint for instant test queries
	ToolsURI = "tools/instanttest/1.0"
)

// InstantTesting
type InstantTesting struct {
	neustar *Neustar
}

// NewInstantTest returns a new InstantTesting object
func NewInstantTest(neustar *Neustar) *InstantTesting {
	return &InstantTesting{
		neustar: neustar,
	}
}

// InstantTestingJobReponse contains the response from the API
type InstantTestingData struct {
	Status       string `json:""`
	Location     string `json:"location"`
	ResponseTime int    `json:"responseTime"`
	ID           string `json:"id"`
}

// InstanceTestingResponse
type InstanceTestingResponse struct {
	Data struct {
		// A list of instant test job locations, their status and the specific location
		Items []InstantTestingData `json:"items"`

		// The ID of the instant test job
		ID string `json:"id"`
	} `json:"data"`
}

// InstantTestingByLocationResponse
type InstantTestingByLocationResponse struct {
	// The ID of the instant test job
	ID string `json:"id"`

	// The current status of the instant test job
	Status string `json:"status"`

	// Available on completion, the resultant har file for this instant
	// test job.
	HARFile string `json:"harFile"`

	// Available on completion, the base64 encoded screenshot for this
	// instant test job.
	Screenshot string `json:"screenshot"`
}

// InstantTestingCreateResponse holds the response from the API on Instant
// Test creation
type InstantTestingCreateResponse struct {
	Data struct {
		Items struct {
			ID       string `json:"id"`
			Loctions []struct {
				ID       string `json:"id"`
				Location string `json:"location"`
			} `json:"locations"`
			Created string `json:"created"`
		} `json:"items"`
	} `json:"data"`
}

// Create creates a new instant test job and return the job id of
// the new instant test job. Url is required. You may optionally
// supply a callback URL. For every stage of the instant test process,
//we will POST the current status of your instant test job
func (i *InstantTesting) Create(url, callback string) (InstantTestingCreateResponse, error) {
	var data InstantTestingCreateResponse
	return data, nil
}

// GetJob retrieves information for a specific instant test job, along
// with information from each location being tested.
func (i *InstantTesting) GetJob(instantTestID string) (InstanceTestingResponse, error) {
	var response *http.Response
	var data InstanceTestingResponse
	response, err := http.Get(
		fmt.Sprintf("%s%s/%s?apikey=%s&sig=%s", BaseURL, ToolsURI, instantTestID, i.neustar.Key, i.neustar.DigitalSignature()))
	if err != nil {
		return InstanceTestingResponse{}, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return InstanceTestingResponse{}, err
	}
	return data, nil
}

// GetJobByLocations retrieves information for a specific instant test job by location.
func (i *InstantTesting) GetJobByLocations(instantTestID, instantTestLocationID string) {}
