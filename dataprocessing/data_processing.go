package dataprocessing

import (
	"math"
	"projects/info-vis-personal-tp/model"
	"time"
)

func LocationsToStatistics(locations []model.Location) (model.LocationStatistics, error) {
	statistics := model.LocationStatistics{
		OnFootTime:        make(map[int]map[int]int64),
		OnVehicleTime:     make(map[int]map[int]int64),
		OnFootDistance:    make(map[int]map[int]int64),
		OnVehicleDistance: make(map[int]map[int]int64),
	}
	return statistics, nil
}

func LocationsToGeoJSON(locations []model.Location) (model.LocationGeoJSON, error) {
	var features []model.Feature
	var longitude float64
	var latitude float64
	for _, location := range locations {
		if time.Unix(0, location.TimestampMs*int64(time.Millisecond)).Year() == 2018 {
			longitude = location.LongitudeE7 / 1E7
			if longitude > 180 {
				longitude = longitude - (math.MaxUint32 / 1E7)
			}
			latitude = location.LatitudeE7 / 1E7
			if latitude > 180 {
				latitude = latitude - (math.MaxUint32 / 1E7)
			}
			features = append(features, model.Feature{
				Type: "Feature",
				Geometry: model.Geometry{
					Type:        "Point",
					Coordinates: []float64{longitude, latitude},
				},
			})
		}
	}
	geoJSON := model.LocationGeoJSON{
		Type:     "FeatureCollection",
		Features: features,
	}
	return geoJSON, nil
}
