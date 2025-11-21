package main

import (
	"fmt"
	"transport/control"
	"transport/veihcles"
)

func main() {
	// Crear centro de control
	center := control.NewControlCenter()

	// Crear vehículos
	bus := veihcles.NewAutonomousBus("BUS-001")
	drone := veihcles.NewDeliveryDrone("DRONE-001")

	// Registrar vehículos
	center.RegisterVehicle(bus)
	center.RegisterVehicle(drone)

	// Asignar misiones
	bus.AssignMission("Transportar 30 pasajeros al centro")
	drone.AssignMission("Entregar paquete médico urgente")

	fmt.Println("=== INICIANDO SIMULACIÓN ===")

	// Demostrar polimorfismo
	center.StartAllVehicles()

	fmt.Println("\n=== REPORTE CONSOLIDADO ===")
	center.ReportAllStatus()

	fmt.Println("\n=== DETENIENDO VEHÍCULOS ===")
	center.StopAllVehicles()

	fmt.Println("\n=== ESTADO FINAL ===")
	center.ReportAllStatus()
}
