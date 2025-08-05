package schema

type Schema struct {
    Tables []Table `json:"tables,omitempty"`
    Fields []Field `json:"fields,omitempty"` // For backward compatibility with simple schemas
    Relationships []Relationship `json:"relationships,omitempty"` // New: Define relationships
}

type Table struct {
    Name   string  `json:"name"`
    Fields []Field `json:"fields"`
}

type Field struct {
    Name        string      `json:"name"`
    Type        string      `json:"type"`
    Required    bool        `json:"required"`
    Constraints *Constraint `json:"constraints,omitempty"` // New: Field-level constraints
}

// New: Relationship constraints between tables/fields
type Relationship struct {
    Type         string `json:"type"` // "foreign_key", "one_to_many", "many_to_many"
    FromTable    string `json:"from_table"`
    FromField    string `json:"from_field"`
    ToTable      string `json:"to_table"`  
    ToField      string `json:"to_field"`
    Cardinality  string `json:"cardinality,omitempty"` // "1:1", "1:many", "many:many"
}

// New: Field-level constraints
type Constraint struct {
    References   *Reference `json:"references,omitempty"`   // Foreign key reference
    DependsOn    string     `json:"depends_on,omitempty"`   // Field dependency
    Pattern      string     `json:"pattern,omitempty"`      // Regex pattern
    MinValue     *int       `json:"min_value,omitempty"`    // Minimum value
    MaxValue     *int       `json:"max_value,omitempty"`    // Maximum value
    UniqueCount  *int       `json:"unique_count,omitempty"` // Number of unique values
}

// New: Reference constraint for foreign keys
type Reference struct {
    Table string `json:"table"`
    Field string `json:"field"`
}