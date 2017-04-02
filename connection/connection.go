package connection

import (
	"bytes"
	"syscall"
)

// Connection describes a new socket connection
type Connection struct {
	FileDescriptor int
	SocketAddr     syscall.Sockaddr
}

func (c Connection) Write(data []byte) error {
	bufSize := 16
	buf := make([]byte, bufSize)
	offset := 0
	l := len(data)

	for (offset+1)*bufSize < l {

		// fill buffer
		// could also just generate a slice from the array and convert it to array again
		for i := 0; i < bufSize; i++ {
			idx := offset*bufSize + i
			buf[i] = data[idx]
		}
		// write it
		_, writeErr := syscall.Write(c.FileDescriptor, buf)
		if writeErr != nil {
			return writeErr
		}
		// adjust vars
		offset++
	}
	if offset*bufSize != l {
		// there is leftover stuff
		diff := (l - offset*bufSize)
		diffbuf := make([]byte, diff)
		for i := 0; i < diff; i++ {
			idx := offset*bufSize + i
			diffbuf[i] = data[idx]
		}
		_, writeErr := syscall.Write(c.FileDescriptor, diffbuf)
		if writeErr != nil {
			return writeErr
		}
	}
	return nil
}

func (c Connection) Read() (bytes.Buffer, error) {
	var result bytes.Buffer
	var readErr error
	bytesRead := 0
	bufSize := 16
	buf := make([]byte, bufSize)
	// initial read
	bytesRead, readErr = syscall.Read(c.FileDescriptor, buf)
	if readErr != nil {
		return result, readErr
	}
	if bytesRead <= 0 {
		return result, nil
	}
	result.Write(buf)
	for bytesRead > 0 {
		bytesRead, readErr = syscall.Read(c.FileDescriptor, buf)
		if readErr != nil {
			return result, readErr
		}
		result.Write(buf)
		if bytesRead < bufSize {
			// stop reading once we got less than our buffer size
			bytesRead = 0
		}
	}
	return result, nil

}
