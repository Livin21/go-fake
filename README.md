# go-fake

A powerful CLI tool for generating fake data based on JSON or SQL schema definitions. Perfect for testing, development, and prototyping.

## Features ✨

- **Multiple Schema Formats**: Support for both JSON and SQL schema definitions
- **Relationship Constraints**: Foreign key relationships and referential integrity
- **Field Constraints**: Min/max values, unique counts, and data validation
- **Smart Output Format**: JSON schemas → JSON files, SQL schemas → CSV files  
- **Multi-Table Support**: Generate separate files for each table in SQL schemas
- **Rich Data Types**: 17+ supported data types for realistic fake data generation
- **Dependency Resolution**: Automatic handling of table dependencies and foreign keys
- **Customizable Output**: Specify number of rows and output file location
- **Type-Safe JSON**: Proper data types in JSON output (numbers, booleans, strings)
- **Fast & Lightweight**: Built with Go 1.24 for optimal performance

## Installation 🚀

### From Source

```bash
git clone https://github.com/livin21/go-fake.git
cd go-fake
make build
```

### Using Go Install

```bash
go install github.com/livin21/go-fake/cmd/generate@latest
```

## Usage 📖

### Basic Usage

```bash
# Generate JSON from JSON schema
./bin/go-fake -schema examples/sample.json -output users.json

# Generate multi-table CSV from SQL schema (creates directory)
./bin/go-fake -schema examples/sample.sql -output company_data
# Creates: company_data/users.csv, company_data/products.csv, company_data/orders.csv

# Specify number of rows
./bin/go-fake -schema examples/comprehensive.json -rows 1000 -output large-dataset
```

### Output Format

The tool automatically determines the output format based on your input schema:

- **JSON Schema Input** (`.json`) → **JSON Output Files**
- **SQL Schema Input** (`.sql`) → **CSV Output Files**
- **Multi-Table Schemas** → **Directory with separate files per table**

### Command Line Options

- `-schema string`: Path to the schema file (JSON or SQL) - **Required**
- `-output string`: Output directory for multi-table schemas or file path for single-table schemas (default: "output.csv" or "output.json")
- `-rows int`: Number of rows to generate (default: 100)
- `-h`: Show help message with supported data types

## Schema Formats 📋

### JSON Schema Format

Support for single-table and multi-table schemas with relationship constraints:

```json
{
  "tables": [
    {
      "name": "users",
      "fields": [
        {
          "name": "id",
          "type": "int",
          "required": true,
          "constraints": {
            "min_value": 1,
            "max_value": 1000
          }
        },
        {
          "name": "email",
          "type": "email", 
          "required": true
        },
        {
          "name": "age",
          "type": "int",
          "required": false,
          "constraints": {
            "min_value": 18,
            "max_value": 80
          }
        }
      ]
    },
    {
      "name": "employees", 
      "fields": [
        {
          "name": "user_id",
          "type": "int",
          "required": true,
          "constraints": {
            "references": {
              "table": "users",
              "field": "id"
            }
          }
        }
      ]
    }
  ],
  "relationships": [
    {
      "type": "foreign_key",
      "from_table": "employees",
      "from_field": "user_id",
      "to_table": "users", 
      "to_field": "id",
      "cardinality": "many:1"
    }
  ]
}
```

### SQL Schema Format

Standard CREATE TABLE syntax with relationship constraints:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    age INTEGER CHECK (age >= 18 AND age <= 80)
);

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    salary DECIMAL(10,2) CHECK (salary >= 30000 AND salary <= 150000),
    hire_date DATE NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Multi-Table Output**: When using SQL schemas with multiple tables, the tool automatically creates an output directory with separate files for each table:
```
output/
├── users.csv
├── products.csv
└── orders.csv
```

## Relationship Constraints 🔗

The tool supports sophisticated relationship constraints for generating realistic, interconnected data:

### Foreign Key Relationships

**JSON Schema**:
```json
{
  "name": "user_id",
  "type": "int", 
  "constraints": {
    "references": {
      "table": "users",
      "field": "id"
    }
  }
}
```

**SQL Schema** (automatically parsed):
```sql
user_id INTEGER NOT NULL REFERENCES users(id)
```

### Field Constraints

- **Min/Max Values**: `"min_value": 18, "max_value": 65`
- **Unique Count**: `"unique_count": 5` (generate only 5 unique values)
- **Value Ranges**: `CHECK (salary >= 30000 AND salary <= 150000)`

### Generation Order

1. **Dependency Analysis**: Identifies table relationships
2. **Primary Tables First**: Generates tables without dependencies 
3. **Foreign Key Resolution**: Uses actual generated values for references
4. **Referential Integrity**: Guarantees all foreign keys are valid

Example with relationships:
```bash
./go-fake -schema examples/relationships.json -rows 10 -output company_data
# Creates directory: company_data/
# - company_data/users.json (primary table)
# - company_data/departments.json (primary table) 
# - company_data/employees.json (references users and departments)
```

For detailed relationship constraint documentation, see **[RELATIONSHIPS.md](RELATIONSHIPS.md)**.

## Supported Data Types 🎯

| Type | Description | Example Output |
|------|-------------|----------------|
| `string`, `varchar`, `text` | Random full names | John Smith |
| `email` | Random email addresses | john.doe@example.com |
| `int`, `integer`, `serial` | Random integers | 42 |
| `float`, `decimal`, `numeric` | Random decimal numbers | 123.45 |
| `bool`, `boolean` | Random true/false | true |
| `date` | Random dates | 2023-05-15 |
| `datetime`, `timestamp` | Random datetimes | 2023-05-15 14:30:22 |
| `phone` | Random phone numbers | (555) 123-4567 |
| `company` | Random company names | TechCorp |
| `address` | Random street addresses | 123 Main St |
| `city` | Random city names | New York |
| `state` | Random state codes | CA |
| `zipcode`, `zip` | Random ZIP codes | 12345 |
| `uuid` | Random UUIDs | 123e4567-e89b-12d3-a456-426614174000 |
| `price` | Random prices | 29.99 |
| `firstname` | Random first names | Alice |
| `lastname` | Random last names | Johnson |

## Examples 💡

### Generate User Data (JSON Output)

```bash
# Create users.json schema
cat > users.json << EOF
{
  "fields": [
    {"name": "id", "type": "uuid", "required": true},
    {"name": "first_name", "type": "firstname", "required": true},
    {"name": "last_name", "type": "lastname", "required": true},
    {"name": "email", "type": "email", "required": true},
    {"name": "phone", "type": "phone", "required": false},
    {"name": "created_at", "type": "datetime", "required": true}
  ]
}
EOF

# Generate 500 users as JSON
./bin/go-fake -schema users.json -rows 500 -output users.json

# Output: users.json with properly typed JSON data
{
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "first_name": "Alice",
      "last_name": "Johnson", 
      "email": "alice.johnson@example.com",
      "phone": "(555) 123-4567",
      "created_at": "2023-05-15 14:30:22"
    }
  ]
}
```

### Generate E-commerce Data (Multi-Table CSV)

```bash
# Using the sample.sql with multiple tables
./bin/go-fake -schema examples/sample.sql -rows 1000 -output ecommerce

# Output:
# Fake data generated and written to 3 files:
#   - ecommerce_users.csv
#   - ecommerce_products.csv  
#   - ecommerce_orders.csv
```

## Development 🛠️

### Project Structure

```
go-fake/
├── cmd/generate/          # CLI application entry point
├── internal/
│   ├── generator/         # Data generation logic
│   ├── parser/           # Schema parsing (JSON/SQL)
│   └── schema/           # Schema types and validation
├── pkg/
│   ├── csv/              # CSV writing utilities  
│   └── faker/            # Fake data providers
├── examples/             # Example schema files
└── Makefile             # Build automation
```

### Building and Testing

```bash
# Install dependencies
make deps

# Build the application
make build

# Run tests
make test

# Run all examples
make run-examples

# Clean build artifacts
make clean
```

### Adding New Data Types

1. Add the generator function to `pkg/faker/providers.go`
2. Add the case to the switch statement in `internal/generator/generator.go`
3. Update the help text in `cmd/generate/main.go`
4. Add tests for the new functionality

## Contributing 🤝

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Roadmap 🗺️

- [ ] Support for more output formats (JSON, XML, Parquet)
- [ ] Custom faker patterns and templates
- [ ] Database direct export (PostgreSQL, MySQL)
- [ ] Web UI for schema creation
- [ ] Docker container support
- [ ] Relationship constraints between fields
- [ ] Data localization (different languages/regions)

## Acknowledgments 🙏

- Inspired by libraries like Faker.js and Python Faker
- Built with the power of Go's standard library
- Thanks to all contributors and users

---

**Happy fake data generation!** 🎉