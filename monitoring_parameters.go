package neustar

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

// AggregateSampleParameters holds the allowed options for getting
// aggregate sample data
type AggregateSampleParameters struct {
	// An ISO 8601 formatted date string or datetime string representing the
	// start date from which you wish to collect samples. Examples: 2012-03-02 or 2012-03-01T12:00
	StartDate string `url:"startDate"`

	// An ISO 8601 formatted date string or datetime string representing the
	// end date from which you wish to collect samples. Examples: 2012-03-02 or 2012-03-01T12:00
	EndDate string `url:"endDate"`

	// From which position in the return list you wish to start. At most, 2000 records will be returned.
	Offset int `url:"offset"`

	// Aggregation period ('day', 'hour')
	Frequency string `url:"frequency"`

	// When selected, the data will be aggregated by the selected 'groupBy'
	GroupBy string `url:"groupBy"`
}
