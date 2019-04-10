package model

type ProcessedData struct {
	TimestampMs int64 `json:"timestampMs"`
	LatitudeE7  int64 `json:"latitudeE7"`
}
