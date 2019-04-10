package readerwriter

import (
	"encoding/json"
	"io/ioutil"
	"projects/info-vis-personal-tp/model"
)

func WriteProcessedData(fileName string, processedData *model.ProcessedData) error {
	asBytes, err := json.Marshal(processedData)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, asBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
