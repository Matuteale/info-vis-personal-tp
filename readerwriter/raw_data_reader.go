package readerwriter

import (
	"encoding/json"
	"io/ioutil"
	"projects/info-vis-personal-tp/model"
)

type RawData struct {
	Locations []model.Location `json:"locations"`
}

func ReadRawData(fileName string) (*[]model.Location, error) {
	var rawData RawData
	rawDataBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawDataBytes, &rawData)
	if err != nil {
		return nil, err
	}
	return &rawData.Locations, nil
}
