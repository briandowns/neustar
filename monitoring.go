package neustar

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	// MonitorURI is the endpoint for calls to the monitoring API
	MonitorURI = "monitor/1.0"

	// LocationsURI is the endpoint for locations calls
	LocationsURI = "/locations"

	// SummaryURI is the endpoint for summary calls
	SummaryURI = "/summary"

	// AggregateURI is the endpoint for aggregate calls
	AggregateURI = "/aggregate"
)

// MonitorTypes is a slice of valid monitor types
var MonitorTypes = []string{"RealBrowserUser", "VirtualUser", "dns"}

// BrowserTypes is a slice of valid browser types
var BrowserTypes = []string{"FF", "CHROME", "IE"}

// UpdateIntervals is a slice of valid intervals
var UpdateIntervals = []int{1, 2, 3, 4, 5, 10, 15, 20, 30, 60}

// AggregateSampleDataFrequency is a slice of valid frequencies
var AggregateSampleDataFrequency = []string{"day", "hour"}

// AggregateSampleGroupBy is a slice of valid groupBy parameters
var AggregateSampleGroupBy = []string{"location", "step"}

// Monitoring holds monitoring config
type Monitoring struct {
	neustar *Neustar
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
	response, err := http.Get(fmt.Sprintf(
		"%s%s?apikey=%s&sig=%s",
		BaseURL, MonitorURI, m.neustar.Key, m.neustar.DigitalSignature()))
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
	response, err := http.Get(fmt.Sprintf(
		"%s%s/%s?apikey=%s&sig=%s",
		BaseURL, MonitorURI, id, m.neustar.Key, m.neustar.DigitalSignature()))
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
	response, err := http.Get(fmt.Sprintf(
		"%s%s/%s?apikey=%s&sig=%s",
		BaseURL, MonitorURI, id, m.neustar.Key, m.neustar.DigitalSignature()))
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
	response, err := http.Get(fmt.Sprintf(
		"%s%s?apikey=%s&sig=%s",
		BaseURL, MonitorURI, m.neustar.Key, m.neustar.DigitalSignature()))
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
func (m *Monitoring) AggregateSampleData(monitorID string, asp *AggregateSampleParameters) ([]AggregateSampleDataResponse, int, error) {
	var response *http.Response
	v, err := query.Values(asp)
	if err != nil {
		return nil, 0, err
	}
	var data map[string]map[string][]AggregateSampleDataResponse
	response, err = http.Get(fmt.Sprintf(
		"%s%s/%s%s?%s&apikey=%s&sig=%s",
		BaseURL, MonitorURI, monitorID, AggregateURI, v.Encode(), m.neustar.Key, m.neustar.DigitalSignature()))
	if err != nil {
		return nil, response.StatusCode, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, response.StatusCode, err
	}
	return data["data"]["items"], response.StatusCode, nil
}

// Summary provides the monitor summary api returns all of the data that is found when looking at your
// list of monitors in the web portal. This includes things such as the average load
// time, sample count and uptime for the day, week, month or year, the last time an
// error occurred, and the last error message.
func (m *Monitoring) Summary(monitorID string) ([]SummaryDataResponse, int, error) {
	var response *http.Response
	var data map[string]map[string][]SummaryDataResponse
	endpoint := fmt.Sprintf("%s%s/%s%s?apikey=%s&sig=%s", BaseURL, MonitorURI, monitorID, SummaryURI, m.neustar.Key, m.neustar.DigitalSignature())
	fmt.Println(endpoint)
	response, err := http.Get(fmt.Sprintf(
		"%s%s/%s%s?apikey=%s&sig=%s",
		BaseURL, MonitorURI, monitorID, SummaryURI, m.neustar.Key, m.neustar.DigitalSignature()))
	if err != nil {
		return nil, response.StatusCode, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, response.StatusCode, err
	}
	return data["data"]["items"], response.StatusCode, nil
}

// Locations gets a list of all monitoring locations available
func (m *Monitoring) Locations() ([]string, int, error) {
	var response *http.Response
	var data map[string]map[string][]string
	response, err := http.Get(fmt.Sprintf(
		"%s%s%s?apikey=%s&sig=%s",
		BaseURL, MonitorURI, LocationsURI, m.neustar.Key, m.neustar.DigitalSignature()))
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

// ValidAggregateSampleDataFrequency validates the given frequency is valid
func ValidAggregateSampleDataFrequency(frequency string) bool {
	for _, i := range AggregateSampleDataFrequency {
		if i == frequency {
			return true
		}
	}
	return false
}

// ValidAggregateSampleGroupBy validates the given groupBy is valid
func ValidAggregateSampleGroupBy(groupBy string) bool {
	for _, i := range AggregateSampleGroupBy {
		if i == groupBy {
			return true
		}
	}
	return false
}
