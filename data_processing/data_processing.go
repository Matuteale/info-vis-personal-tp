package data_processing

import (
	"projects/info-vis-personal-tp/model"
)

func ProcessLocations(locations *[]model.Location) (*model.ProcessedData, error) {
	proccesedData := model.ProcessedData{
		OnFootTime:        make(map[int]map[int]int64),
		OnVehicleTime:     make(map[int]map[int]int64),
		OnFootSegments:    make(map[int]map[int][]model.Segment),
		OnVehicleSegments: make(map[int]map[int][]model.Segment),
	}
	return &proccesedData, nil
}
