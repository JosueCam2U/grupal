package veihcles

type AutonomousVehicle interface {
	Start()
	Stop()
	ReportStatus() string
	AssignMission(mission string)
}
