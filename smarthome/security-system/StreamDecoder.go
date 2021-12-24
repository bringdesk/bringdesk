package security_system

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type StreamState int

const (
	StreamStateHeader = StreamState(1)
	StreamStateFrame  = StreamState(2)
	StreamStateReset  = StreamState(3)
)

type StreamDecoder struct {
	stream          bytes.Buffer /* */
	streamState     StreamState  /* */
	streamFrameSize int          /* */
	frameCount      int
	processFrame    func(frame []byte)
}

type ProcessState int

const (
	ProcessStateUnknown = ProcessState(0)
	ProcessStatePending = ProcessState(1)
	ProcessStateReady   = ProcessState(2)
)

func (self *StreamDecoder) processData() (ProcessState, error) {

	if self.streamState == StreamStateHeader {
		//--myboundary
		//Content-Type: image/jpeg
		//Content-Length: 61775
		raw, err1 := self.stream.ReadString('\n')
		if err1 != nil {
			log.Printf("Read error: err = %#v", err1)
			return ProcessStatePending, nil
		}
		row := strings.Trim(raw, "\r\n")
		//log.Printf("row = \n```\n%s\n```", row)
		if strings.HasPrefix(row, "--") {
			// TODO - check boundary name ...
			return ProcessStateReady, nil
		} else if strings.Contains(row, ":") {
			parts := strings.SplitN(row, ":", 2)
			if len(parts) == 2 {
				headerName := parts[0]
				headerValue := parts[1]
				if headerName == "Content-Length" {
					newHeaderValue := strings.Trim(headerValue, " \t\n\r")
					//log.Printf("ContentLength is %q", newHeaderValue)

					size, err1 := strconv.ParseUint(newHeaderValue, 10, 64)
					if err1 != nil {
						log.Printf("Parse error: err = %q", err1)
					}
					self.streamFrameSize = int(size)
				}
			}
			return ProcessStateReady, nil
		} else if row == "" {
			// End headers start content ...
			//log.Printf("Start frame #%d processing...", self.frameCount)
			self.streamState = 2
			return ProcessStateReady, nil
		}

		return ProcessStateUnknown, fmt.Errorf("wrong invariant")

	} else if self.streamState == StreamStateFrame {

		if self.streamFrameSize == 0 {
			return ProcessStateUnknown, fmt.Errorf("no frame size")
		}

		if self.stream.Len() > self.streamFrameSize {
			var frame []byte = make([]byte, self.streamFrameSize)
			_, err := self.stream.Read(frame)
			if err != nil {
				log.Printf("Frame %d complete with error: err = %#v", err)
			}
			//
			if self.processFrame != nil {
				self.processFrame(frame)
			}
			//
			self.streamFrameSize = 0
			self.streamState = StreamStateReset
			self.frameCount = self.frameCount + 1
			return ProcessStateReady, nil
		}

		return ProcessStatePending, nil

	} else if self.streamState == StreamStateReset {
		self.stream.ReadString('\n')
		self.streamState = StreamStateHeader
		//log.Printf("Total frame count: %d", self.frameCount)
		return ProcessStateReady, nil
	}

	return ProcessStateUnknown, fmt.Errorf("wrong state")

}

func (self *StreamDecoder) SetProcessFrame(processFrame func([]byte)) {
	self.processFrame = processFrame
}

func (self *StreamDecoder) Decode() error {

	for {
		processState, err1 := self.processData()
		if err1 != nil {
			return err1
		}
		if processState == ProcessStatePending {
			break
		}
	}

	return nil

}

func (self *StreamDecoder) Write(chunk []byte) error {

	if self.stream.Len() > 128*1024 {
		//log.Printf("Camera protocol framing error. No memory.")
		panic("Camera protocol framing error. No memory.")
	}
	self.stream.Write(chunk)

	//log.Printf("Stream new data size = %d", self.stream.Len())

	return nil
}

func (self *StreamDecoder) processFrameSave(frame []byte) {
	// TODO - make process frame ...
	frameFile := fmt.Sprintf("frame_%d.jpg", self.frameCount)
	stream, err1 := os.Create(frameFile)
	if err1 != nil {
		return
	}
	defer stream.Close()
	stream.Write(frame)
}

func NewStreamDecoder() *StreamDecoder {
	newStreamDecoder := new(StreamDecoder)
	newStreamDecoder.streamState = StreamStateHeader
	return newStreamDecoder
}
