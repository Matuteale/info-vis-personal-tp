package model

type Location struct {
	TimestampMs int64   `json:"timestampMs"`
	LatitudeE7  float64 `json:"latitudeE7"`
	LongitudeE7 float64 `json:"longitudeE7"`
	Altitude    *int64  `json:"altitude,omitempty"`
	Accuracy    int     `json:"accuracy"`
}
