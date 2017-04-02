package http

import (
	"strings"

	"github.com/prayerslayer/tcp/connection"
)

// HTTPRequest represents a simple marshalled HTTP request
type HTTPRequest struct {
	Method     string
	Path       string
	Protocol   string
	Connection connection.Connection
}

func ParseHTTPRequest(req string) HTTPRequest {
	lines := strings.Split(req, "\r\n")
	if len(lines) > 0 {
		firstLineItems := strings.Split(lines[0], " ")
		return HTTPRequest{
			Method:   firstLineItems[0],
			Path:     firstLineItems[1],
			Protocol: firstLineItems[2],
		}
	}

	return HTTPRequest{
		Method:   "GET",
		Path:     "/",
		Protocol: "HTTP/1.1",
	}

}
