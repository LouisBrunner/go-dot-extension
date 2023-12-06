package gencommon

import (
	"encoding/json"
	"os"
)

func ReadAPI(path string) (*ExtensionAPI, error) {
	inFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var api ExtensionAPI
	err = json.Unmarshal(inFile, &api)
	if err != nil {
		return nil, err
	}
	return &api, nil
}
