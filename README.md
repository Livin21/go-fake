# go-fake: Super fast fake data generator

A powerful, AI-enhanced command-line tool for generating realistic fake data at scale. Built with Go 1.24, go-fake intelligently parses JSON and SQL schemas to create high-quality test data with support for complex relationships, 40+ data types, and enterprise-grade performance optimizations. Perfect for developers, testers, and data engineers who need reliable fake data for development, testing, and prototyping.

## Features ✨

- **🤖 AI-Enhanced Field Inference**: OpenAI integration for intelligent field type detection
- **🧠 Intelligent Pattern Matching**: 40+ supported data types with smart field recognition
- **📋 Multiple Schema Formats**: Support for both JSON and SQL schema definitions
- **🔗 Relationship Constraints**: Foreign key relationships and referential integrity
- **⚙️ Field Constraints**: Min/max values, unique counts, and data validation
- **📁 Smart Output Format**: JSON schemas → JSON files, SQL schemas → CSV files  
- **🔄 Format Override**: Force JSON or CSV output regardless of input schema type
- **🗂️ Multi-Table Support**: Generate separate files for each table in SQL schemas
- **🎯 Rich Data Types**: 40+ supported data types for realistic fake data generation
- **🔄 Dependency Resolution**: Automatic handling of table dependencies and foreign keys
- **🛠️ Customizable Output**: Specify number of rows and output file location
- **✅ Type-Safe JSON**: Proper data types in JSON output (numbers, booleans, strings)
- **🚀 Cross-Platform Releases**: Automated releases for Linux, Windows, macOS (AMD64/ARM64)
- **⚡ Fast & Lightweight**: Built with Go 1.24 for optimal performance

## Installation 🚀

### Pre-built Binaries (Recommended)

Download the latest release for your platform from [GitHub Releases](https://github.com/Livin21/go-fake/releases):

**Linux:**
```bash
# AMD64
wget https://github.com/Livin21/go-fake/releases/latest/download/go-fake-linux-amd64
chmod +x go-fake-linux-amd64
sudo mv go-fake-linux-amd64 /usr/local/bin/go-fake

# ARM64
wget https://github.com/Livin21/go-fake/releases/latest/download/go-fake-linux-arm64
chmod +x go-fake-linux-arm64
sudo mv go-fake-linux-arm64 /usr/local/bin/go-fake
```

**macOS:**
```bash
# Intel Macs
wget https://github.com/Livin21/go-fake/releases/latest/download/go-fake-darwin-amd64
chmod +x go-fake-darwin-amd64
sudo mv go-fake-darwin-amd64 /usr/local/bin/go-fake

# Apple Silicon Macs
wget https://github.com/Livin21/go-fake/releases/latest/download/go-fake-darwin-arm64
chmod +x go-fake-darwin-arm64
sudo mv go-fake-darwin-arm64 /usr/local/bin/go-fake
```

**Windows:**
Download the appropriate `.exe` file for your architecture and add it to your PATH.

### Using Go Install

```bash
go install github.com/Livin21/go-fake/cmd/generate@latest
```

### From Source

```bash
git clone https://github.com/livin21/go-fake.git
cd go-fake
make build
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
- `-perf`: Enable performance optimizations (parallel generation, caching)
- `-workers int`: Number of parallel workers (0 = auto-detect CPU cores)
- `-batch int`: Batch size for row generation (higher = more memory, faster generation)
- `-verbose`: Enable verbose logging with detailed execution information
- `-version`: Show version information and feature status
- `-h`: Show help message with supported data types

### Performance Optimizations 🚀

Enable high-speed data generation with performance flags:

```bash
# Basic performance optimization
./bin/go-fake -schema large_schema.sql -rows 10000 -perf

# Advanced tuning for large datasets
./bin/go-fake -schema enterprise.sql -rows 100000 -perf -workers 8 -batch 2000

# Memory-constrained environments
./bin/go-fake -schema schema.sql -rows 50000 -perf -workers 4 -batch 500
```

**Performance Features:**
- **Parallel Table Generation**: Generate multiple tables concurrently
- **Configurable Worker Pools**: Control parallel processing based on system capabilities
- **Batch Processing**: Memory-efficient row generation with configurable batch sizes
- **Field Inference Caching**: Eliminate repeated pattern matching for similar fields
- **Auto-detection**: Automatically detect optimal settings based on CPU cores

**Performance Results:**
- **6 tables × 5,000 rows**: ~50ms generation time
- **Scales with CPU cores**: Performance improves with available hardware
- **Memory optimized**: Efficient memory usage with batch processing
- **Enterprise ready**: Handles large multi-table schemas efficiently

For detailed performance tuning, see **[docs/PERFORMANCE_GUIDE.md](docs/PERFORMANCE_GUIDE.md)**.

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

### High-Performance Generation Examples

```bash
# Large dataset generation with performance optimizations
./bin/go-fake -schema examples/performance_test.sql -rows 10000 -perf

# Enterprise-scale generation with custom tuning
./bin/go-fake -schema enterprise_schema.sql -rows 100000 -perf -workers 8 -batch 2000 -verbose

# Memory-constrained environment optimization
./bin/go-fake -schema large_schema.sql -rows 50000 -perf -workers 4 -batch 500

# Performance comparison output:
# Standard: 6 tables × 5,000 rows in ~50ms
# Optimized: 6 tables × 5,000 rows in ~50ms with parallel processing
# Scales with: CPU cores, batch sizes, and table independence

# CI/CD pipeline optimized generation
./bin/go-fake -schema test_data.sql -rows 1000 -perf -workers 2 -batch 1000 -output /tmp/testdata

# Performance monitoring with verbose logging
./bin/go-fake -schema schema.sql -rows 5000 -perf -verbose
# Shows: Generation time, worker utilization, memory usage, per-table timing
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
│   │   ├── performance.go # Performance optimizations & parallel processing
│   │   └── openai.go      # OpenAI API integration
│   ├── parser/           # Schema parsing (JSON/SQL)
│   └── schema/           # Schema types and validation
├── pkg/
│   ├── csv/              # CSV writing utilities  
│   └── faker/            # Fake data providers (40+ types)
├── examples/             # Example schema files
├── docs/                 # Documentation
│   ├── OPENAI_INTEGRATION.md # AI integration guide
│   └── PERFORMANCE_GUIDE.md  # Performance optimization guide
└── Makefile             # Build automation
```

### Building and Testing

```bash
# Install dependencies
make deps

# Build the application
make build

# Build with release optimizations
make build-release

# Build for all platforms
make release

# Build and test release locally
make release-local

# Run tests
make test

# Run all examples
make run-examples

# Clean build artifacts
make clean
```

### Release Process

This project uses an **automated release process** that reads from `CHANGELOG.md`:

1. **Update CHANGELOG.md:**
   ```markdown
   ## [Unreleased]
   
   ### Added
   - 📊 **New Feature** with detailed description
   - 🎯 **Another Feature** explaining the benefit
   
   ### Enhanced
   - Improved existing functionality
   - Better performance in specific areas
   
   ### Fixed
   - Bug fixes and corrections
   ```

2. **Create a new release:**
   ```bash
   ./scripts/release.sh v1.2.0
   ```
   
   The script will:
   - ✅ Validate `CHANGELOG.md` has content in `[Unreleased]` section
   - 📝 Extract release notes from `[Unreleased]` section  
   - 🔄 Update `CHANGELOG.md` by moving content to versioned section
   - 🔢 Update version in `cmd/generate/main.go`
   - 🧪 Run full test suite and build tests
   - 📤 Create git commit and tag with extracted release notes

3. **GitHub Actions automatically:**
   - 🏗️ Builds binaries for 6 platforms (Linux, Windows, macOS - AMD64/ARM64)
   - 📋 Extracts the same release notes from `CHANGELOG.md`
   - 🎁 Creates GitHub release with formatted notes and download instructions
   - 📎 Uploads binaries and checksums
   - 🌐 Updates Go module registry

**Safety Features:**
- Multiple validation layers prevent errors
- Automatic backup and recovery on failure  
- Clean error messages guide issue resolution
- Full test suite must pass before release

**Detailed Documentation:** See [docs/RELEASE_PROCESS.md](docs/RELEASE_PROCESS.md) for complete instructions.

**Supported Platforms:**
- Linux (AMD64, ARM64)
- Windows (AMD64, ARM64) 
- macOS (Intel, Apple Silicon)

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
- [x] ✅ **Automated Releases**: GitHub Actions pipeline with multi-platform binaries
- [x] ✅ **Performance Optimizations**: Parallel table generation, worker pools, and batch processing
- [x] ✅ **Field Inference Caching**: Optimized pattern matching with intelligent caching
- [x] ✅ **Verbose Logging**: Detailed execution visibility and performance monitoring
- [ ] 🔄 **Support for more output formats** (XML, Parquet, Avro)
- [ ] 🔄 **Custom faker patterns and templates**
- [ ] 🔄 **Database direct export** (PostgreSQL, MySQL, MongoDB)
- [ ] 🔄 **Web UI for schema creation** and real-time preview  
- [ ] 🔄 **Docker container support** and Kubernetes deployment
- [ ] 🔄 **Data localization** (different languages/regions)
- [ ] 🔄 **Advanced AI features** (schema generation from descriptions)
- [ ] 🔄 **Streaming generation** for extremely large datasets (>10M rows)
- [ ] 🔄 **Plugin system** for custom data generators

## Acknowledgments 🙏

- Inspired by libraries like Faker.js and Python Faker
- Built with the power of Go's standard library
- OpenAI GPT models for intelligent field inference
- Thanks to all contributors and users

---

**Happy fake data generation with AI!** 🤖🎉
