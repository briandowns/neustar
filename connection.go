package neustar

import (
	"fmt"
	"net/http"
	"runtime"
)

// Conn
type Connection struct{}

// NewConnection
func NewConnection() *Conn { return &Conn{} }

// NewRequest
func (c *Conn) NewRequest(method, path, query string) (*Request, error) {
	var uri string
	// If query parameters are provided, the add them to the URL,
	// otherwise, leave them out
	if len(query) > 0 {
		uri = fmt.Sprintf("%s://%s:%s%s?%s", c.Protocol, host, portNum, path, query)
	} else {
		uri = fmt.Sprintf("%s://%s:%s%s", c.Protocol, host, portNum, path)
	}
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "neustar/"+Version+" ("+runtime.GOOS+"-"+runtime.GOARCH+")")

	newRequest := &Request{
		Request:      req,
		hostResponse: hr,
	}
	return newRequest, nil
}
