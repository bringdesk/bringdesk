package security_system

import (
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"log"
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
	secureCameraAddress string
	texture             *sdl.Texture
	surface             *sdl.Surface
	cameraStream        *CameraStream
}

func NewSecuritySystemWidget() *SecuritySystemWidget {
	newSecuritySystemWidget := new(SecuritySystemWidget)
	newSecuritySystemWidget.recoverCameraAddress()
	newSecuritySystemWidget.start()
	return newSecuritySystemWidget
}

func (self *SecuritySystemWidget) start() {
	go func() {
		for {
			self.cameraStream = NewCameraStream()
			self.cameraStream.SetAddr(self.secureCameraAddress)
			self.cameraStream.SetProcessFrame(self.processFrame)
			err1 := self.cameraStream.StartStream()
			if err1 != nil {
				log.Printf("Secure widget: err = %#v", err1)
			}
			self.cameraStream.Close()
			/* Wait 10 sec. */
			time.Sleep(10 * time.Second)
		}
	}()
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

func (self *SecuritySystemWidget) ProcessEvent(e *evt.Event) {
}

func (self *SecuritySystemWidget) Render() {
	self.BaseWidget.Render()

	mainRenderer := ctx.GetRenderer()

	newRect := sdl.Rect{X: int32(self.X), Y: int32(self.Y), W: int32(self.Width), H: int32(self.Height)}

	self.frameSync.Lock()
	defer self.frameSync.Unlock()

	/**/
	if self.surface != nil {
		if self.texture != nil {
			self.texture.Destroy()
			self.texture = nil
		}
		texture, err2 := mainRenderer.CreateTextureFromSurface(self.surface)
		if err2 != nil {
			panic(err2)
		}
		self.texture = texture
		//
		mainRenderer.Copy(self.texture, nil, &newRect)
	}

	/* Draw red overlay on disconnect state */
	var cameraActive bool = false
	if self.cameraStream != nil {
		cameraActive = self.cameraStream.IsStreamActive()
	}
	if !cameraActive {
		mainRenderer.SetDrawColor(255, 0, 0, 128)
		mainRenderer.FillRect(&newRect)
	}

}

func (self *SecuritySystemWidget) processFrame(frame []byte) {

	self.frameSync.Lock()
	defer self.frameSync.Unlock()

	self.frame = frame

	/* Create surface */
	rwops, err1 := sdl.RWFromMem(self.frame)
	if err1 != nil {
		log.Printf("error with RWFromMem")
	}
	defer rwops.Close()

	/**/
	if self.surface != nil {
		self.surface.Free()
		self.surface = nil
	}

	//

	surface, err1 := img.LoadJPGRW(rwops)
	if err1 != nil {
		panic(err1)
	}
	self.surface = surface

}
