package security_system

import (
	"log"
	"net/http"
	"time"
)

type CameraStream struct {
	secureCameraAddr string /* Remote camera address */
	streamActive     bool
	processFrame     func(frame []byte)
}

func NewCameraStream() *CameraStream {
	return new(CameraStream)
}

func (self *CameraStream) StartStream() error {

	log.Printf("Camers session is start...")

	client := http.Client{
		Timeout: 5 * time.Minute,
	}
	request, err1 := http.NewRequest("GET", self.secureCameraAddr, nil)
	if err1 != nil {
		return err1
	}
	request.Header.Add("User-Agent", "FireFox")

	resp, err2 := client.Do(request)
	if err2 != nil {
		return err2
	}
	defer resp.Body.Close()

	/* Check server response */
	log.Printf("resp = %#v", resp)

	streamDecoder := NewStreamDecoder()
	streamDecoder.SetProcessFrame(self.processFrame)

	// TODO - check "Content-Type":[]string{"multipart/x-mixed-replace; boundary=--myboundary"}

	window := make([]byte, 16*1024) // Make 8 kB chunk memory
	self.streamActive = true
	var err error
	for {

		size, err1 := resp.Body.Read(window)
		if err1 != nil {
			err = err1
			break
		}
		//log.Printf("Camera stream RX chunk: size = %d", size)

		err2 := streamDecoder.Write(window[:size])
		if err2 != nil {
			err = err2
			break
		}

		err3 := streamDecoder.Decode()
		if err3 != nil {
			err = err3
			break
		}

	}
	self.streamActive = false
	log.Printf("Camers session is complete.")

	return err
}

func (self *CameraStream) SetAddr(secureCameraAddr string) {
	self.secureCameraAddr = secureCameraAddr
}

func (self *CameraStream) SetProcessFrame(processFrame func(frame []byte)) {
	self.processFrame = processFrame
}

func (self *CameraStream) Close() {

}

func (self *CameraStream) IsStreamActive() bool {
	return self.streamActive
}
