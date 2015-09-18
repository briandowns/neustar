package neustar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// AlertURI is the endpoint for calls to the scripting API
	AlertURI = "alert/1.0"

	// PolicyURI is the endpoint for policy calls
	PolicyURI = "/policy"
)

// Strikes is a slice of valid strikes
var Strikes = []int{1, 2, 3}

// NewAlertPolicyParameters holds the parameters needed to pass to the NewAlertPolicy method
type NewAlertPolicyParameters struct {
	// Name of the alert policy
	Name string

	// EmailAddresses is a string slice of comma-separated email addresses associated with
	// the alert policy (e.g. ["alert@mycompany.com","myemail@gmail.com"].
	EmailAddresses []string

	// Strikes before triggering an alert.
	Strikes int

	// Description for the alert policy
	Description string
}

// NewAlertPolicyResponse
type NewAlertPolicyResponse struct {
	// Name of the alert policy
	Name string

	// EmailAddresses associated with the alert policy
	EmailAddresses []string

	// Description of the alert policy
	Description string

	// Strikes for the alert policy
	Strikes int
}

// ListAlertPoliciesResponse holds the response from the ListAlertPolicies call
type ListAlertPoliciesResponse struct {
	// ID of the alert policy
	ID string

	// Name of the alert policy
	Name string

	// EmailAddresses is a slice of strings containing email addresses associated
	// with the alert policy
	EmailAddresses []string

	// The description of the alert policy
	Description string

	// Strikes before triggering an alert.
	Strikes int

	// AdvancedEdit holds a flag that will return true if the policy is an Advanced
	// Alert Policy
	AdvancedEdit bool
}

// Alerting holds alerting config
type Alerting struct {
	neustar *Neustar
}

// NewAlertPolicy creates a new Alert policy
func (a *Alerting) NewAlertPolicy(napp *NewAlertPolicyParameters) (NewAlertPolicyResponse, error) {
	buffer, err := json.Marshal(napp)
	if err != nil {
		return NewAlertPolicyResponse{}, err
	}
	body := bytes.NewBuffer(buffer)
	var data NewAlertPolicyResponse
	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s/%s?apikey=%s&sig=%s", BaseURL, AlertURI, PolicyURI, a.neustar.Key, a.neustar.DigitalSignature()),
		body)
	if err != nil {
		return NewAlertPolicyResponse{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return NewAlertPolicyResponse{}, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return NewAlertPolicyResponse{}, err
	}
	return data, nil
}

// ListAlertPolicies retrieves a list of policies ordered by date in descending order.
func (a *Alerting) ListAlertPolicies() (ListAlertPoliciesResponse, error) {
	var response *http.Response
	var data ListAlertPoliciesResponse
	response, err := http.Get(fmt.Sprintf("%s%s/%s?apikey=%s&sig=%s", BaseURL, AlertURI, PolicyURI, a.neustar.Key, a.neustar.DigitalSignature()))
	if err != nil {
		return ListAlertPoliciesResponse{}, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return ListAlertPoliciesResponse{}, err
	}
	return data, nil
}

// ValidStrikes makes sure that the given strike is valid
func ValidStrikes(strike int) bool {
	for _, i := range Strikes {
		if i == strike {
			return true
		}
	}
	return false
}
