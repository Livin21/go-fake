package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"go-fake/internal/generator"
	"go-fake/internal/parser"
	"go-fake/internal/schema"
)

const version = "v1.1.0"

func main() {
	schemaFile := flag.String("schema", "", "Path to the schema file (JSON or SQL)")
	outputFile := flag.String("output", "output.csv", "Output directory for multi-table schemas or file path for single-table schemas")
	numRows := flag.Int("rows", 100, "Number of rows to generate")
	showVersion := flag.Bool("version", false, "Show version information")
	enableAI := flag.Bool("ai", false, "Enable OpenAI-powered field inference (requires OPENAI_API_KEY)")
	outputFormat := flag.String("format", "", "Override output format (json or csv). If not specified, format is auto-detected from schema type")
	
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "go-fake v%s - AI-Enhanced Fake Data Generator\n\n", version)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS]\n\n", "go-fake")
		fmt.Fprintf(flag.CommandLine.Output(), "Output Format:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  JSON schemas (.json) -> JSON output files (default)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  SQL schemas (.sql)   -> CSV output files (default)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Multi-table schemas  -> Creates directory with separate files per table\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Use -format to override automatic format detection\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nAI Enhancement:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Set OPENAI_API_KEY environment variable to enable AI-powered field inference\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Use -ai flag to enable AI mode for ambiguous field detection\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Supported field types:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Basic: string, int, float, bool, date, datetime\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Identity: email, name, firstname, lastname, username, uuid\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Contact: phone, address, city, state, zipcode, country\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Business: company, jobtitle, department, category, price\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Technical: url, image, ipaddress, macaddress, version, filename\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Security: password, creditcard, bankaccount, ssn, license\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Content: text, hashtag, color, product, brand, skill\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Measurements: age, height, weight, temperature, longitude, latitude\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  System: status, priority, duration, gender\n")
	}
	flag.Parse()

	if *showVersion {
		fmt.Printf("go-fake v%s - AI-Enhanced Fake Data Generator\n", version)
		fmt.Println("Features:")
		fmt.Println("  - Intelligent field type detection")
		fmt.Println("  - 40+ supported data types")
		fmt.Println("  - Relationship constraints")
		fmt.Println("  - Directory-based output")
		fmt.Printf("  - OpenAI integration (%s)\n", getAIStatus())
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

	// Parse schema based on file extension
	var schemaData schema.Schema
	var err error

	if strings.HasSuffix(strings.ToLower(*schemaFile), ".json") {
		schemaData, err = parser.ParseJSONSchema(*schemaFile)
		if err != nil {
			log.Fatalf("Error parsing JSON schema: %v", err)
		}
	} else {
		schemaData, err = parser.ParseSQLSchema(*schemaFile)
		if err != nil {
			log.Fatalf("Error parsing SQL schema: %v", err)
		}
	}

	// Generate fake data with optional AI enhancement
	var generatedFiles []string
	if *enableAI {
		fmt.Printf("AI-enhanced mode enabled (OpenAI API: %s)\n", getAIStatus())
		generatedFiles, err = generator.GenerateWithAIAndFormat(&schemaData, *numRows, *outputFile, *outputFormat)
	} else {
		generatedFiles, err = generator.GenerateWithFormat(&schemaData, *numRows, *outputFile, *outputFormat)
	}

	if err != nil {
		log.Fatalf("Error generating data: %v", err)
	}

	// Display results
	if len(generatedFiles) == 1 {
		fmt.Printf("Fake data generated and written to: %s\n", generatedFiles[0])
	} else {
		fmt.Printf("Fake data generated and written to %d files:\n", len(generatedFiles))
		for _, file := range generatedFiles {
			fmt.Printf("  - %s\n", file)
		}
	}
}

func getAIStatus() string {
	if os.Getenv("OPENAI_API_KEY") != "" {
		return "Available"
	}
	return "Not configured (set OPENAI_API_KEY)"
}
