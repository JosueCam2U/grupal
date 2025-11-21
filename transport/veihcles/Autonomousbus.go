package veihcles

import (
	"fmt"
)

type AutonomousBus struct {
	ID             string
	CurrentMission string
	IsRunning      bool
	PassengerCount int
	Route          string
}

func NewAutonomousBus(id string) *AutonomousBus {
	return &AutonomousBus{
		ID:             id,
		Route:          "Ruta Centro",
		PassengerCount: 0,
	}
}

func (b *AutonomousBus) Start() {
	if !b.IsRunning {
		b.IsRunning = true
		b.PassengerCount = 30
		fmt.Printf("Autobús %s inició viaje con %d pasajeros\n", b.ID, b.PassengerCount)
	}
}

func (b *AutonomousBus) Stop() {
	if b.IsRunning {
		b.IsRunning = false
		fmt.Printf("Autobús %s se detuvo. Pasajeros desembarcaron\n", b.ID)
		b.PassengerCount = 0
	}
}

func (b *AutonomousBus) ReportStatus() string {
	status := "Detenido"
	if b.IsRunning {
		status = "En ruta"
	}
	return fmt.Sprintf("[Bus %s] %s - %s | Pasajeros: %d", b.ID, status, b.CurrentMission, b.PassengerCount)
}

func (b *AutonomousBus) AssignMission(mission string) {
	b.CurrentMission = mission
}
