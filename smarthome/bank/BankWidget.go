package bank

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/bringdesk/bringdesk/widgets"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  string `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type BankResponse struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
}

type ConvertRate struct {
	currency    rune
	convertRate float64
}

type BankWidget struct {
	convertRates   []ConvertRate /* Conversation rate     */
	updateIsActive bool          /* Update routine        */
	updateDate     string        /* Bank update           */

}

func NewBankWidget() *BankWidget {
	newBankWidget := new(BankWidget)
	newBankWidget.updateIsActive = true
	go func() {
		for newBankWidget.updateIsActive {
			/* Update data */
			newBankWidget.updateData()
			/* Wait 10 minute */
			time.Sleep(10 * time.Minute)
		}
	}()
	return newBankWidget
}

func (self *BankWidget) updateData() {
	log.Printf("BankWidget: Update currency conversion rates is start")
	/* Receive conversion XML structure */
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	req, _ := http.NewRequest("GET", "http://www.cbr.ru/scripts/XML_daily.asp", nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Problem with CBR result: err = %#v", err)
		return
	}
	defer resp.Body.Close()
	var out bytes.Buffer
	io.Copy(&out, resp.Body)
	log.Printf("out = %s", out.String())

	/* Parse XML structure */
	var bankResponse BankResponse
	data := out.Bytes()
	newDecoder := xml.NewDecoder(bytes.NewReader(data))
	newDecoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}
	err1 := newDecoder.Decode(&bankResponse)
	if err1 != nil {
		log.Printf("error unmarshal XML wit curencies: err = %#v", err1)
	}
	log.Printf("result = %#v", bankResponse)

	/* Add USD an EUR */
	self.updateDate = bankResponse.Date
	self.convertRates = nil
	for _, valute := range bankResponse.Valutes {
		if valute.CharCode == "USD" {
			log.Printf("USD = %#v", valute)
			//
			newValue := ParseCurrency(valute.Value)
			//
			newConvertRate := ConvertRate{}
			newConvertRate.currency = '$'
			newConvertRate.convertRate = newValue
			self.convertRates = append(self.convertRates, newConvertRate)

		} else if valute.CharCode == "EUR" {
			log.Printf("EUR = %#v", valute)
			//
			newValue := ParseCurrency(valute.Value)
			//
			newConvertRate := ConvertRate{}
			newConvertRate.currency = '€'
			newConvertRate.convertRate = newValue
			self.convertRates = append(self.convertRates, newConvertRate)
		}
	}

	/* Debug */
	log.Printf("self.convertRates = %#v", self.convertRates)

}

func ParseCurrency(value string) float64 {
	log.Printf("Parse value: value = %#v", value)
	// "73,7330"
	newValue := strings.Replace(value, ",", ".", 1)
	log.Printf("ISO number value: %#v", newValue)
	res, err1 := strconv.ParseFloat(newValue, 64)
	if err1 != nil {
		log.Printf("Currency convert error: err = %#v", err1)
	}
	return res
}

func (self *BankWidget) Render() {

	for idx, c := range self.convertRates {
		currencyWidget := widgets.NewTextWidget("", 21)
		currencyWidget.SetText(fmt.Sprintf("%c - %.02f", c.currency, c.convertRate))
		currencyWidget.SetColor(0, 0, 0, 0)
		currencyWidget.SetRect(100, 100+idx*20, 100, 100)
		currencyWidget.Render()
		currencyWidget.Destroy()
	}

}
