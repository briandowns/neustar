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

// InstantTestingJobReponse
type InstantTestingJobReponse struct {
	ID string `json:"id"`
}

// Create creates a new instant test job and return the job id of
// the new instant test job. Url is required. You may optionally
// supply a callback URL. For every stage of the instant test process,
//we will POST the current status of your instant test job
func (i *InstantTesting) Create() {}

// GetJob retrieves information for a specific instant test job, along
// with information from each location being tested.
func (i *InstantTesting) GetJob(instantTestID string) (interface{}, error) {
	var response *http.Response
	var data map[string]map[string][]Monitor
	response, err := http.Get(fmt.Sprintf("%s%s/%s?apikey=%s&sig=%s", BaseURL, ToolsURI, instantTestID, i.neustar.Key, i.neustar.DigitalSignature()))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data["data"]["items"], nil
}

// GetJobByLocations retrieves information for a specific instant test job by location.
func (i *InstantTesting) GetJobByLocations(instantTestID, instantTestLocationID string) {}
