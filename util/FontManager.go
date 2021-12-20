package util

import (
	"github.com/veandco/go-sdl2/ttf"
	"log"
	"os"
	"path"
)

type FontManager struct {
	searchPaths []string
	useCount    int
	fontAliases map[string]string
}

func NewFontManager() *FontManager {
	newFontManager := &FontManager{
		fontAliases: make(map[string]string),
	}
	return newFontManager
}

func (self *FontManager) registerFont(name string, fontPath string) {
	log.Printf("Register fonr: name = %s path = %s", name, fontPath)
	self.fontAliases[name] = fontPath
}

func (self *FontManager) discoverFonts(searchPath string) error {
	files, err1 := os.ReadDir(searchPath)
	if err1 != nil {
		return err1
	}
	for _, file := range files {
		newFontPath := path.Join(searchPath, file.Name())
		log.Printf("Detect %s", newFontPath)
		newFont, err2 := ttf.OpenFont(newFontPath, 0)
		if err2 != nil {
			log.Printf("Load font error: err = %#v", err2)
		}
		fontName := newFont.FaceFamilyName()
		newFont.Close()
		//
		self.registerFont(fontName, newFontPath)
	}
	return nil
}

func (self *FontManager) SetSearchPath(searchPath string) {
	/* Step 1. Save search paths */
	self.searchPaths = append(self.searchPaths, searchPath)
	/* Step 2. Discover fonts */
	self.discoverFonts(searchPath)
}

func (self *FontManager) Dump() {
	log.Printf("FontMannager: use count = %d", self.useCount)
}

func (self *FontManager) GetFontByName(name string) string {
	for fntName, fntPath := range self.fontAliases {
		if fntName == name {
			return fntPath
		}
	}
	return ""
}

func (self *FontManager) Acquire(fontName string, fontSize int) (*ttf.Font, error) {
	self.useCount += 1

	/* Search font by name */
	fontPath := self.GetFontByName(fontName)
	if fontPath == "" {
		fontPath = self.GetFontByIndex(0)
	}
	if fontPath == "" {
		panic("no fonts?")
	}

	/* Acquire resource */
	newFont, err2 := ttf.OpenFont(fontPath, fontSize)
	if err2 != nil {
		return nil, err2
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

func (self *FontManager) GetFontByIndex(index int) string {
	for _, fontPath := range self.fontAliases {
		return fontPath
	}
	return ""
}
