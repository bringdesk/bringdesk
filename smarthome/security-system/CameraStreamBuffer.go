package security_system

import (
	"bytes"
	"fmt"
	"io"
)

type CameraStreamBuffer struct {
	buffer     []byte
	bufferSize int
}

func (self *CameraStreamBuffer) ReadLine() (string, error) {
	parts := bytes.SplitN(self.buffer, []byte("\n"), 2)
	if len(parts) == 2 {
		self.buffer = parts[1]
		return string(parts[0]), nil
	} else {
		return "", io.EOF
	}
}

func (self *CameraStreamBuffer) Len() int {
	frameLen := len(self.buffer)
	return frameLen
}

func (self *CameraStreamBuffer) Read(frame []byte) (int, error) {
	frameSize := cap(frame)
	bufferSize := len(self.buffer)
	if bufferSize < frameSize {
		return 0, io.EOF
	}
	copy(frame, self.buffer[:frameSize])
	self.buffer = self.buffer[frameSize:]
	return frameSize, nil
}

func (self *CameraStreamBuffer) Write(chunk []byte) error {
	bufferSize := len(self.buffer)
	chunkSize := len(chunk)
	//log.Printf("CameraStreamBuffer: buffer = %d chunk = %d", bufferSize, chunkSize)
	if bufferSize+chunkSize > self.bufferSize {
		return fmt.Errorf("out of camera buffer")
	}
	self.buffer = append(self.buffer, chunk...)
	return nil
}

func NewCameraStreamBuffer() *CameraStreamBuffer {
	var bufferSize int = 512 * 1024
	newCameraStreamBuffer := &CameraStreamBuffer{
		bufferSize: bufferSize,
	}
	return newCameraStreamBuffer
}
