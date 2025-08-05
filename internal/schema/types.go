package schema

type Schema struct {
    Tables []Table `json:"tables,omitempty"`
    Fields []Field `json:"fields,omitempty"` // For backward compatibility with simple schemas
}

type Table struct {
    Name   string  `json:"name"`
    Fields []Field `json:"fields"`
}

type Field struct {
    Name     string `json:"name"`
    Type     string `json:"type"`
    Required bool   `json:"required"`
}