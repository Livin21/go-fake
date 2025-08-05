package parser

import (
	"bufio"
	"go-fake/internal/schema"
	"os"
	"regexp"
	"strconv"
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
	var relationships []schema.Relationship
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
			field := parseFieldDefinitionWithConstraints(line, currentTable.Name, &relationships)
			if field.Name != "" {
				currentTable.Fields = append(currentTable.Fields, field)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return schema.Schema{}, err
	}

	// Add parsed relationships to schema
	s.Relationships = relationships

	return s, nil
}

// extractTableName extracts the table name from a CREATE TABLE statement.
func extractTableName(line string) string {
	// Regex to match CREATE TABLE table_name
	re := regexp.MustCompile(`CREATE\s+TABLE\s+(\w+)`)
	matches := re.FindStringSubmatch(strings.ToUpper(line))
	if len(matches) > 1 {
		// Return the original case from the line
		originalRe := regexp.MustCompile(`CREATE\s+TABLE\s+(\w+)`)
		originalMatches := originalRe.FindStringSubmatch(line)
		if len(originalMatches) > 1 {
			return originalMatches[1]
		}
	}
	return ""
}

// parseFieldDefinitionWithConstraints parses a field definition line and extracts constraints and relationships
func parseFieldDefinitionWithConstraints(line, tableName string, relationships *[]schema.Relationship) schema.Field {
	// Clean up the line
	line = strings.TrimSpace(line)
	line = strings.TrimSuffix(line, ",")
	
	if line == "" || strings.HasPrefix(line, "PRIMARY KEY") || strings.HasPrefix(line, "CONSTRAINT") {
		return schema.Field{}
	}

	parts := strings.Fields(line)
	if len(parts) < 2 {
		return schema.Field{}
	}

	field := schema.Field{
		Name:        parts[0],
		Type:        mapSQLType(parts[1]),
		Required:    false,
		Constraints: &schema.Constraint{},
	}

	// Parse constraints from the line
	upperLine := strings.ToUpper(line)
	
	// Check for NOT NULL
	if strings.Contains(upperLine, "NOT NULL") {
		field.Required = true
	}

	// Check for CHECK constraints with min/max values
	checkRe := regexp.MustCompile(`CHECK\s*\(\s*(\w+)\s*>=\s*(\d+)\s*AND\s*(\w+)\s*<=\s*(\d+)\s*\)`)
	if matches := checkRe.FindStringSubmatch(upperLine); len(matches) == 5 {
		if min, err := strconv.Atoi(matches[2]); err == nil {
			field.Constraints.MinValue = &min
		}
		if max, err := strconv.Atoi(matches[4]); err == nil {
			field.Constraints.MaxValue = &max
		}
	}

	// Check for REFERENCES (foreign key)
	refRe := regexp.MustCompile(`REFERENCES\s+(\w+)\s*\(\s*(\w+)\s*\)`)
	if matches := refRe.FindStringSubmatch(upperLine); len(matches) == 3 {
		field.Constraints.References = &schema.Reference{
			Table: matches[1],
			Field: matches[2],
		}
		
		// Add to relationships
		*relationships = append(*relationships, schema.Relationship{
			Type:      "foreign_key",
			FromTable: tableName,
			FromField: field.Name,
			ToTable:   matches[1],
			ToField:   matches[2],
			Cardinality: "many:1",
		})
	}

	// If no constraints were set, remove the empty constraints object
	if field.Constraints.References == nil && field.Constraints.MinValue == nil && field.Constraints.MaxValue == nil {
		field.Constraints = nil
	}

	return field
}

// mapSQLType maps SQL types to internal types.
func mapSQLType(sqlType string) string {
	sqlType = strings.ToUpper(sqlType)
	
	switch {
	case strings.Contains(sqlType, "SERIAL"):
		return "int"
	case strings.Contains(sqlType, "INTEGER") || strings.Contains(sqlType, "INT"):
		return "int"
	case strings.Contains(sqlType, "VARCHAR") || strings.Contains(sqlType, "TEXT"):
		return "string"
	case strings.Contains(sqlType, "DECIMAL") || strings.Contains(sqlType, "NUMERIC") || strings.Contains(sqlType, "FLOAT"):
		return "float"
	case strings.Contains(sqlType, "BOOLEAN") || strings.Contains(sqlType, "BOOL"):
		return "boolean"
	case strings.Contains(sqlType, "DATE"):
		return "date"
	case strings.Contains(sqlType, "TIMESTAMP") || strings.Contains(sqlType, "DATETIME"):
		return "timestamp"
	default:
		return "string"
	}
}