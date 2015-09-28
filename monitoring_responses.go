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

// SLASettings hols SLA settings
type SLASettings struct {
	RunningAvgDuration interface{} `json:"runningAvgDuration"`
	Uptime             float64     `json:"uptime"`
	LoadTime           float64     `json:"loadtime"`
}

// AggregateSampleDataResponse holds the return from the API list call
type AggregateSampleDataResponse struct {
	Data struct {
		Count int                       `json:"count"`
		Items []AggregateSampleResponse `json:"items"`
	} `json:"data"`
}

// AggregateSampleResponse holds the returned data from the call
type AggregateSampleResponse struct {
	Count      int    `json:"count"`
	Uptime     string `json:"uptime"`
	Min        int    `json:"min"`
	Max        int    `json:"max"`
	Date       string `json:"date"`
	Avg        string `json:"avg"`
	STDDev     string `json:"stdDev"`
	Location   string `json:"location"`
	StepName   string `json:"stepName"`
	StepNumber int    `json:"stepNumber"`
	TP50       int    `json:"tp50"`
	TP90       int    `json:"tp90"`
}

// SamplesDataResponse holds a response from a call to the Samples endpoint
type SamplesDataResponse struct {
	Data struct {
		Count int `json:"count"`
		Items []struct {
			Status          string `json:"status"`
			BytesReceived   int    `json:"bytesReceived"`
			ErrorLineNumber int    `json:"errorLineNumber"`
			Location        string `json:"location"`
			StartTime       string `json:"startTime"`
			Duration        int    `json:"duration"`
			ID              string `json:"id"`
		} `json:"items"`
	} `json:"data"`
}

// AlertPolicy holds the alert policy data
type AlertPolicy struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Monitor hold monitoring data
type Monitor struct {
	// The ID of the monitor
	ID string `json:"id"`

	// The ID of the alerting policy associated with this monitor
	//AlertPolicy interface{} `json:"alertPolicy"`
	AlertPolicy AlertPolicy

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
	LastSampleAt string `json:"lastSampleAt"`

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
	SLASettings  SLASettings  `json:"slaSettings,omitempty"`
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
	LastErrorType    string `json:"lastErrorType"`
	LastErrorMessage string `json:"lastErrorMessage"`
}

// CreateMonitorResponse contains the response from a call to create a new monitor
type CreateMonitorResponse struct {
	ID      string `json:"id"`
	Created string `json:"created"`
}

// RawSampleDataResponse holds the response given when the RawSampleData function is
// called
type RawSampleDataResponse struct {
	Data struct {
		AgentIPAddr   string `json:"agentIpAddr"`
		Browser       string `json:"browser"`
		BytesReceived int    `json:"bytesReceived"`
		Data          struct {
			Har struct {
				Log struct {
					Steps []struct {
						Duration       int           `json:"duration"`
						Label          string        `json:"label"`
						NameValuePairs []interface{} `json:"nameValuePairs"`
						StartTime      string        `json:"startTime"`
						Step           int           `json:"step"`
						TimePaused     int           `json:"timePaused"`
					} `json:"_steps"`
					Browser struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"browser"`
					Creator struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"creator"`
					Entries []struct {
						WSID    int      `json:"_wsid"`
						Cache   struct{} `json:"cache"`
						Pageref string   `json:"pageref"`
						Request struct {
							BodySize int           `json:"bodySize"`
							Cookies  []interface{} `json:"cookies"`
							Headers  []struct {
								Name  string `json:"name"`
								Value string `json:"value"`
							} `json:"headers"`
							HeadersSize int           `json:"headersSize"`
							HTTPVersion string        `json:"httpVersion"`
							Method      string        `json:"method"`
							QueryString []interface{} `json:"queryString"`
							URL         string        `json:"url"`
						} `json:"request"`
						Response struct {
							BodySize int `json:"bodySize"`
							Content  struct {
								MimeType string `json:"mimeType"`
								Size     int    `json:"size"`
							} `json:"content"`
							Cookies []struct {
								Expires string `json:"expires"`
								Name    string `json:"name"`
								Path    string `json:"path"`
								Value   string `json:"value"`
							} `json:"cookies"`
							Headers []struct {
								Name  string `json:"name"`
								Value string `json:"value"`
							} `json:"headers"`
							HeadersSize int    `json:"headersSize"`
							HTTPVersion string `json:"httpVersion"`
							RedirectURL string `json:"redirectURL"`
							Status      int    `json:"status"`
							StatusText  string `json:"statusText"`
						} `json:"response"`
						ServerIPAddress string `json:"serverIPAddress"`
						StartedDateTime string `json:"startedDateTime"`
						Time            int    `json:"time"`
						Timings         struct {
							Blocked int `json:"blocked"`
							Connect int `json:"connect"`
							DNS     int `json:"dns"`
							Receive int `json:"receive"`
							Send    int `json:"send"`
							Ssl     int `json:"ssl"`
							Wait    int `json:"wait"`
						} `json:"timings"`
					} `json:"entries"`
					Location string `json:"location"`
					Pages    []struct {
						ID          string `json:"id"`
						PageTimings struct {
							DOMComplete                int `json:"_domComplete"`
							DOMContentLoadedEventEnd   int `json:"_domContentLoadedEventEnd"`
							DOMContentLoadedEventStart int `json:"_domContentLoadedEventStart"`
							DOMInteractive             int `json:"_domInteractive"`
							DOMLoading                 int `json:"_domLoading"`
							LoadEventEnd               int `json:"_loadEventEnd"`
							LoadEventStart             int `json:"_loadEventStart"`
							OnContentLoad              int `json:"onContentLoad"`
							OnLoad                     int `json:"onLoad"`
						} `json:"pageTimings"`
						StartedDateTime string `json:"startedDateTime"`
						Title           string `json:"title"`
					} `json:"pages"`
					Version string `json:"version"`
				} `json:"log"`
			} `json:"har"`
			MonitorID  string `json:"monitorId"`
			Screenshot string `json:"screenshot"`
		} `json:"data"`
		Duration   string        `json:"duration"`
		Items      []interface{} `json:"items"`
		Location   string        `json:"location"`
		Offset     int           `json:"offset"`
		ScriptInfo struct {
			LineNumber    interface{} `json:"lineNumber"`
			ScriptBody    string      `json:"scriptBody"`
			ScriptLink    string      `json:"scriptLink"`
			ScriptName    string      `json:"scriptName"`
			ScriptVersion string      `json:"scriptVersion"`
		} `json:"scriptInfo"`
		ScriptName string `json:"scriptName"`
		StartTime  string `json:"startTime"`
		Status     string `json:"status"`
		StatusCode int    `json:"statusCode"`
		Total      int    `json:"total"`
	} `json:"data"`
}
