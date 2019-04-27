package main

import (
	"log"
	"os"
	"projects/info-vis-personal-tp/dataprocessing"
	"projects/info-vis-personal-tp/readerwriter"
)

func main() {
	err := os.RemoveAll("processed_data/")
	if err != nil {
		log.Fatalf("Error removing files: %v", err)
	}
	err = os.MkdirAll("processed_data/", 0777)
	if err != nil {
		log.Fatalf("Error removing files: %v", err)
	}
	println("Reading raw full location history...")
	locations, err := readerwriter.ReadRawData("raw_data/full_location_history.json")
	if err != nil {
		println("Error reading full location history. Error: %v", err)
		return
	}
	println("Processing raw full location history statistics...")
	statistics, err := dataprocessing.LocationsToStatistics(locations)
	if err != nil {
		println("Error while processing location history data to statistics. Error: %v", err)
		return
	}
	println("Writing processed location history statistics...")
	err = readerwriter.WriteProcessedData("processed_data/statistics.json", statistics)
	if err != nil {
		println("Error writing processed location history statistics. Error: %v", err)
		return
	}
	println("Processing raw full location history geoJSON...")
	geoJSON, err := dataprocessing.LocationsToGeoJSON(locations, 2018)
	if err != nil {
		println("Error while processing location history data to statistics. Error: %v", err)
		return
	}
	println("Writing processed location history geoJSON...")
	err = readerwriter.WriteProcessedData("processed_data/locations_2014_2019.geojson", geoJSON)
	if err != nil {
		println("Error writing processed location history geoJSON. Error: %v", err)
		return
	}
	println("Success")
}
