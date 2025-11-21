package components

type GPS struct {
	Latitude  float64
	Longitude float64
}

func NewGPS() *GPS {
	return &GPS{
		Latitude:  -0.2298, // Quito
		Longitude: -78.5249,
	}
}

func (g *GPS) GetLocation() (float64, float64) {
	return g.Latitude, g.Longitude
}
