package util

type GeoLocationManager struct {
}

type Location struct {
	Latitude  float64
	Longitude float64
}

func NewGeoLocationManager() *GeoLocationManager {
	return new(GeoLocationManager)
}

func (self *GeoLocationManager) GetLocation() *Location {
	return nil
}
