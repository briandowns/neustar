package neustar

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

// PingSettings is an object containing all PING-related settings:
// {"timeout": int, "host": string}.
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

// AggregateSampleDataResponse holds the returned data from the call
type AggregateSampleDataResponse struct {
	Count      int     `json:"count"`
	Uptime     float64 `json:"uptime"`
	Min        int     `json:"min"`
	Max        int     `json:"max"`
	Date       string  `json:"date"`
	Avg        float64 `json:"avg"`
	STDDev     float64 `json:"stdDev"`
	Location   string  `json:"location"`
	StepName   string  `json:"stepName"`
	StepNumber int     `json:"stepNumber"`
	TP50       float64 `json:"tp50"`
	TP90       float64 `json:"tp90"`
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

// SummaryDataResponse holds the response from the Summary endpoint
type SummaryDataResponse struct {
	// The general status of the monitor, is the monitor erroing or in success state
	GeneralStatus string `json:"generalStatus"`

	// The average uptime for the current week, as a percentage
	AvgUptimeWeek float64 `json:"avgUptimeWeek"`

	// The ID for the last sample taken. This ID can be used with the sample or aggregate APIs
	LastSampleID string `json:"lastSampleId"`

	// The average uptime for the current quarter, as a percentage
	AvgUptimeQuarter float64 `json:"avgUptimeQuarter"`

	// The average uptime for the current year, as a percentage
	AvgUptimeYear float64 `json:"avgUptimeYear"`

	// The number of samples this month
	SampleCountMonth int `json:"sampleCountMonth"`

	// The average uptime this quarter, as a percentage
	AvgLoadtimeQuarter int `json:"avgLoadtimeQuarter"`

	// The average load time this day, in milliseconds
	AvgLoadtimeDay int `json:"avgLoadtimeDay"`

	// The number of samples this day
	SampleCountDay int `json:"sampleCountDay"`

	// The number of samples this year
	SampleCountYear int `json:"sampleCountYear"`

	// The last sampe time data
	LastSampleTimePaused int `json:"lastSampleTimePaused"`

	// When the last sample was taken
	LastSampleAt string `json:"lastSampleAt"`

	// The average load time this month, in milliseconds
	AvgLoadtimeMonth int `json:"avgLoadtimeMonth"`

	// The average load time this quarter, in milliseconds
	SampleCountQuarter int `json:"sampleCountQuarter"`

	// The monitor status: either 'Alerting', 'Warning', 'Scheduled', 'Active', 'Maintenance', 'Off'
	Status string `json:"status"`

	// The average load time this week, in milliseconds
	AvgLoadtimeWeek int `json:"avgLoadtimeWeek"`

	// Load time 90th percentile for the previous UTC day, in milliseconds
	TP90 int `json:"tp90"`

	// The last sample status of the monitor
	LastSampleStatus string `json:"lastSampleStatus"`

	// The average uptime this day, as a percentage
	AvgUptimeDay float64 `json:"avgUptimeDay"`

	// The average load time this year, in milliseconds
	AvgLoadtimeYear int `json:"avgLoadtimeYear"`

	// The duration of the last sample, in milliseconds
	LastSampleDuration int `json:"lastSampleDuration"`

	// Load time 50th percentile for the previous UTC day, in milliseconds
	TP50 int `json:"tp50"`

	// The number of samples this week
	SampleCountWeek int `json:"sampleCountWeek"`

	// The average uptime this month, as a percentage
	AvgUptimeMonth float64 `json:"avgUptimeMonth"`

	// When the last error occured
	LastErrorAt      string `json:"lastErrorAt"`
	LastErrorID      string `json:"lastErrorId"`
	LastErrorType    int    `json:"lastErrorType"`
	LastErrorMessage int    `json:"lastErrorMessage"`
}
