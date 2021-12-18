package bank

import (
	"log"
	"time"
)

type ConvertRate struct {
	currency    rune
	convertRate float64
}

type BankWidget struct {
	convertRates   []*ConvertRate /* Conversation rate     */
	updateIsActive bool           /* Update routine        */
}

func NewBankWidget() *BankWidget {
	newBankWidget := new(BankWidget)
	newBankWidget.updateIsActive = true
	go newBankWidget.updateData()
	return newBankWidget
}

func (self *BankWidget) StateRestore() {
	/* Register USD rate conversation */
	convertRate1 := new(ConvertRate)
	convertRate1.currency = '$'
	convertRate1.convertRate = 77.45
	self.convertRates = append(self.convertRates, convertRate1)
	/* Register EURO rate conversation */
	convertRate2 := new(ConvertRate)
	convertRate2.currency = '€'
	convertRate2.convertRate = 87.45
	self.convertRates = append(self.convertRates, convertRate2)
}

func (self *BankWidget) updateData() {
	for self.updateIsActive {
		log.Printf("BankWidget: Update currency conversion rates is start")
		/* Receive conversion XML structure */
		// TODO - ... http://www.cbr.ru/scripts/XML_daily.asp
		/* Parse XML structure */
		// TODO - ...
		/* Wait 10 minute */
		time.Sleep(10 * time.Minute)
	}
}

func (self *BankWidget) Render() {

}
