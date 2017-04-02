package main

import (
	"log"
	"strconv"
	"strings"
	"syscall"

	c "github.com/prayerslayer/tcp/connection"
	http "github.com/prayerslayer/tcp/http"
	workers "github.com/prayerslayer/tcp/workers"
)

func stringToIP(ipString string) [4]byte {
	//log.Printf("String to IP: %s", ipString)
	parts := strings.Split(ipString, ".")
	numbers := [4]byte{0, 0, 0, 0}
	for i, part := range parts {
		n, e := strconv.ParseUint(part, 10, 8)
		//log.Printf("%+v => %+v (%+v)", part, n, e)
		if e != nil {
			log.Fatal(e)
		}
		numbers[i] = uint8(n)
	}
	return numbers
}

func main() {
	// socket setup
	fd, socketErr := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if socketErr != nil {
		log.Fatal(socketErr)
	}
	log.Printf("Socket descriptor: %+v", fd)

	bindErr := syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 3000,
		Addr: stringToIP("127.0.0.1"),
	})
	if bindErr != nil {
		log.Fatal(bindErr)
	}
	log.Printf("Bound to 127.0.0.1:3000")

	listenErr := syscall.Listen(fd, 128)
	if listenErr != nil {
		log.Fatal(listenErr)
	}
	log.Printf("Listening...")

	conns := make(chan c.Connection)
	reqs := make(chan http.HTTPRequest)
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go workers.ConnectionAcceptor(fd, conns)
	}

	for j := 0; j < 10; j++ {
		go workers.ConnectionWorker(conns, reqs)
	}

	for k := 0; k < 10; k++ {
		go workers.RequestWorker(reqs)
	}

	<-done
}
