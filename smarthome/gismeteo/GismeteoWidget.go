package gismeteo

import (
	"bytes"
	"encoding/json"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"io"
	"log"
	"net/http"
	"time"
)

type GismeteoWidget struct {
	widgets.BaseWidget
	error string /* Код ошибки             */
}

type GismeteoErrorResponse struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Response struct {
	} `json:"response"`
}

type GismeteoResponse struct {
	Kind string /* Тип погодных данных: Obs - наблюдение, Frc - Прогноз */
	Date struct {
		UTC   string /* По стандарту UTC                                     */
		Unix  int    /* В формате UNIX по стандарту UTC                      */
		Local string /* По локальному времени географического объекта        */
	}
	Temperature struct { /* Температура                                          */
		Air struct { /* Воздух                                               */
			C float32 /* В градусах Цельсия                                   */
		}
		Comfort struct {
			C float32 /* В градусах Цельсия                                   */
		}
		Water struct {
			C float32 /* В градусах Цельсия                                   */
		}
	}
	Description struct { /* Описание погоды                                      */
		Full string /* Полное описание                                      */
	}
	Humidity struct { /* Влажность                                            */
		Percent int /* В процентах                                          */
	}
	Pressure struct {
		/* mm_hg_atm */
	}
	Cloudiness struct { /* Облачность                                           */
		Percent int
		Type    int /* По шкале от 0 до 3:                                  */
		/*               0 - Ясно,                              */
		/*               1 - Малооблачно,                       */
		/*               2 - Облачно,                           */
		/*               3 - Пасмурно,                          */
		/*             101 - Переменная облачность              */
	}
	Storm struct { /* Гроза                                                */
		Prediction bool /* Вероятность грозы                                    */
	}
	Precipitation struct { /* Осадки                                               */
		Type int /* Тип осадков:                                         */
		/*            0	Нет осадков                           */
		/*            1	Дождь                                 */
		/*            2	Снег                                  */
		/*            3	Смешанные осадки                      */
		Amount    *float32 /* Количество осадков в мм.                             */
		Intensity int      /* Интенсивность осадков                                */
	}
	Phenomenon int      /* Код погодного явления                                */
	Icon       string   /* Иконка погоды                                        */
	Gm         int      /* Геомагнитное поле                                    */
	Wind       struct { /* Ветер                                                */
		Direction struct {
			Degree int /* В градусах                                           */
		}
		Speed struct {
			/* m_s */ /* В метрах в секунду                                   */
		}
	}
}

func NewGismeteoWidget() *GismeteoWidget {
	newGismeteoWidget := new(GismeteoWidget)
	go func() {
		for {
			newGismeteoWidget.updateData()
			/* Wait 10 min */
			time.Sleep(30 * time.Minute)
		}
	}()
	newGismeteoWidget.updateData()
	return newGismeteoWidget
}

func (self *GismeteoWidget) updateData() {

	/* Step 1. Download response */
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	req, _ := http.NewRequest("GET", "https://api.gismeteo.net/v2/weather/current/?latitude=54.35&longitude=52.52", nil)
	req.Header.Add("X-Gismeteo-Token", "56b30cb255.3443075")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Problem with Gismeto result: err = %#v", err)
		return
	}
	defer resp.Body.Close()
	var out bytes.Buffer
	io.Copy(&out, resp.Body)
	log.Printf("out = %s", out.String())

	/* Step 2. Parse Gismeteo response */
	var errorResponse GismeteoErrorResponse
	err1 := json.Unmarshal(out.Bytes(), &errorResponse)
	if err1 != nil {
		log.Printf("Parse error: err = %#v", err1)
	} else {
		self.error = errorResponse.Meta.Message
	}

}

func (self *GismeteoWidget) ProcessEvent(e *evt.Event) {
}

func (self *GismeteoWidget) Render() {

	self.BaseWidget.Render()

	if self.error != "" {
		errorMessage := widgets.NewTextWidget("", 16)
		errorMessage.SetColor(255, 0, 0, 128)
		errorMessage.SetRect(self.X, self.Y, self.Width, self.Height)
		errorMessage.SetText(self.error)
		errorMessage.Render()
		errorMessage.Destroy()
	}

}
