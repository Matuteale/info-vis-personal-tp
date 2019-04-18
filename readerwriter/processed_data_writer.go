package readerwriter

import (
	"encoding/json"
	"io/ioutil"
)

func WriteProcessedData(fileName string, processedData interface{}) error {
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
