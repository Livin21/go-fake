package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"go-fake/internal/generator"
	"go-fake/internal/parser"
)

const version = "1.0.0"

func main() {
	schemaFile := flag.String("schema", "", "Path to the schema file (JSON or SQL)")
	outputFile := flag.String("output", "output.csv", "Path to the output CSV file")
	numRows := flag.Int("rows", 100, "Number of rows to generate")
	showVersion := flag.Bool("version", false, "Show version information")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "go-fake v%s - Generate fake data based on JSON or SQL schema files.\n\n", version)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS]\n\n", "go-fake")
		fmt.Fprintf(flag.CommandLine.Output(), "Output Format:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  JSON schemas (.json) -> JSON output files\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  SQL schemas (.sql)   -> CSV output files\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nSupported field types:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  string, varchar, text - Random names\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  email                 - Random email addresses\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  int, integer, serial  - Random integers\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  float, decimal        - Random decimal numbers\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  bool, boolean         - Random true/false\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  date                  - Random dates\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  datetime, timestamp   - Random datetimes\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  phone                 - Random phone numbers\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  company               - Random company names\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  address               - Random street addresses\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  city                  - Random city names\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  state                 - Random state codes\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  zipcode, zip          - Random ZIP codes\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  uuid                  - Random UUIDs\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  price                 - Random prices\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  firstname             - Random first names\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  lastname              - Random last names\n")
	}
	flag.Parse()

	if *showVersion {
		fmt.Printf("go-fake v%s\n", version)
		return
	}

	if *schemaFile == "" {
		flag.Usage()
		log.Fatal("\nError: Schema file is required")
	}

	// Check if schema file exists
	if _, err := os.Stat(*schemaFile); os.IsNotExist(err) {
		log.Fatalf("Error: Schema file '%s' does not exist", *schemaFile)
	}

	var schema interface{}
	var err error

	if strings.HasSuffix(strings.ToLower(*schemaFile), ".json") {
		schema, err = parser.ParseJSONSchema(*schemaFile)
	} else {
		schema, err = parser.ParseSQLSchema(*schemaFile)
	}

	if err != nil {
		log.Fatalf("Error parsing schema: %v", err)
	}

	// Determine output format based on input schema type
	var format generator.OutputFormat
	if strings.HasSuffix(strings.ToLower(*schemaFile), ".json") {
		format = generator.FormatJSON
	} else {
		format = generator.FormatCSV
	}

	generatedFiles, err := generator.GenerateDataFiles(schema, *numRows, *outputFile, format)
	if err != nil {
		log.Fatalf("Error generating data: %v", err)
	}

	if len(generatedFiles) == 1 {
		fmt.Printf("Fake data generated and written to %s\n", generatedFiles[0])
	} else {
		fmt.Printf("Fake data generated and written to %d files:\n", len(generatedFiles))
		for _, file := range generatedFiles {
			fmt.Printf("  - %s\n", file)
		}
	}
}