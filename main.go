package main

import (
	"projects/info-vis-personal-tp/data_processing"
	"projects/info-vis-personal-tp/readerwriter"
)

func main() {
	println("Reading full location history...")
	locations, err := readerwriter.ReadRawData("raw_data/full_location_history.json")
	if err != nil {
		println("Error reading full location history. Error: %v", err)
		return
	}
	processedData, err := data_processing.ProcessLocations(locations)
	if err != nil {
		println("Error while processing location history data. Error: %v", err)
		return
	}
	err = readerwriter.WriteProcessedData("processed_data/processed_location_history_data.json", processedData)
	if err != nil {
		println("Error writing processed location history data. Error: %v", err)
		return
	}
}
