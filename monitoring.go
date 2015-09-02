package neustar

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// MonitorURI is the endpoint for calls to the monitoring API
	MonitorURI = "monitor/1.0"

	// LocationsURI is the endpoint for locations calls
	LocationsURI = "/locations"

	// SummaryURI is the summary endpoint for summary calls
	SummaryURI = "/summary"
)

// MonitorTypes is a slice of valid monitor types
var MonitorTypes = []string{"RealBrowserUser", "VirtualUser", "dns"}

// BrowserTypes is a slice of valid browser types
var BrowserTypes = []string{"FF", "CHROME", "IE"}

// UpdateIntervals is a slice of valid intervals
var UpdateIntervals = []int{1, 2, 3, 4, 5, 10, 15, 20, 30, 60}

// UpdateMonitorParameters holds the allowed options for updating
// a monitor
type UpdateMonitorParameters struct {
	Name        string `json:"name"`
	Description string `json:"name"`
	Interval    int    `json:"interval"`
	TestScript  string `json:"testScript"`
	Locations   string `json:"locations"`
	AlertPolicy string `json:"alertPolicy"`
	Browser     string `json:"browser"`
	Active      string `json:"Active"`
}

// CreateMonitorParameters holds the parameters needed by the create
// monitor endpoint
type CreateMonitorParameters struct {
	// The name of the monitor
	Name string `json:"name"`

	// A description of what this monitor is for
	Description string `json:"description"`

	// How often the monitoring script will run for each of the locations
	Interval int `json:"interval"`

	// The id of the test script that this monitor should run
	TestScript string `json:"testScript"`

	// A CSV list of locations that this monitor should run from
	Locations string `json:"Locations"`

	// The id of the alert policy that this monitor should run
	AlertPolicy string `json:"alertPolicy"`

	// Specifies the browser type that this monitor should use. Note: IE is
	// available for Enterprise customers only
	Browser string `json:"browser"`

	// Enables or disables this monitor from taking samples
	Active string `json:"active"`

	// Set to network monitor type such as 'dns'. See related settings below.
	// Leave this blank for script-based monitors. Note: this interface will
	// not allow you to test network monitor creation. Please use your API client.
	Type string `json:"type"`

	DNSSettings  DNSSettings  `json:"dnsSettings"`
	PingSettings PingSettings `json:"pingSettings"`
	PopSettings  PopSettings  `json:"popSettings"`
	PortSettings PortSettings `json:"portSettings"`
	SMTPSettings SMTPSettings `json:"smtpSettings"`
}

// AggregateParameters holds the allowed options for getting
// aggregate sample data
type AggregateParameters struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Offset    int    `json:"offset"`
	Frequency string `json:"frequency"`
	GroupBy   string `json:"groupBy"`
}

// AggregateSampleDataResponse holds the returned data fro mthe call
type AggregateSampleDataResponse struct {
	Count      int         `json:"count"`
	Uptime     float64     `json:"uptime"`
	Min        int         `json:"min"`
	Max        int         `json:"max"`
	Date       interface{} `json:"date"`
	Avg        float64     `json:"avg"`
	STDDev     float64     `json:"stdDev"`
	Location   string      `json:"location"`
	StepName   string      `json:"stepName"`
	StepNumber int         `json:"stepNumber"`
	TP50       float64     `json:"tp50"`
	TP90       float64     `json:"tp90"`
}

// DNSSettings is a an object containing all DNS-related settings:
// {"timeout": int, "lookups": array}. The "lookups" array contains
// JSON objects with this format: {"lookupType": string ("A" or "AAAA"),
// "authoritative": boolean, "hostname": string, "dnsServer": string, "expectedIps":
// string of comma-separated IP addresses}
type DNSSettings struct {
	LookupType    string `json:"lookupType"`
	Authoritative bool   `json:"authoritative"`
	Hostname      string `json:"hostname"`
	DNSServer     string `json:"dnsServer"`
	ExpectedIPs   string `json:"expectedIps"`
}

// PingSettings is an object containing all PING-related settings: {"timeout": int, "host": string}.
type PingSettings struct {
	Timeout int    `json:"timeout"`
	Host    string `json:"host"`
}

// PortSettings holds port settings
type PortSettings struct {
	Timeout          int    `json:"timeout"`
	Server           string `json:"server"`
	Port             int    `json:"port"`
	Protocol         string `json:"protocol"`
	Command          string `json:"command"`
	ExpectedResponse string `json:"expected_response"`
	DataFormat       string `json:"data_format"`
}

// SMTPSettings holds SMTP setttings
type SMTPSettings struct {
	Timeout int    `json:"timeout"`
	Server  string `json:"server"`
	Email   string `json:"email"`
}

// PopSettings holds port settings
type PopSettings struct {
	Timeout  int    `json:"timeout"`
	Server   string `json:"server"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Monitoring holds monitoring config
type Monitoring struct {
	neustar *Neustar
}

// Monitor hold monitoring data
type Monitor struct {
	// The ID of the monitor
	ID string `json:"id"`

	// The ID of the alerting policy associated with this monitor
	AlertPolicy interface{} `json:"alertPolicy"`

	// A list of monitoring locations that this monitor is run from
	Locations []string `json:"locations"`

	// The version, id and name of the script associated with this monitor
	Script Script `json:"script"`

	// The description of the monitor
	Description string `json:"description"`

	// How often this monitor runs
	Interval int `json:"interval"`

	// The name of this monitor
	Name string `json:"name"`

	// The time of the last monitoring sample for this monitor
	LastSampleAt interface{} `json:"lastSampleAt"`

	// Describes whether this monitor is actively monitoring or not
	Active bool `json:"active"`

	// Whether this monitor is in a maintenance window or not
	InMaintenanceWindow bool `json:"inMaintenanceWindow"`

	// Describes the type of browser that monitor is using, 'FF' for
	// Firefox or 'CHROME' for Chrome, or 'IE' for Internet Explorer
	Browser string `json:"browser"`

	// The type of monitor ('RealBrowserUser', 'VirtualUser', 'dns')
	Type string `json:"type"`

	SMTPSettings SMTPSettings `json:"smtpSettings"`
	SLASettings  interface{}  `json:"slaSettings,omitempty"`
	DNSSettings  DNSSettings  `json:"dnsSettings,omitempty"`
	PopSettings  PopSettings  `json:"popSettings,omitempty"`
	PortSettings PortSettings `json:"portSettings,omitempty"`
	PingSettings PingSettings `json:"pingSettings"`
}

// NewMonitor creates a new Monitoring object
func NewMonitor(neustar *Neustar) *Monitoring {
	return &Monitoring{
		neustar: neustar,
	}
}

// Create creates a new monitor and returns the monitor id of the newly
// created monitor. Name, interval, testScript and locations are required.
// Use the Get Monitoring Locations api to retrieve a list of monitoring locations.
func (m *Monitoring) Create() {}

// List retrieves a list of all monitors associated with your account,
//along with information about each. The monitor id that is returned
// is used to make other api calls.
func (m *Monitoring) List() ([]Monitor, int, error) {
	var response *http.Response
	var data map[string]map[string][]Monitor
	response, err := http.Get(fmt.Sprintf("%s%s?apikey=%s&sig=%s", BaseURL, MonitorURI, m.neustar.Key, m.neustar.DigitalSignature()))
	if err != nil {
		return nil, response.StatusCode, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, response.StatusCode, err
	}
	return data["data"]["items"], response.StatusCode, nil
}

// Get retrieves information for a specific monitor associated with your
// account. The monitor id that is returned is used to make other api calls.
func (m *Monitoring) Get(id string) ([]Monitor, int, error) {
	var response *http.Response
	var data map[string]map[string][]Monitor
	response, err := http.Get(fmt.Sprintf("%s%s/%s?apikey=%s&sig=%s", BaseURL, MonitorURI, id, m.neustar.Key, m.neustar.DigitalSignature()))
	if err != nil {
		return nil, response.StatusCode, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, response.StatusCode, err
	}
	return data["data"]["items"], response.StatusCode, nil
}

// Update changes some or all of the parameters of an existing monitor.
// Requires the monitor ID retrieved from the List Monitors api.
func (m *Monitoring) Update() {}

// Delete deletes the given monitor, stopping it from monitoring and removing
// all its monitoring data.
func (m *Monitoring) Delete(id string) (int, error) {
	var response *http.Response
	response, err := http.Get(fmt.Sprintf("%s%s/%s?apikey=%s&sig=%s", BaseURL, MonitorURI, id, m.neustar.Key, m.neustar.DigitalSignature()))
	if err != nil {
		return response.StatusCode, err
	}
	if response.StatusCode != 200 {
		return response.StatusCode, nil
	}
	return response.StatusCode, nil
}

// Samples returns all samples associated to this monitor for a given time period.
// This data is returned at a high level, which timing for the overall sample. To
// get the details for the specific sample, call the get raw sample data api. At a
// maximum, this api will return 2000 samples. If there are more than 2000 results
// returned, the 'more' field will be set to true and you can make another api call
// specifying an offset which would be equal to the number of results returned in the
// first api call plus the offset of that call.
func (m *Monitoring) Samples() ([]string, int, error) {
	var response *http.Response
	var data map[string]map[string][]string
	response, err := http.Get(fmt.Sprintf("%s%s?apikey=%s&sig=%s", BaseURL, MonitorURI, m.neustar.Key, m.neustar.DigitalSignature()))
	if err != nil {
		return nil, response.StatusCode, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, response.StatusCode, err
	}
	return data["data"]["items"], response.StatusCode, nil
}

// RawSampleData retrieves the raw, HTTP Archive (HAR) data for a particular sample
func (m *Monitoring) RawSampleData(monitorID, sampleID string) {}

// AggregateSampleData retrieves the aggregated sample information for a given period
// of time. You can choose to aggregate the data for each hour or each day. This is
// more effecient than getting all the individual samples for a period of time and
// performing the aggregation yourself.
func (m *Monitoring) AggregateSampleData(monitorID string, asd AggregateSampleDataResponse) (int, error) {
	var response *http.Response
	response, err := http.Get(fmt.Sprintf("%s%s/%s?apikey=%s&sig=%s", BaseURL, MonitorURI, monitorID, m.neustar.Key, m.neustar.DigitalSignature()))
	if err != nil {
		return response.StatusCode, err
	}
	if response.StatusCode != 200 {
		return response.StatusCode, nil
	}
	return response.StatusCode, nil
}

// Summary provides the monitor summary api returns all of the data that is found when looking at your
// list of monitors in the web portal. This includes things such as the average load
// time, sample count and uptime for the day, week, month or year, the last time an
// error occurred, and the last error message.
func (m *Monitoring) Summary(monitorID string) {}

// Locations gets a list of all monitoring locations available
func (m *Monitoring) Locations() ([]string, int, error) {
	var response *http.Response
	var data map[string]map[string][]string
	response, err := http.Get(fmt.Sprintf("%s%s%s?apikey=%s&sig=%s", BaseURL, MonitorURI, LocationsURI, m.neustar.Key, m.neustar.DigitalSignature()))
	if err != nil {
		return nil, response.StatusCode, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, response.StatusCode, err
	}
	return data["data"]["items"], response.StatusCode, nil
}

// ValidMonitorType validates the given monitor type is valid
func ValidMonitorType(monitorType string) bool {
	for _, i := range MonitorTypes {
		if i == monitorType {
			return true
		}
	}
	return false
}

// ValidBrowserType validates the given browser type is valid
func ValidBrowserType(browserType string) bool {
	for _, i := range BrowserTypes {
		if i == browserType {
			return true
		}
	}
	return false
}

// ValidUpdateInterval validates the given interval is valid
func ValidUpdateInterval(interval int) bool {
	for _, i := range UpdateIntervals {
		if i == interval {
			return true
		}
	}
	return false
}
