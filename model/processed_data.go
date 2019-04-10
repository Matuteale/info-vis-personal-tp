package model

type ProcessedData struct {
	OnFootTime        map[int]map[int]int64     `json:"onFootTimeSec"`    // By year and month
	OnVehicleTime     map[int]map[int]int64     `json:"onVehicleTimeSec"` // By year and month
	OnFootSegments    map[int]map[int][]Segment `json:"onFootSegments"`
	OnVehicleSegments map[int]map[int][]Segment `json:"OnVehicleSegments"`
}

type Segment struct {
	From SegmentPoint  `json:"onFootSegments"`
	To   *SegmentPoint `json:"onFootSegments,omitempty"`
}

type SegmentPoint struct {
	TimestampMs int64  `json:"timestampMs"`
	LatitudeE7  int64  `json:"latitudeE7"`
	LongitudeE7 int64  `json:"longitudeE7"`
	Altitude    *int64 `json:"altitude,omitempty"`
}
