package security_system

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

type SecuritySystemWidget struct {
	widgets.BaseWidget
	frameSync           sync.Mutex
	frame               []byte
	streamActive        bool
	surface             *sdl.Surface
	secureCameraAddress string
}

func NewSecuritySystemWidget() *SecuritySystemWidget {
	newSecuritySystemWidget := new(SecuritySystemWidget)
	newSecuritySystemWidget.recoverCameraAddress()
	go func() {
		for {
			newSecuritySystemWidget.startStream()
			time.Sleep(10 * time.Second)
		}
	}()
	return newSecuritySystemWidget
}

func (self *SecuritySystemWidget) recoverCameraAddress() {

	/* Step 0. Prepare reding user home directory */
	userDirName, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Todoist error API token reading: err = %#v", err)
	}

	/* Step 1. Prepare Todoist token path */
	newTokenPath := path.Join(userDirName, ".sca")
	log.Printf("openweathermap token path: %#v", newTokenPath)

	/* Step 2. Reading content with token */
	content, err := ioutil.ReadFile(newTokenPath)
	if err != nil {
		log.Printf("openweathermap error API token reading: err = %#v", err)
	}
	self.secureCameraAddress = strings.Trim(string(content), " \r\n\t")

}

func (self *SecuritySystemWidget) startStream() {
	client := http.Client{
		Timeout: 5 * time.Minute,
	}
	request, err := http.NewRequest("GET", self.secureCameraAddress, nil)
	request.Header.Add("User-Agent", "FireFox")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	/* Check server response */
	log.Printf("resp = %#v", resp)

	streamDecoder := NewStreamDecoder()
	streamDecoder.SetProcessFrame(self.processFrame)

	// TODO - check "Content-Type":[]string{"multipart/x-mixed-replace; boundary=--myboundary"}

	window := make([]byte, 16*1024) // Make 8 kB chunk memory
	self.streamActive = true
	for {

		size, err1 := resp.Body.Read(window)
		if err1 != nil {
			break
		}
		//log.Printf("Camera stream RX chunk: size = %d", size)

		err2 := streamDecoder.Write(window[:size])
		if err2 != nil {
			break
		}

		err3 := streamDecoder.Decode()
		if err3 != nil {
			break
		}

	}
	self.streamActive = false
	log.Printf("Camers session is complete.")

}

func (self *SecuritySystemWidget) ProcessEvent(e *evt.Event) {
}

func (self *SecuritySystemWidget) Render() {
	self.BaseWidget.Render()

	self.frameSync.Lock()
	defer self.frameSync.Unlock()

	/* Save image */
	mainRenderer := ctx.GetRenderer()

	newRect := sdl.Rect{X: int32(self.X), Y: int32(self.Y), W: int32(self.Width), H: int32(self.Height)}
	texture, _ := mainRenderer.CreateTextureFromSurface(self.surface)
	mainRenderer.Copy(texture, nil, &newRect)
	texture.Destroy()

	/* Draw red overlay on disconnect state */
	if !self.streamActive {
		mainRenderer.SetDrawColor(255, 0, 0, 128)
		mainRenderer.FillRect(&newRect)
	}

}

func (self *SecuritySystemWidget) processFrame(frame []byte) {

	if len(frame) > 0 {
		self.frameSync.Lock()
		defer self.frameSync.Unlock()

		self.frame = frame

		/* Create surface */
		rwops, _ := sdl.RWFromMem(self.frame)

		surface, _ := img.LoadJPGRW(rwops)
		if self.surface != nil {
			self.surface.Free()
		}
		self.surface = surface

		rwops.Close()
		rwops.Free()

	}

}
