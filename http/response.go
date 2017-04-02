package http

import (
	"bytes"
	"log"
	"strconv"
)

// HTTPResponse represents a response to a http request
type HTTPResponse struct {
	OriginalRequest HTTPRequest
	Body            string
	Status          int
	Headers         []string
}

func (response *HTTPResponse) ToBytes() []byte {
	var serialized bytes.Buffer
	// status line
	serialized.WriteString(response.OriginalRequest.Protocol)
	serialized.WriteString(" ")
	serialized.WriteString(strconv.FormatInt(int64(response.Status), 10))
	serialized.WriteString(" OK")
	serialized.WriteString("\r\n")
	// Headers
	for _, header := range response.Headers {
		serialized.WriteString(header)
		serialized.WriteString("\r\n")
	}
	if len(response.Headers) > 0 {
		serialized.WriteString("\r\n")
	}
	// body
	serialized.WriteString(response.Body)
	log.Printf("<-- %+v", serialized.String())
	return serialized.Bytes()
}
