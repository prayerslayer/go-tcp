package http

func Middleware(req *HTTPRequest, res *HTTPResponse) {
	// here be many functions o_O
	res.Body = "\"PONG\""
	res.Status = 200
	res.Headers = []string{"Connection: close"}
}
