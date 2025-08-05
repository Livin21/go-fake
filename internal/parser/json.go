package parser

import (
	"encoding/json"
	"go-fake/internal/schema"
	"io/ioutil"
	"os"
)

// ParseJSONSchema reads a JSON file and returns a structured schema.
func ParseJSONSchema(filePath string) (schema.Schema, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return schema.Schema{}, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return schema.Schema{}, err
	}

	var s schema.Schema
	if err := json.Unmarshal(bytes, &s); err != nil {
		return schema.Schema{}, err
	}

	return s, nil
}