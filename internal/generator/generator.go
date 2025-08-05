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

// RelationshipData stores generated data for relationship constraints
type RelationshipData struct {
	TableData map[string][]map[string]interface{} // table_name -> rows of data
	References map[string][]interface{} // table.field -> list of generated values
}

// GenerateDataFiles generates fake data and creates separate files for each table.
// The format is determined by the outputFormat parameter.
func GenerateDataFiles(schemaInterface interface{}, numRows int, outputPath string, format OutputFormat) ([]string, error) {
	// Convert the schema to the expected type
	s, ok := schemaInterface.(schema.Schema)
	if !ok {
		return nil, ErrInvalidSchema
	}

	// Initialize relationship data tracker
	relData := &RelationshipData{
		TableData: make(map[string][]map[string]interface{}),
		References: make(map[string][]interface{}),
	}

	var generatedFiles []string

	// Handle multi-table schemas (SQL)
	if len(s.Tables) > 0 {
		// First pass: Generate data for tables without dependencies
		for _, table := range s.Tables {
			if !hasReferences(table.Fields) {
				data := generateTableDataWithConstraints(table.Fields, numRows, table.Name, relData, s.Relationships)
				relData.TableData[table.Name] = data
				populateReferences(table.Name, table.Fields, data, relData)
			}
		}

		// Second pass: Generate data for tables with dependencies
		for _, table := range s.Tables {
			if hasReferences(table.Fields) {
				data := generateTableDataWithConstraints(table.Fields, numRows, table.Name, relData, s.Relationships)
				relData.TableData[table.Name] = data
				populateReferences(table.Name, table.Fields, data, relData)
			}
		}

		// Write all generated data to files
		for _, table := range s.Tables {
			var filename string
			var err error
			
			if format == FormatJSON {
				filename = getOutputFilename(outputPath, table.Name, ".json")
				err = writeJSONFileArray(filename, relData.TableData[table.Name])
			} else {
				data := convertToStringSlicesWithHeaders(relData.TableData[table.Name], table.Fields)
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

// hasReferences checks if any field in the table has reference constraints
func hasReferences(fields []schema.Field) bool {
	for _, field := range fields {
		if field.Constraints != nil && field.Constraints.References != nil {
			return true
		}
	}
	return false
}

// generateTableDataWithConstraints generates table data while respecting relationship constraints
func generateTableDataWithConstraints(fields []schema.Field, numRows int, tableName string, relData *RelationshipData, relationships []schema.Relationship) []map[string]interface{} {
	var rows []map[string]interface{}
	
	for i := 0; i < numRows; i++ {
		row := make(map[string]interface{})
		
		for _, field := range fields {
			value := generateConstrainedValue(field, relData, tableName, i)
			row[field.Name] = value
		}
		
		rows = append(rows, row)
	}
	
	return rows
}

// generateConstrainedValue generates a value for a field considering its constraints
func generateConstrainedValue(field schema.Field, relData *RelationshipData, tableName string, rowIndex int) interface{} {
	// Handle reference constraints (foreign keys)
	if field.Constraints != nil && field.Constraints.References != nil {
		refKey := field.Constraints.References.Table + "." + field.Constraints.References.Field
		if values, exists := relData.References[refKey]; exists && len(values) > 0 {
			// Pick a random value from the referenced table
			return values[rand.IntN(len(values))]
		}
	}
	
	// Handle dependent fields
	if field.Constraints != nil && field.Constraints.DependsOn != "" {
		// This could be used for conditional generation based on other fields
		// For now, we'll generate normally but this can be extended
	}
	
	// Handle unique count constraints
	if field.Constraints != nil && field.Constraints.UniqueCount != nil {
		uniqueValues := generateUniqueValues(field, *field.Constraints.UniqueCount)
		return uniqueValues[rowIndex % len(uniqueValues)]
	}
	
	// Generate value based on type with min/max constraints
	return generateConstrainedFakeValue(field)
}

// generateConstrainedFakeValue generates a fake value with min/max constraints
func generateConstrainedFakeValue(field schema.Field) interface{} {
	switch field.Type {
	case "string", "varchar", "text":
		return faker.GenerateName()
	case "email":
		return faker.GenerateEmail()
	case "int", "integer", "serial":
		min := 1
		max := 1000
		if field.Constraints != nil {
			if field.Constraints.MinValue != nil {
				min = *field.Constraints.MinValue
			}
			if field.Constraints.MaxValue != nil {
				max = *field.Constraints.MaxValue
			}
		}
		return rand.IntN(max-min+1) + min
	case "float", "decimal", "numeric":
		value, _ := strconv.ParseFloat(faker.GenerateFloat(), 64)
		return value
	case "bool", "boolean":
		return rand.IntN(2) == 1
	case "uuid":
		return faker.GenerateUUID()
	case "phone":
		return faker.GeneratePhone()
	case "address":
		return faker.GenerateAddress()
	case "company":
		return faker.GenerateCompany()
	case "date", "timestamp":
		return faker.GenerateDate()
	default:
		return faker.GenerateName()
	}
}

// generateUniqueValues generates a set of unique values for a field
func generateUniqueValues(field schema.Field, count int) []interface{} {
	seen := make(map[interface{}]bool)
	values := make([]interface{}, 0, count)
	
	for len(values) < count {
		value := generateConstrainedFakeValue(field)
		if !seen[value] {
			seen[value] = true
			values = append(values, value)
		}
	}
	
	return values
}

// populateReferences stores generated values for use as foreign key references
func populateReferences(tableName string, fields []schema.Field, data []map[string]interface{}, relData *RelationshipData) {
	for _, field := range fields {
		refKey := tableName + "." + field.Name
		values := make([]interface{}, len(data))
		
		for i, row := range data {
			values[i] = row[field.Name]
		}
		
		relData.References[refKey] = values
	}
}

// convertToStringSlicesWithHeaders converts map data to string slices for CSV output with headers
func convertToStringSlicesWithHeaders(data []map[string]interface{}, fields []schema.Field) [][]string {
	// Create result slice with space for header + data rows
	result := make([][]string, len(data)+1)
	
	// Create header row
	header := make([]string, len(fields))
	for i, field := range fields {
		header[i] = field.Name
	}
	result[0] = header
	
	// Create data rows
	for i, row := range data {
		rowData := make([]string, len(fields))
		for j, field := range fields {
			if value, exists := row[field.Name]; exists {
				rowData[j] = fmt.Sprintf("%v", value)
			}
		}
		result[i+1] = rowData
	}
	
	return result
}

// convertToStringSlices converts map data to string slices for CSV output
func convertToStringSlices(data []map[string]interface{}, fields []schema.Field) [][]string {
	result := make([][]string, len(data))
	
	for i, row := range data {
		rowData := make([]string, len(fields))
		for j, field := range fields {
			if value, exists := row[field.Name]; exists {
				rowData[j] = fmt.Sprintf("%v", value)
			}
		}
		result[i] = rowData
	}
	
	return result
}

// writeJSONFileArray writes JSON array data to a file
func writeJSONFileArray(filename string, data []map[string]interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData := map[string]interface{}{
		"data": data,
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(jsonData)
}