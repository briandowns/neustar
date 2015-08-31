package neustar

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

// API holds the provided access keys
type API struct {
	Key    string
	Secret string
}

// Configuration contains configuration.. duh
type Configuration struct {
	API
}

// DigitalSignature creates an MD5 hash of the key, the secret and a timestamp
func (c *Configuration) DigitalSignature() string {
	now := time.Now()
	epoch := now.Unix()
	data := md5.Sum(
		[]byte(fmt.Sprintf("%s%s%s",
			c.API.Key, c.API.Secret, strconv.FormatInt(epoch, 10)),
		),
	)
	return fmt.Sprintf("%x", data)
}

// LoadConfig builds a config obj
func LoadConfig(key, secret string) *Configuration {
	return &Configuration{
		API{
			Key:    key,
			Secret: secret,
		},
	}
}
