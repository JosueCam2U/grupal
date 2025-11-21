package veihcles

import (
	"fmt"
)

type DeliveryDrone struct {
	ID             string
	CurrentMission string
	IsRunning      bool
	IsFlying       bool
	PackageWeight  float64
}

func NewDeliveryDrone(id string) *DeliveryDrone {
	return &DeliveryDrone{
		ID:            id,
		PackageWeight: 2.5,
	}
}

func (d *DeliveryDrone) Start() {
	if !d.IsRunning {
		d.IsRunning = true
		d.IsFlying = true
		fmt.Printf("Dron %s despegó con paquete de %.1f kg\n", d.ID, d.PackageWeight)
	}
}

func (d *DeliveryDrone) Stop() {
	if d.IsRunning {
		d.IsRunning = false
		d.IsFlying = false
		fmt.Printf("Dron %s aterrizó y entregó paquete\n", d.ID)
		d.PackageWeight = 0
	}
}

func (d *DeliveryDrone) ReportStatus() string {
	status := "En tierra"
	if d.IsFlying {
		status = "Volando"
	}
	return fmt.Sprintf("[Drone %s] %s - %s | Peso: %.1fkg", d.ID, status, d.CurrentMission, d.PackageWeight)
}

func (d *DeliveryDrone) AssignMission(mission string) {
	d.CurrentMission = mission
}
