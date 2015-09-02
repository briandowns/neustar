package neustar

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

const (
	// BaseURL is the base URL endpoint
	BaseURL = "http://api.neustar.biz/performance/"

	// Version is the current version of this library
	Version = "0.1"
)

// ReturnedAPIError represents what the API returns on error
type ReturnedAPIError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Param   string `json:"param"`
}

// Neustar holds the provided access keys
type Neustar struct {
	// Neustar API key
	Key string

	// Neustar API secret
	Secret string
}

// NewNeustar creates a new Neustar object
func NewNeustar(key, secret string) *Neustar {
	return &Neustar{
		Key:    key,
		Secret: secret,
	}
}

// DigitalSignature creates an MD5 hash of the key, the secret and a timestamp
func (n *Neustar) DigitalSignature() string {
	now := time.Now()
	epoch := now.Unix()
	data := md5.Sum(
		[]byte(fmt.Sprintf("%s%s%s",
			n.Key, n.Secret, strconv.FormatInt(epoch, 10)),
		),
	)
	return fmt.Sprintf("%x", data)
}

// APIError represents what the API returns on error
var APIError map[string]ReturnedAPIError

// MonitoringErrorCodes holds API returned error codes
var MonitoringErrorCodes = map[string]string{
	"MON_0000": "Resource :resource not available",
	"MON_0001": "Item with Id :itemid not found",
	"MON_0002": "Duplicate name :name found.",
	"MON_0003": ":missingfield is/are required fields",
	"MON_0004": ":value is an invalid value.",
	"MON_0005": ":value is an invalid value type. Type :valuetype expected",
	"MON_0006": ":field is an invalid field",
	"MON_0007": "Valid :field id not supplied",
	"MON_0008": "Empty body received in POST",
	"MON_0009": "Empty body received in PUT",
	"MON_0010": "System cannot complete the request",
	"MON_0011": "User does not have permission to perform the action",
	"MON_9999": "An internal error has occured",
}

// RealUserMeasurementsErrorCodes holds API returned error codes
var RealUserMeasurementsErrorCodes = map[string]string{
	"RUM_000": "Request has been throttled",
	"RUM_001": "Database internal error",
	"RUM_002": "Generic internal error",
	"RUM_003": "Database found data inconsistencies",
	"RUM_004": "Generic Forbidden error",
	"RUM_005": "Mandatory parameter is missing",
	"RUM_006": "Parameter is not valid",
}
