package control

import "transport/veihcles"

type ControlCenter struct {
	vehicles []veihcles.AutonomousVehicle
}

func NewControlCenter() *ControlCenter {
	return &ControlCenter{
		vehicles: make([]veihcles.AutonomousVehicle, 0),
	}
}

func (cc *ControlCenter) RegisterVehicle(vehicle veihcles.AutonomousVehicle) {
	cc.vehicles = append(cc.vehicles, vehicle)
}

func (cc *ControlCenter) StartAllVehicles() {
	for _, vehicle := range cc.vehicles {
		vehicle.Start()
	}
}

func (cc *ControlCenter) StopAllVehicles() {
	for _, vehicle := range cc.vehicles {
		vehicle.Stop()
	}
}

func (cc *ControlCenter) ReportAllStatus() {
	for _, vehicle := range cc.vehicles {
		println(vehicle.ReportStatus())
	}
}

// Demostración de polimorfismo
func (cc *ControlCenter) OperateSingleVehicle(vehicle veihcles.AutonomousVehicle) {
	println("\n--- Operando vehículo individual (Polimorfismo) ---")
	vehicle.AssignMission("Misión especial por polimorfismo")
	vehicle.Start()
	println(vehicle.ReportStatus())
	vehicle.Stop()
}
