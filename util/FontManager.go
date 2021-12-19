package util

import (
	"github.com/veandco/go-sdl2/ttf"
	"log"
	"path"
)

type FontManager struct {
	searchPaths    []string
	useCount      int
	fontAliases   map[string]string
}

func NewFontManager() *FontManager {
	return new(FontManager)
}

func (self *FontManager) SetSearchPath(searchPath string) {
	self.searchPaths = append(self.searchPaths, searchPath)
}

func (self *FontManager) Dump() {
	log.Printf("FontMannager: use count = %d", self.useCount)
}

func (self *FontManager) Acquire(fontName string, fontSize int) (*ttf.Font, error) {
	self.useCount += 1

	/* Acquire resource */
	fontFile := "PublicSans-Regular.otf"

	var newFont *ttf.Font
	for _, searchPath := range self.searchPaths {
		newPath := path.Join(searchPath, fontFile)
		var err2 error
		newFont, err2 = ttf.OpenFont(newPath, fontSize)
		if err2 == nil {
			// TODO - logging error ..
		}
	}

	return newFont, nil
}

func (self *FontManager) Release(font *ttf.Font) {
	self.useCount -= 1
	/* Release resource */
	font.Close()
}

func (self *FontManager) GetUseFontCount() int {
	return self.useCount
}