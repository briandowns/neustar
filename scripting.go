package neustar

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// ScriptURI is the endpoint for calls to the scripting API
	ScriptURI = "script/1.0"

	// PolicyURI is the endpoint for policy calls
	CreateURI = "/url"

	// AllScripts is the endpoint for AllScripts calls
	AllScriptsURI = "/AllScripts"

	// ValidSciptsURI is the endpiont for ValidScripts calls
	ValidSciptsURI = "/ValidScripts"

	// InvalidScriptsURI is the endpiont for ValidScripts calls
	InvalidScriptsURI = "/InvalidScripts"

	// UploadBodyURI is the endpiont for ValidScripts calls
	UploadBodyURI = "/upload/body"
)

// Script holds a representation of a script
type Script struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Name    string `json:"name"`
}

// ScriptingListResponse holds the response from the list endpoint
type ScriptingListResponse struct {
	// The name of the alert policy
	Name string `json:"name"`

	// The ID of the alert policy
	ID string `json:"id"`

	// The description of the alert policy
	Description string `json:"description"`

	// This flag will return true if the policy is an Advanced Alert Policy
	AdvancedEdit bool `json:"advancedEdit"`

	// The email addresses associated with the alert policy
	EmailAddress []string `json:"emailAddresses"`

	// The number of strikes before triggering an alert.
	Strikes int `json:"strikes"`

	ScriptInput          interface{} `json:"scriptInput"`
	EscalationPolicy     interface{} `json:"escalationPolicy"`
	FeedIDs              interface{} `json:"feedIds"`
	PagerDutyServiceKey  interface{} `json:"pagerDutyServiceKey"`
	PagerDutyConnectName interface{} `json:"pagerDutyConnectionName"`
	PagerDutyAccount     interface{} `json:"pagerDutyAccount"`
	SMS                  interface{} `json:"sms"`
	PolicyGroup          interface{} `json:"policyGroup"`
	History              interface{} `json:"history"`
	AccountID            string      `json:"accountId"`
	AlertType            interface{} `json:"alertType"`
	ScriptBody           string      `json:"scriptBody"`
	LastUser             string      `json:"lastUser"`
	Created              string      `json:"created"`
	InUse                bool        `json:"inUse"`
	Modified             string      `json:"modified"`
	LatestVersion        string      `json:"latestVersion"`
}

// ScriptCreateParameters holds the parameters passed in to
// create a new script
type ScriptCreateParameters struct {
	// The name of the alert policy
	Name string `json:"name"`

	// An array of comma-separated email addresses
	// associated with the alert policy e.g. ["alert@mycompany.com","myemail@gmail.com"].
	EmailAddresses []string `json:"emailAddresses"`

	// The number of strikes before triggering an alert.
	Strikes int `json:"strikes"`

	// A description for the alert policy
	Description string `json:"description"`
}

// Scripting holds scripting config
type Scripting struct {
	neustar *Neustar
}

// NewScript creates a new Scripting object
func NewScript(neustar *Neustar) *Scripting {
	return &Scripting{
		neustar: neustar,
	}
}

// ScriptDataResponse holds the return from the API list call
type ScriptDataResponse struct {
	Data struct {
		Total  int                     `json:"total"`
		Offset int                     `json:"offset"`
		More   bool                    `json:"more"`
		Items  []ScriptingListResponse `json:"items"`
	} `json:"data"`
}

// Create creates a new Alert policy
func (s *Scripting) Create() {}

// List retrieves a list of policies ordered by date in descending order.
func (s *Scripting) List() ([]ScriptingListResponse, int, error) {
	var response *http.Response
	var data ScriptDataResponse
	response, err := http.Get(fmt.Sprintf("%s%s?apikey=%s&sig=%s", BaseURL, ScriptURI, s.neustar.Key, s.neustar.DigitalSignature()))
	if err != nil {
		return nil, response.StatusCode, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, response.StatusCode, err
	}
	return data.Data.Items, response.StatusCode, nil
}

// ListValidTestScripts retrieves a list of valid test scripts
func (s *Scripting) ListValidTestScripts() {}

// ListInvalidTestScripts retrieves a list of invalid test scripts
func (s *Scripting) ListInvalidTestScripts() {}

// UploadTestScriptFile
func (s *Scripting) UploadTestScriptFile() {}

// CloneTestScriptFile
func (s *Scripting) CloneTestScriptFile() {}
