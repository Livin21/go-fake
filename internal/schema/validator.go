package schema

import "errors"

func ValidateSchema(schema Schema) error {
	// Check if schema has either tables or fields
	if len(schema.Tables) == 0 && len(schema.Fields) == 0 {
		return errors.New("schema must have at least one table or field")
	}

	// Validate tables if present
	if len(schema.Tables) > 0 {
		for _, table := range schema.Tables {
			if table.Name == "" {
				return errors.New("table name cannot be empty")
			}
			if len(table.Fields) == 0 {
				return errors.New("table must have at least one field")
			}
			if err := validateFields(table.Fields); err != nil {
				return err
			}
		}
	}

	// Validate fields if present (backward compatibility)
	if len(schema.Fields) > 0 {
		if err := validateFields(schema.Fields); err != nil {
			return err
		}
	}

	return nil
}

func validateFields(fields []Field) error {
	fieldNames := make(map[string]bool)
	for _, field := range fields {
		if field.Name == "" {
			return errors.New("field name cannot be empty")
		}
		if fieldNames[field.Name] {
			return errors.New("field names must be unique")
		}
		fieldNames[field.Name] = true

		if field.Type == "" {
			return errors.New("field type cannot be empty")
		}
	}
	return nil
}