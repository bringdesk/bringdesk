package util

import (
	"github.com/veandco/go-sdl2/ttf"
	"log"
	"os"
	"path"
)

type FontMeta struct {
	Name string
	Path string
}

type FontCache struct {
	name  string
	size  int
	cache *ttf.Font
}

type FontManager struct {
	searchPaths []string
	useCount    int
	fonts       []FontMeta
	caches      []FontCache
}

func NewFontManager() *FontManager {
	newFontManager := &FontManager{}
	return newFontManager
}

func (self *FontManager) registerFont(name string, fontPath string) {
	log.Printf("Register fonr: name = %s path = %s", name, fontPath)
	/* Step 1. Check exists */
	// TODO - check already exists ...
	/* Step 2. Register */
	newFontMeta := FontMeta{
		Name: name,
		Path: fontPath,
	}
	self.fonts = append(self.fonts, newFontMeta)
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
	for _, fnt := range self.fonts {
		if fnt.Name == name {
			return fnt.Path
		}
	}
	return ""
}

func (self *FontManager) Acquire(fontName string, fontSize int) (*ttf.Font, error) {
	self.useCount += 1

	/* Search in cache */
	for _, c := range self.caches {
		if c.name == fontName && c.size == fontSize {
			return c.cache, nil
		}
	}

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

	/* Save in cache */
	newCache := FontCache{
		name:  fontName,
		size:  fontSize,
		cache: newFont,
	}
	self.caches = append(self.caches, newCache)

	return newFont, nil
}

func (self *FontManager) Release(font *ttf.Font) {
	self.useCount -= 1
	/* Search in cache */
	// TODO - search in chache and release node ...
	/* Release resource */
	//font.Close()
}

func (self *FontManager) GetUseFontCount() int {
	return self.useCount
}

func (self *FontManager) GetFontByIndex(index int) string {
	for idx, fnt := range self.fonts {
		if idx == index {
			return fnt.Path
		}
	}
	return ""
}

func (self *FontManager) GetFontCount() int {
	return len(self.fonts)
}

func (self *FontManager) GetFontCacheCount() int {
	return len(self.caches)
}
