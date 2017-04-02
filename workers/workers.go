package workers

import (
	"log"
	"syscall"

	c "github.com/prayerslayer/tcp/connection"
	http "github.com/prayerslayer/tcp/http"
)

func ConnectionAcceptor(fd int, newConnChan chan<- c.Connection) {
	log.Printf("Awaiting connection...")
	for {
		newfd, sa, err := syscall.Accept(fd)
		if err != nil {
			log.Printf("Error during connection: %+v", err)
			continue
		}
		newConnChan <- c.Connection{
			FileDescriptor: newfd,
			SocketAddr:     sa,
		}
	}
}

func ConnectionWorker(newConnChan <-chan c.Connection, requestChan chan<- http.HTTPRequest) {
	for {
		newConn := <-newConnChan
		request, readErr := newConn.Read()
		if readErr != nil {
			log.Printf("Error during read: %+v", readErr)
			continue // ????
		}
		if request.Len() > 0 {
			httpRequest := http.ParseHTTPRequest(request.String())
			httpRequest.Connection = newConn
			requestChan <- httpRequest
		}
	}
}

func RequestWorker(requests <-chan http.HTTPRequest) {
	for {
		response := &http.HTTPResponse{}
		request := <-requests
		log.Printf("--> %s %s", request.Method, request.Path)
		response.OriginalRequest = request
		http.Middleware(&request, response)
		writeErr := request.Connection.Write(response.ToBytes())
		if writeErr != nil {
			log.Printf("Error during response write %+v", writeErr)
		}
		closeErr := syscall.Close(request.Connection.FileDescriptor)
		if closeErr != nil {
			log.Printf("Could not close socket %+v: %+v", request.Connection.FileDescriptor, closeErr)
		}
	}
}
