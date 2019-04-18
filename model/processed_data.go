package model

type LocationStatistics struct {
	OnFootTime        map[int]map[int]int64 `json:"onFootTimeSec,omitempty"`     // By year and month
	OnVehicleTime     map[int]map[int]int64 `json:"onVehicleTimeSec,omitempty"`  // By year and month
	OnFootDistance    map[int]map[int]int64 `json:"OnFootDistance,omitempty"`    // By year and month
	OnVehicleDistance map[int]map[int]int64 `json:"onVehicleDistance,omitempty"` // By year and month
}

type LocationGeoJSON struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type     string   `json:"type"`
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
