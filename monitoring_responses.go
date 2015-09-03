package neustar

// AggregateSampleDataResponse holds the returned data from the call
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
