package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"projects/info-vis-personal-tp/model"
)

type RawData struct {
	Locations []model.Location `json:"locations"`
}

func ReadData(fileName string) *[]model.Location {
	var rawData RawData
	rawDataBytes, err := ioutil.ReadFile(fmt.Sprintf("%s", fileName))
	if err != nil {
		panic(fmt.Sprintf("error reading data file. err: %v", err))
	}
	err = json.Unmarshal(rawDataBytes, &rawData)
	if err != nil {
		panic(fmt.Sprintf("error parsing data file. err: %v", err))
	}
	return &rawData.Locations
}
