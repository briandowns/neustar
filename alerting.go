package neustar

import (
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

// NewAlertPolicyParameters
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

// ListAlertPoliciesResponse
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

// Create creates a new Alert policy
func (a *Alerting) Create() {}

// List retrieves a list of policies ordered by date in descending order.
func (a *Alerting) List() ([]ScriptingListResponse, int, error) {
	var response *http.Response
	var data ScriptDataResponse
	response, err := http.Get(fmt.Sprintf("%s%s?apikey=%s&sig=%s", BaseURL, AlertURI, a.neustar.Key, a.neustar.DigitalSignature()))
	if err != nil {
		return nil, response.StatusCode, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, response.StatusCode, err
	}
	return data.Data.Items, response.StatusCode, nil
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
