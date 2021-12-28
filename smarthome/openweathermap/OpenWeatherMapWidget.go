package openweathermap

import (
	"encoding/json"
	"fmt"
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type OpenWeatherMapWidget struct {
	widgets.BaseWidget
	error    string /* Error message          */
	apiToken string /* API token              */
	out      string /* Temp out               */
}

type OpenWeatherMapErrorResponse struct {
	Cod     int    `json:"cod"`
	Message string `json:"message"`
}

type OpenWeatherMapResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Snow struct {
		H float64 `json:"1h"`
	} `json:"snow"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func NewOpenWeatherMapWidget() *OpenWeatherMapWidget {
	newOpenWeatherMapWidget := new(OpenWeatherMapWidget)
	newOpenWeatherMapWidget.recoverToken()
	go func() {
		for {
			newOpenWeatherMapWidget.updateData()
			/* Wait 10 min */
			time.Sleep(30 * time.Minute)
		}
	}()
	return newOpenWeatherMapWidget
}

func (self *OpenWeatherMapWidget) recoverToken() {

	/* Step 0. Prepare reding user home directory */
	userDirName, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Todoist error API token reading: err = %#v", err)
	}

	/* Step 1. Prepare Todoist token path */
	newTokenPath := path.Join(userDirName, ".openweathermap")
	log.Printf("openweathermap token path: %#v", newTokenPath)

	/* Step 2. Reading content with token */
	content, err := ioutil.ReadFile(newTokenPath)
	if err != nil {
		log.Printf("openweathermap error API token reading: err = %#v", err)
	}
	self.apiToken = strings.Trim(string(content), " \r\n\t")

}

func (self *OpenWeatherMapWidget) updateData() error {

	mainNetworkManager := ctx.GetNetworkManager()
	//mainGeolocationManager :=
	//ourGelocation := mainGeolocationManager.GetGelocation()
	lat := 59.998335
	lon := 30.363723
	ourLang := "RU"

	/* Step 1. Download response */
	req, err1 := mainNetworkManager.MakeRequest("OpenWeatherMapWidget", "GET", "http://api.openweathermap.org/data/2.5/weather", 15)
	if err1 != nil {
		return err1
	}
	req.AddQueryParam("lat", fmt.Sprintf("%f", lat))
	req.AddQueryParam("lon", fmt.Sprintf("%f", lon))
	req.AddQueryParam("units", "metric")
	req.AddQueryParam("lang", ourLang)
	req.AddQueryParam("appid", self.apiToken)

	resp, err2 := mainNetworkManager.Perform(req)
	if err2 != nil {
		return err2
	}
	newContent := resp.Bytes()

	/* Step 2. Process error response */
	var weatherErrorResponse OpenWeatherMapErrorResponse
	err3 := json.Unmarshal(newContent, &weatherErrorResponse)
	if err3 != nil {
		return err3
	}
	self.error = weatherErrorResponse.Message

	/* Step 2. Parse OpenWeatherMap response */
	var weatherResponse OpenWeatherMapResponse
	err4 := json.Unmarshal(newContent, &weatherResponse)
	if err4 != nil {
		return err4
	}
	log.Printf("weatherResponse = %#v", weatherResponse)
	self.out = fmt.Sprintf("Temp = %.1f (%s)", weatherResponse.Main.Temp, weatherResponse.Name)

	return nil
}

func (self *OpenWeatherMapWidget) ProcessEvent(e *evt.Event) {
}

func (self *OpenWeatherMapWidget) Render() {

	self.BaseWidget.Render()

	if self.error != "" {
		errorMessage := widgets.NewTextWidget("", 16)
		errorMessage.SetColor(255, 0, 0, 128)
		errorMessage.SetRect(self.X, self.Y, self.Width, self.Height)
		errorMessage.SetText(self.error)
		errorMessage.Render()
		errorMessage.Destroy()
	}

	if self.out != "" {
		outMessage := widgets.NewTextWidget("", 16)
		outMessage.SetColor(0, 0, 0, 0)
		outMessage.SetRect(self.X, self.Y, self.Width, self.Height)
		outMessage.SetText(self.out)
		outMessage.Render()
		outMessage.Destroy()
	}

}
