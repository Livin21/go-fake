package parser

import (
	"bufio"
	"go-fake/internal/schema"
	"os"
	"regexp"
	"strings"
)

// ParseSQLSchema reads an SQL file and returns a structured schema with multiple tables.
func ParseSQLSchema(filePath string) (schema.Schema, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return schema.Schema{}, err
	}
	defer file.Close()

	var s schema.Schema
	scanner := bufio.NewScanner(file)
	
	var currentTable *schema.Table
	var inTableDefinition bool
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "--") {
			continue // Skip empty lines and comments
		}

		// Check for CREATE TABLE statement
		if strings.HasPrefix(strings.ToUpper(line), "CREATE TABLE") {
			tableName := extractTableName(line)
			if tableName != "" {
				currentTable = &schema.Table{
					Name:   tableName,
					Fields: []schema.Field{},
				}
				inTableDefinition = true
			}
			continue
		}

		// Check for end of table definition
		if strings.Contains(line, ");") {
			if currentTable != nil {
				s.Tables = append(s.Tables, *currentTable)
				currentTable = nil
			}
			inTableDefinition = false
			continue
		}

		// Parse field definitions within table
		if inTableDefinition && currentTable != nil {
			field := parseFieldDefinition(line)
			if field.Name != "" {
				currentTable.Fields = append(currentTable.Fields, field)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return schema.Schema{}, err
	}

	return s, nil
}

// extractTableName extracts table name from CREATE TABLE statement
func extractTableName(line string) string {
	// Match "CREATE TABLE tablename" or "CREATE TABLE tablename ("
	re := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+(\w+)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// parseFieldDefinition parses a field definition line like "id SERIAL PRIMARY KEY,"
func parseFieldDefinition(line string) schema.Field {
	// Remove trailing comma and parentheses
	line = strings.TrimSuffix(strings.TrimSpace(line), ",")
	line = strings.TrimSuffix(line, "(")
	
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return schema.Field{}
	}

	fieldName := parts[0]
	fieldType := strings.ToLower(parts[1])
	
	// Map SQL types to our faker types
	mappedType := mapSQLType(fieldType)
	
	// Check if field is required (not NULL, PRIMARY KEY, etc.)
	isRequired := strings.Contains(strings.ToUpper(line), "NOT NULL") || 
				 strings.Contains(strings.ToUpper(line), "PRIMARY KEY")

	return schema.Field{
		Name:     fieldName,
		Type:     mappedType,
		Required: isRequired,
	}
}

// mapSQLType maps SQL types to our faker types
func mapSQLType(sqlType string) string {
	switch {
	case strings.Contains(sqlType, "serial"), strings.Contains(sqlType, "int"):
		return "int"
	case strings.Contains(sqlType, "varchar"), strings.Contains(sqlType, "text"), strings.Contains(sqlType, "char"):
		return "string"
	case strings.Contains(sqlType, "decimal"), strings.Contains(sqlType, "numeric"), strings.Contains(sqlType, "float"):
		return "decimal"
	case strings.Contains(sqlType, "timestamp"), strings.Contains(sqlType, "datetime"):
		return "datetime"
	case strings.Contains(sqlType, "date"):
		return "date"
	case strings.Contains(sqlType, "bool"):
		return "bool"
	default:
		return "string"
	}
}