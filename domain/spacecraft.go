package domain

type SpacecraftStatus string

const (
	// SpacecraftStatusOperational is the status of a spacecraft that is operational and available to use.
	SpacecraftStatusOperational SpacecraftStatus = "operational"
	// SpacecraftStatusMaintenance is the status of a spacecraft that is under maintenance.
	SpacecraftStatusMaintenance SpacecraftStatus = "maintenance"
	// SpacecraftStatusDecomission is the status of a spacecraft that is decomissioned/destroyed.
	SpacecraftStatusDecomission SpacecraftStatus = "decomission"
	// SpacecraftStatusDamaged is the status of a spacecraft that is damaged and not available to use and scheduled for repair.
	SpacecraftStatusDamaged SpacecraftStatus = "damaged"
	// SpacecraftStatusUnavailable is the status of a spacecraft that is unavailable to use.
	SpacecraftStatusUnavailable SpacecraftStatus = "unavailable"
	// SpacecraftStatusUnknownis the fallback status
	SpacecraftStatusUnknown SpacecraftStatus = "unknown"
)

func StringToSpacecraftStatus(s string) SpacecraftStatus {
	switch s {
	case string(SpacecraftStatusOperational):
		return SpacecraftStatusOperational
	case string(SpacecraftStatusMaintenance):
		return SpacecraftStatusMaintenance
	case string(SpacecraftStatusDecomission):
		return SpacecraftStatusDecomission
	case string(SpacecraftStatusDamaged):
		return SpacecraftStatusDamaged
	case string(SpacecraftStatusUnavailable):
		return SpacecraftStatusUnavailable
	default:
		return SpacecraftStatusUnknown
	}
}

// Spacecraft holds the application data of a spacecraft.
// Handles the mapping of the data from the database to the application.
// Defines the JSON schema of the API.
type Spacecraft struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Class    string     `json:"class"`
	Crew     uint32     `json:"crew"`            // uint because we don't expect negative values
	Image    string     `json:"image,omitempty"` // optional IMO
	Value    float64    `json:"value"`
	Status   string     `json:"status"`
	Armament []Armament `json:"armament,omitempty"`
}

// Armament holds the application data of a Armament.
// Handles the mapping of the data from the database to the application.
// Defines the JSON schema of the API.
type Armament struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Quantity string `json:"qty"`
}

// SpacecraftFilters holds the filters that can be applied to the GetSpacecrafts method. all are optional
type SpacecraftFilters struct {
	Name   string           `json:"name,omitempty"`
	Class  string           `json:"class,omitempty"`
	Status SpacecraftStatus `json:"status,omitempty"`
}
