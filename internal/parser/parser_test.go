package parser

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseJSONSchema(t *testing.T) {
	// Create a temporary JSON file for testing
	jsonContent := `{
		"fields": [
			{"name": "id", "type": "int", "required": true},
			{"name": "name", "type": "string", "required": false}
		]
	}`
	
	tmpFile, err := ioutil.TempFile("", "test-schema-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	
	if _, err := tmpFile.Write([]byte(jsonContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()
	
	result, err := ParseJSONSchema(tmpFile.Name())
	if err != nil {
		t.Fatalf("ParseJSONSchema() error = %v", err)
	}
	
	if len(result.Fields) != 2 {
		t.Errorf("Expected 2 fields, got %d", len(result.Fields))
	}
	
	if result.Fields[0].Name != "id" || result.Fields[0].Type != "int" {
		t.Errorf("First field incorrect: got %+v", result.Fields[0])
	}
}

func TestParseSQLSchema(t *testing.T) {
	// Create a temporary SQL file for testing with CREATE TABLE syntax
	sqlContent := `CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);`
	
	tmpFile, err := ioutil.TempFile("", "test-schema-*.sql")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	
	if _, err := tmpFile.Write([]byte(sqlContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()
	
	result, err := ParseSQLSchema(tmpFile.Name())
	if err != nil {
		t.Fatalf("ParseSQLSchema() error = %v", err)
	}
	
	if len(result.Tables) != 1 {
		t.Errorf("Expected 1 table, got %d", len(result.Tables))
	}
	
	if len(result.Tables) > 0 {
		table := result.Tables[0]
		if table.Name != "users" {
			t.Errorf("Expected table name 'users', got '%s'", table.Name)
		}
		
		if len(table.Fields) != 3 {
			t.Errorf("Expected 3 fields, got %d", len(table.Fields))
		}
		
		if len(table.Fields) > 0 && (table.Fields[0].Name != "id" || table.Fields[0].Type != "int") {
			t.Errorf("First field incorrect: got %+v", table.Fields[0])
		}
	}
}