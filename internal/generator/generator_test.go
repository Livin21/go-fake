package generator

import (
	"go-fake/internal/schema"
	"testing"
)

func TestGenerateData(t *testing.T) {
	tests := []struct {
		name          string
		schema        interface{}
		wantError     bool
		wantRowCount  int // We'll check the number of rows instead of exact content
		wantColCount  int // We'll check the number of columns
	}{
		{
			name: "Test with simple schema",
			schema: schema.Schema{
				Fields: []schema.Field{
					{Name: "id", Type: "int", Required: true},
					{Name: "name", Type: "string", Required: true},
				},
			},
			wantError:    false,
			wantRowCount: 101, // 100 data rows + 1 header row
			wantColCount: 2,
		},
		{
			name: "Test with table schema",
			schema: schema.Schema{
				Tables: []schema.Table{
					{
						Name: "users",
						Fields: []schema.Field{
							{Name: "id", Type: "int", Required: true},
							{Name: "name", Type: "string", Required: true},
						},
					},
				},
			},
			wantError:    false,
			wantRowCount: 101, // 100 data rows + 1 header row
			wantColCount: 2,
		},
		{
			name:      "Test with invalid schema",
			schema:    "invalid",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateData(tt.schema, 100)
			
			if tt.wantError {
				if err == nil {
					t.Errorf("GenerateData() expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("GenerateData() unexpected error: %v", err)
				return
			}
			
			if len(got) != tt.wantRowCount {
				t.Errorf("GenerateData() row count = %v, want %v", len(got), tt.wantRowCount)
			}
			
			if len(got) > 0 && len(got[0]) != tt.wantColCount {
				t.Errorf("GenerateData() column count = %v, want %v", len(got[0]), tt.wantColCount)
			}
		})
	}
}