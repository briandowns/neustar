package neustar

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Configuration contains configuration.. duh
type Configuration struct {
	API struct {
		Key    string `json:"key"`
		Secret string `json:"secret"`
	} `json:"api"`
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
func LoadConfig(cf string) (*Configuration, error) {
	confFile, err := os.Open(cf)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(confFile)
	conf := &Configuration{}
	err = decoder.Decode(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
