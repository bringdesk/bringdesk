package widgets

import "image/color"

type ButtonWidget struct {
	name     string     /* Button name                 */
	callback func()     /* Callback operation on click */
	color    color.RGBA /* Button color                */
}

func (self *ButtonWidget) SetCallback(callback func()) {
	self.callback = callback
}

func NewButtonWidget() *ButtonWidget {
	newButtonWidget := new(ButtonWidget)
	return newButtonWidget
}
