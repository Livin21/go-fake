# go-fake

A powerful CLI tool for generating fake data based on JSON or SQL schema definitions. Perfect for testing, development, and prototyping.

## Features ✨

- **🤖 AI-Enhanced Field Inference**: OpenAI integration for intelligent field type detection
- **🧠 Intelligent Pattern Matching**: 40+ supported data types with smart field recognition
- **📋 Multiple Schema Formats**: Support for both JSON and SQL schema definitions
- **🔗 Relationship Constraints**: Foreign key relationships and referential integrity
- **⚙️ Field Constraints**: Min/max values, unique counts, and data validation
- **📁 Smart Output Format**: JSON schemas → JSON files, SQL schemas → CSV files  
- **� Format Override**: Force JSON or CSV output regardless of input schema type
- **�🗂️ Multi-Table Support**: Generate separate files for each table in SQL schemas
- **🎯 Rich Data Types**: 40+ supported data types for realistic fake data generation
- **🔄 Dependency Resolution**: Automatic handling of table dependencies and foreign keys
- **🛠️ Customizable Output**: Specify number of rows and output file location
- **✅ Type-Safe JSON**: Proper data types in JSON output (numbers, booleans, strings)
- **⚡ Fast & Lightweight**: Built with Go 1.24 for optimal performance

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

# Override output format
./bin/go-fake -schema examples/sample.json -format csv -output users.csv
./bin/go-fake -schema examples/sample.sql -format json -output users_json/

# Check version and feature status
./bin/go-fake --version
# Shows: AI integration status, supported data types, feature list
```

### Output Format

The tool automatically determines the output format based on your input schema, but you can override this behavior:

- **JSON Schema Input** (`.json`) → **JSON Output Files** (default)
- **SQL Schema Input** (`.sql`) → **CSV Output Files** (default)
- **Multi-Table Schemas** → **Directory with separate files per table**
- **Format Override** → Use `-format json` or `-format csv` to override automatic detection

#### Format Override Examples

```bash
# Force JSON schema to output CSV format
./bin/go-fake -schema schema.json -format csv -output data.csv

# Force SQL schema to output JSON format  
./bin/go-fake -schema schema.sql -format json -output data_dir/

# Multi-table SQL schema as JSON files in directory
./bin/go-fake -schema multi_table.sql -format json -output json_output/
# Creates: json_output/users.json, json_output/products.json, etc.
```

### Command Line Options

- `-schema string`: Path to the schema file (JSON or SQL) - **Required**
- `-output string`: Output directory for multi-table schemas or file path for single-table schemas (default: "output.csv" or "output.json")
- `-rows int`: Number of rows to generate (default: 100)
- `-format string`: Override output format (`json` or `csv`). If not specified, format is auto-detected from schema type
- `-ai`: Enable OpenAI-powered field inference for ambiguous field names (requires OPENAI_API_KEY)
- `-version`: Show version information and feature status
- `-h`: Show help message with supported data types

### AI Enhancement 🤖

Set the `OPENAI_API_KEY` environment variable to enable AI-powered field inference:

```bash
export OPENAI_API_KEY="your-openai-api-key"

# Use AI mode for enhanced field detection
./bin/go-fake -schema schema.json -ai -output enhanced_data.json
```

**AI Benefits:**
- **Intelligent Field Inference**: Analyzes ambiguous field names like `user_handle`, `secret_code`, `network_endpoint`
- **Contextual Data Generation**: Generates more realistic, contextually appropriate data
- **Graceful Fallback**: Falls back to intelligent pattern matching when API is unavailable
- **Enhanced Schema Understanding**: Provides suggestions and documentation for complex schemas

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

**The tool intelligently detects field types using multi-layered inference:**
1. **Exact Pattern Matching**: Direct field name matches (e.g., `email` → email addresses)
2. **Partial Pattern Matching**: Substring matches (e.g., `user_email` → email addresses)  
3. **Semantic Understanding**: Context analysis (e.g., `contact_method` → phone/email)
4. **AI Enhancement**: OpenAI analysis for ambiguous fields (with `-ai` flag)
5. **Regex Analysis**: Pattern-based detection for complex field names

### Comprehensive Data Type Support

| Category | Types | Example Output |
|----------|-------|----------------|
| **Basic Types** | `string`, `int`, `float`, `bool`, `date`, `datetime` | "Sample Text", 42, 123.45, true, "2023-05-15" |
| **Identity** | `email`, `name`, `firstname`, `lastname`, `username`, `uuid` | alice@example.com, John Doe, ultrajohnpro |
| **Contact Info** | `phone`, `address`, `city`, `state`, `zipcode`, `country` | (555) 123-4567, 123 Main St, New York, CA |
| **Business** | `company`, `jobtitle`, `department`, `category`, `price` | TechCorp, Software Engineer, Engineering |
| **Technical** | `url`, `image`, `ipaddress`, `macaddress`, `version`, `filename` | https://example.com, 192.168.1.1, v1.2.3 |
| **Security** | `password`, `creditcard`, `bankaccount`, `ssn`, `license` | $eC9rE!@, 4532-1234-5678-9012 |
| **Content** | `text`, `hashtag`, `color`, `product`, `brand`, `skill` | #trending, #FF5733, Wireless Headphones |
| **Measurements** | `age`, `height`, `weight`, `temperature`, `longitude`, `latitude` | 28, 5.8, 165.5, 72.3°F, -122.4194 |
| **System** | `status`, `priority`, `duration`, `gender` | active, high, 2h 30m, male |

### AI-Enhanced Field Detection Examples

```bash
# Standard detection
"user_email" → email addresses
"phone_number" → phone numbers  
"company_name" → company names

# AI-enhanced detection (with -ai flag)
"network_endpoint" → IP addresses or URLs
"secret_code" → secure passwords
"personal_identifier" → UUIDs or IDs
"business_entity" → company names
"contact_method" → names or emails
"measurement_reading" → decimal numbers
```

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
#   - ecommerce/users.csv
#   - ecommerce/products.csv  
#   - ecommerce/orders.csv
```

### AI-Enhanced Generation Examples

```bash
# Generate data with AI field inference
export OPENAI_API_KEY="your-api-key"

# Schema with ambiguous field names
cat > ambiguous_schema.json << EOF
{
  "fields": [
    {"name": "user_handle", "type": "string"},
    {"name": "secret_code", "type": "string"}, 
    {"name": "network_endpoint", "type": "string"},
    {"name": "business_entity", "type": "string"}
  ]
}
EOF

# Standard mode (pattern matching)
./bin/go-fake -schema ambiguous_schema.json -output standard.json

# AI-enhanced mode (intelligent inference)
./bin/go-fake -schema ambiguous_schema.json -ai -output enhanced.json

# Format override with AI enhancement
./bin/go-fake -schema ambiguous_schema.json -ai -format csv -output enhanced.csv

# AI Output Examples:
# user_handle: "ultrajohnpro", "megaalice123"
# secret_code: "$eC9rE!@", "Kbt^Aje^fYS" 
# network_endpoint: "192.168.1.1", "api.example.com"
# business_entity: "TechCorp", "InnovateLab"
```

### Format Override Examples

```bash
# Convert JSON schema to CSV output
./bin/go-fake -schema user_profile.json -format csv -output users.csv

# Convert SQL schema to JSON output
./bin/go-fake -schema database.sql -format json -output json_data/

# Multi-table SQL to individual JSON files
./bin/go-fake -schema ecommerce.sql -format json -rows 500 -output ecommerce_json/
# Creates: ecommerce_json/users.json, ecommerce_json/products.json, etc.

# Error handling for invalid formats
./bin/go-fake -schema schema.json -format xml -output data
# Error: unsupported format: xml (supported: json, csv)
```

## Development 🛠️

### Project Structure

```
go-fake/
├── cmd/generate/          # CLI application entry point
├── internal/
│   ├── generator/         # Data generation logic
│   │   ├── generator.go   # Core generation functions
│   │   ├── intelligent.go # Intelligent field type inference
│   │   └── openai.go      # OpenAI API integration
│   ├── parser/           # Schema parsing (JSON/SQL)
│   └── schema/           # Schema types and validation
├── pkg/
│   ├── csv/              # CSV writing utilities  
│   └── faker/            # Fake data providers (40+ types)
├── examples/             # Example schema files
├── docs/                 # Documentation
│   └── OPENAI_INTEGRATION.md # AI integration guide
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
2. Add the pattern to `internal/generator/intelligent.go` inference maps
3. Add the case to the switch statement in `generateValueByType()`
4. Update the help text in `cmd/generate/main.go`
5. Add tests for the new functionality

### AI Integration Development

For OpenAI integration development, see **[docs/OPENAI_INTEGRATION.md](docs/OPENAI_INTEGRATION.md)** for:
- API integration patterns
- Field inference strategies  
- Fallback mechanisms
- Cost optimization techniques

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

- [x] ✅ **AI-Enhanced Field Inference**: OpenAI integration for intelligent field detection
- [x] ✅ **40+ Data Types**: Comprehensive fake data generation across multiple domains  
- [x] ✅ **Intelligent Pattern Matching**: Multi-layered field type inference
- [x] ✅ **Relationship Constraints**: Foreign key relationships between fields
- [x] ✅ **Directory-based Output**: Organized multi-table file generation
- [x] ✅ **Format Override**: JSON/CSV output format control regardless of input schema
- [ ] 🔄 **Support for more output formats** (XML, Parquet, Avro)
- [ ] 🔄 **Custom faker patterns and templates**
- [ ] 🔄 **Database direct export** (PostgreSQL, MySQL, MongoDB)
- [ ] 🔄 **Web UI for schema creation** and real-time preview  
- [ ] 🔄 **Docker container support** and Kubernetes deployment
- [ ] 🔄 **Data localization** (different languages/regions)
- [ ] 🔄 **Advanced AI features** (schema generation from descriptions)
- [ ] 🔄 **Performance optimizations** for large datasets (>1M rows)
- [ ] 🔄 **Plugin system** for custom data generators

## Acknowledgments 🙏

- Inspired by libraries like Faker.js and Python Faker
- Built with the power of Go's standard library
- OpenAI GPT models for intelligent field inference
- Thanks to all contributors and users

---

**Happy fake data generation with AI!** 🤖🎉