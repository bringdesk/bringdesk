package skin

import "image/color"

type BgType int

const (
	BgTypeColor     = BgType(1)
	BgTypeImage     = BgType(2)
	BgTypeAnimation = BgType(3)
)

type Skin struct {
	bgType       BgType      /* Background color    */
	bgColor      color.Color /* Background color    */
	bgImage      string      /* Background image    */
	acentColor   color.Color /* Main skin color     */
	DisplayIndex int         /* Display index       */
}

func (self *Skin) GetBgImage() string {
	return self.bgImage
}

func (self *Skin) SetBgImage(bgImage string) {
	self.bgImage = bgImage
}

func NewSkin() *Skin {
	return new(Skin)
}
