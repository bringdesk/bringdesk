package gismeteo

type GismeteoWidget struct {
}

func NewGismeteoWidget() *GismeteoWidget {
	return new(GismeteoWidget)
}

func updateData() {
	// curl -H 'X-Gismeteo-Token: 56b30cb255.3443075' 'https://api.gismeteo.net/v2/weather/current/4368/'

}

func (self *GismeteoWidget) Render() {

}
