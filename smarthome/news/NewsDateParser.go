package news

import "time"

type NewsDateParser struct {
	chars  []rune    /* Date runes             */
	pos    int       /* Parser cursor position */
	result time.Time /* Result                 */
}

func NewNewsDateParser() *NewsDateParser {
	return new(NewsDateParser)
}

func (self *NewsDateParser) Parse(pubDate string) (time.Time, error) {
	/* Step 1. Save runes */

	/* Step 2. Parse process*/
	self.parseWeekday() // "Thu"
	self.parseComma()   // ","
	self.parseSpace()   // " "
	self.parseDate()    // "23 Dec 2021"
	self.parseSpace()   // " "
	self.parseTime()    // "21:23:33"
	self.parseSpace()   // " "
	self.parseZone()    // "+0300"

	return self.result, nil
}

func (self *NewsDateParser) parseTime() error {
	self.parseNumber()  // "23"
	self.parseRune(':') // ":"
	self.parseNumber()  // "43"
	self.parseRune(':') // ":"
	self.parseNumber()  // "16"
	return nil
}

func (self *NewsDateParser) parseDate() error {
	self.parseNumber() // "23"
	self.parseSpace()  // " "
	self.parseMonth()  // "Dec"
	self.parseSpace()  // " "
	self.parseNumber() // 2021
	return nil
}

func (self *NewsDateParser) parseWeekday() error {
	self.peekString("Mon")
	self.peekString("Thu")
	return nil
}

func (self *NewsDateParser) parseComma() error {
	return nil
}

func (self *NewsDateParser) parseSpace() error {
	return nil
}

func (self *NewsDateParser) parseNumber() error {
	return nil
}

func (self *NewsDateParser) parseMonth() {

}

func (self *NewsDateParser) parseRune(ch rune) {

}

func (self *NewsDateParser) parseZone() error {
	return nil
}

func (self *NewsDateParser) peekString(value string) {

}
