package utils

import (
	"encoding/json"
	"os"
	"sync"
)

var mutex sync.Mutex

func ReadJson[T any](filePath string, v *T) error {
	mutex.Lock()
	defer mutex.Unlock()
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func WriteJson[T any](filePath string, v T) error {
	mutex.Lock()
	defer mutex.Unlock()
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
