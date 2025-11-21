package components

type MissionData struct {
	Type        string
	Description string
	Completed   bool
}

func NewMissionData(missionType, description string) *MissionData {
	return &MissionData{
		Type:        missionType,
		Description: description,
		Completed:   false,
	}
}

func (m *MissionData) Complete() {
	m.Completed = true
}
