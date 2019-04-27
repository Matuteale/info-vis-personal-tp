package dataprocessing

import (
	"math"
	"projects/info-vis-personal-tp/model"
	"time"
)

func LocationsToGeoJSON(locations []model.Location, year int) (model.LocationGeoJSON, error) {
	geoJSON := model.LocationGeoJSON{
		Type: "FeatureCollection",
	}
	if len(locations) == 0 {
		return geoJSON, nil
	}
	var features []model.Feature
	for _, location := range locations {
		features = append(features, model.Feature{
			Type: "Feature",
			Geometry: model.Geometry{
				Type:        "Point",
				Coordinates: []float64{convertToGPSFormat(location.LongitudeE7), convertToGPSFormat(location.LatitudeE7)},
			},
		})
	}
	geoJSON.Features = features
	return geoJSON, nil
}

func LocationsToStatistics(locations []model.Location) (model.LocationStatistics, error) {
	statistics := model.LocationStatistics{
		OnFootTime:        make(map[int]map[int]map[int]int64),
		OnVehicleTime:     make(map[int]map[int]map[int]int64),
		OnFootDistance:    make(map[int]map[int]map[int]float64),
		OnVehicleDistance: make(map[int]map[int]map[int]float64),
	}
	if len(locations) == 0 {
		return statistics, nil
	}
	var startIndex int
	for !isValidLocation(locations[startIndex]) {
		startIndex++
	}
	lastLocation := locations[startIndex]
	var lastLocationDate = time.Unix(0, lastLocation.TimestampMs*int64(time.Millisecond))
	var auxOnFootDistance float64
	var auxOnVehicleDistance float64
	var auxOnFootTime int64
	var auxOnVehicleTime int64
	var currentLocationDate time.Time
	var auxYear int
	var auxMonth int
	var auxDay int
	for i := startIndex + 1; i < len(locations); i++ {
		currentLocation := locations[i]
		if !isValidLocation(currentLocation) {
			continue
		}
		currentLocationDate = time.Unix(0, currentLocation.TimestampMs*int64(time.Millisecond))
		if isSameTrackDate(lastLocationDate, currentLocationDate) {
			if isOnVehicle(lastLocation, currentLocation) {
				auxOnVehicleDistance += getDistanceBetween(lastLocation, currentLocation)
				auxOnVehicleTime += (currentLocation.TimestampMs - lastLocation.TimestampMs)
			} else if isOnFoot(lastLocation, currentLocation) {
				auxOnFootDistance += getDistanceBetween(lastLocation, currentLocation)
				auxOnFootTime += (currentLocation.TimestampMs - lastLocation.TimestampMs)
			}
		} else {
			auxYear = lastLocationDate.Year()
			auxMonth = int(lastLocationDate.Month())
			auxDay = lastLocationDate.Day()
			if auxOnVehicleTime > 0 {
				if statistics.OnVehicleDistance[auxYear] == nil {
					statistics.OnVehicleDistance[auxYear] = make(map[int]map[int]float64)
				}
				if statistics.OnVehicleDistance[auxYear][auxMonth] == nil {
					statistics.OnVehicleDistance[auxYear][auxMonth] = make(map[int]float64)
				}
				if statistics.OnVehicleTime[auxYear] == nil {
					statistics.OnVehicleTime[auxYear] = make(map[int]map[int]int64)
				}
				if statistics.OnVehicleTime[auxYear][auxMonth] == nil {
					statistics.OnVehicleTime[auxYear][auxMonth] = make(map[int]int64)
				}
				statistics.OnVehicleDistance[auxYear][auxMonth][auxDay] += auxOnVehicleDistance
				statistics.OnVehicleTime[auxYear][auxMonth][auxDay] += (auxOnVehicleTime / 1000)
			}
			if auxOnFootTime > 0 {
				if statistics.OnFootDistance[auxYear] == nil {
					statistics.OnFootDistance[auxYear] = make(map[int]map[int]float64)
				}
				if statistics.OnFootDistance[auxYear][auxMonth] == nil {
					statistics.OnFootDistance[auxYear][auxMonth] = make(map[int]float64)
				}
				if statistics.OnFootTime[auxYear] == nil {
					statistics.OnFootTime[auxYear] = make(map[int]map[int]int64)
				}
				if statistics.OnFootTime[auxYear][auxMonth] == nil {
					statistics.OnFootTime[auxYear][auxMonth] = make(map[int]int64)
				}
				statistics.OnFootDistance[auxYear][auxMonth][auxDay] += auxOnFootDistance
				statistics.OnFootTime[auxYear][auxMonth][auxDay] += (auxOnFootTime / 1000)
			}
			auxOnVehicleDistance = 0
			auxOnVehicleTime = 0
			auxOnFootDistance = 0
			auxOnFootTime = 0
		}
		lastLocationDate = currentLocationDate
		lastLocation = currentLocation
	}
	return statistics, nil
}

func isSameTrackDate(lastLocationDate time.Time, currentLocationDate time.Time) bool {
	return lastLocationDate.Year() == currentLocationDate.Year() && lastLocationDate.Month() == currentLocationDate.Month() && lastLocationDate.Day() == currentLocationDate.Day() && ((currentLocationDate.Hour()*60+currentLocationDate.Minute())-(lastLocationDate.Hour()*60+lastLocationDate.Minute())) <= 10
}

func convertToGPSFormat(axis float64) float64 {
	axis /= 1E7
	if axis > 180 {
		axis -= (math.MaxUint32 / 1E7)
	}
	return axis
}

func getDistanceBetween(location model.Location, other model.Location) float64 {
	var lon1 = convertToGPSFormat(location.LongitudeE7) * math.Pi / 180
	var lat1 = convertToGPSFormat(location.LatitudeE7) * math.Pi / 180
	var lon2 = convertToGPSFormat(other.LongitudeE7) * math.Pi / 180
	var lat2 = convertToGPSFormat(other.LatitudeE7) * math.Pi / 180
	var a = math.Pow(math.Sin((lat2-lat1)/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((lon2-lon1)/2), 2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return 6371E3 * c
}

func isValidLocation(location model.Location) bool {
	return location.Accuracy < 1000
}

func isOnFoot(lastLocation model.Location, currentLocation model.Location) bool {
	dist := getDistanceBetween(lastLocation, currentLocation)
	distTime := (currentLocation.TimestampMs - lastLocation.TimestampMs)
	return dist/float64((distTime/1000)) <= 3
}

func isOnVehicle(lastLocation model.Location, currentLocation model.Location) bool {
	dist := getDistanceBetween(lastLocation, currentLocation)
	distTime := (currentLocation.TimestampMs - lastLocation.TimestampMs)
	return dist/float64((distTime/1000)) > 3 && dist/float64((distTime/1000)) < 45
}
