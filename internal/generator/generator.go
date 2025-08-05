package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-fake/internal/schema"
	"go-fake/pkg/csv"
	"go-fake/pkg/faker"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var ErrInvalidSchema = errors.New("invalid schema type")

// OutputFormat represents the desired output format
type OutputFormat int

const (
	FormatCSV OutputFormat = iota
	FormatJSON
)

// GenerateDataFiles generates fake data and creates separate files for each table.
// The format is determined by the outputFormat parameter.
func GenerateDataFiles(schemaInterface interface{}, numRows int, outputPath string, format OutputFormat) ([]string, error) {
	// Convert the schema to the expected type
	s, ok := schemaInterface.(schema.Schema)
	if !ok {
		return nil, ErrInvalidSchema
	}

	var generatedFiles []string

	// Handle multi-table schemas (SQL)
	if len(s.Tables) > 0 {
		for _, table := range s.Tables {
			var filename string
			var err error
			
			if format == FormatJSON {
				data := generateTableDataAsJSON(table.Fields, numRows, table.Name)
				filename = getOutputFilename(outputPath, table.Name, ".json")
				err = writeJSONFile(filename, data)
			} else {
				data := generateTableData(table.Fields, numRows)
				filename = getOutputFilename(outputPath, table.Name, ".csv")
				err = csv.WriteCSV(filename, data)
			}
			
			if err != nil {
				return generatedFiles, fmt.Errorf("error writing %s: %v", filename, err)
			}
			
			generatedFiles = append(generatedFiles, filename)
		}
	} else if len(s.Fields) > 0 {
		// Handle single-table schemas (JSON or simple format)
		var filename string
		var err error
		
		if format == FormatJSON {
			data := generateTableDataAsJSON(s.Fields, numRows, "data")
			filename = outputPath
			if filename == "" || strings.HasSuffix(filename, ".csv") {
				filename = strings.TrimSuffix(filename, ".csv") + ".json"
				if filename == ".json" {
					filename = "output.json"
				}
			}
			err = writeJSONFile(filename, data)
		} else {
			data := generateTableData(s.Fields, numRows)
			filename = outputPath
			if filename == "" {
				filename = "output.csv"
			}
			err = csv.WriteCSV(filename, data)
		}
		
		if err != nil {
			return generatedFiles, fmt.Errorf("error writing %s: %v", filename, err)
		}
		
		generatedFiles = append(generatedFiles, filename)
	} else {
		return nil, errors.New("schema contains no tables or fields")
	}

	return generatedFiles, nil
}

// GenerateData generates fake data based on the provided schema (backward compatibility).
func GenerateData(schemaInterface interface{}, numRows int) ([][]string, error) {
	// Convert the schema to the expected type
	s, ok := schemaInterface.(schema.Schema)
	if !ok {
		return nil, ErrInvalidSchema
	}

	// For backward compatibility, use the first table or the fields
	var fields []schema.Field
	if len(s.Tables) > 0 {
		fields = s.Tables[0].Fields
	} else {
		fields = s.Fields
	}

	return generateTableData(fields, numRows), nil
}

// generateTableData generates fake data for a specific table's fields
func generateTableData(fields []schema.Field, numRows int) [][]string {
	var data [][]string
	headers := make([]string, len(fields))
	for i, field := range fields {
		headers[i] = field.Name
	}
	data = append(data, headers)

	// Generate fake data for each row
	for i := 0; i < numRows; i++ {
		row := make([]string, len(fields))
		for j, field := range fields {
			row[j] = generateFakeValue(field)
		}
		data = append(data, row)
	}

	return data
}

// generateTableDataAsJSON generates fake data as JSON objects
func generateTableDataAsJSON(fields []schema.Field, numRows int, tableName string) map[string]interface{} {
	var records []map[string]interface{}

	// Generate fake data for each row
	for i := 0; i < numRows; i++ {
		record := make(map[string]interface{})
		for _, field := range fields {
			value := generateFakeValue(field)
			// Try to convert numeric strings to actual numbers for JSON
			if field.Type == "int" || field.Type == "integer" || field.Type == "serial" {
				if intVal, err := strconv.Atoi(value); err == nil {
					record[field.Name] = intVal
				} else {
					record[field.Name] = value
				}
			} else if field.Type == "float" || field.Type == "decimal" || field.Type == "numeric" || field.Type == "price" {
				if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
					record[field.Name] = floatVal
				} else {
					record[field.Name] = value
				}
			} else if field.Type == "bool" || field.Type == "boolean" {
				record[field.Name] = value == "true"
			} else {
				record[field.Name] = value
			}
		}
		records = append(records, record)
	}

	return map[string]interface{}{
		tableName: records,
	}
}

// getOutputFilename generates the output filename based on table name and extension
func getOutputFilename(outputPath, tableName, ext string) string {
	if outputPath == "" {
		return tableName + ext
	}
	
	// Replace the base filename with table-specific name
	dir := filepath.Dir(outputPath)
	currentExt := filepath.Ext(outputPath)
	base := strings.TrimSuffix(filepath.Base(outputPath), currentExt)
	return filepath.Join(dir, fmt.Sprintf("%s_%s%s", base, tableName, ext))
}

// writeJSONFile writes JSON data to a file
func writeJSONFile(filename string, data map[string]interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// generateFakeValue generates a fake value based on the field type.
func generateFakeValue(field schema.Field) string {
	switch field.Type {
	case "string", "varchar", "text":
		return faker.GenerateName()
	case "email":
		return faker.GenerateEmail()
	case "int", "integer", "serial":
		return strconv.Itoa(rand.IntN(1000) + 1)
	case "float", "decimal", "numeric":
		return faker.GenerateFloat()
	case "bool", "boolean":
		return faker.GenerateBool()
	case "date":
		return faker.GenerateDate()
	case "datetime", "timestamp":
		return faker.GenerateDateTime()
	case "phone":
		return faker.GeneratePhone()
	case "company":
		return faker.GenerateCompany()
	case "address":
		return faker.GenerateAddress()
	case "city":
		return faker.GenerateCity()
	case "state":
		return faker.GenerateState()
	case "zipcode", "zip":
		return faker.GenerateZipCode()
	case "uuid":
		return faker.GenerateUUID()
	case "price":
		return faker.GeneratePrice()
	case "firstname":
		return faker.GenerateFirstName()
	case "lastname":
		return faker.GenerateLastName()
	default:
		return faker.GenerateName() // Default to name for unknown types
	}
}