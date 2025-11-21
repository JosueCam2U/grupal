package components

type Battery struct {
	Level    float64
	Capacity string
}

func NewBattery(level float64) *Battery {
	return &Battery{
		Level:    level,
		Capacity: "100kWh",
	}
}

func (b *Battery) GetLevel() float64 {
	return b.Level
}
