package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"go-fake/internal/generator"
	"go-fake/internal/parser"
	"go-fake/internal/schema"
	"go-fake/pkg/logger"
)

const version = "v1.3.0"

func main() {
	schemaFile := flag.String("schema", "", "Path to the schema file (JSON or SQL)")
	outputFile := flag.String("output", "output.csv", "Output directory for multi-table schemas or file path for single-table schemas")
	numRows := flag.Int("rows", 100, "Number of rows to generate")
	showVersion := flag.Bool("version", false, "Show version information")
	enableAI := flag.Bool("ai", false, "Enable OpenAI-powered field inference (requires OPENAI_API_KEY)")
	outputFormat := flag.String("format", "", "Override output format (json or csv). If not specified, format is auto-detected from schema type")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	enablePerf := flag.Bool("perf", false, "Enable performance optimizations (parallel generation, caching)")
	workers := flag.Int("workers", 0, "Number of parallel workers (0 = auto-detect CPU cores)")
	batchSize := flag.Int("batch", 1000, "Batch size for row generation (higher = more memory, faster generation)")
	
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
		fmt.Fprintf(flag.CommandLine.Output(), "Logging:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Use -verbose flag to enable detailed execution logging\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Shows schema parsing, data generation progress, and timing information\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Performance:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Use -perf flag to enable parallel generation and caching optimizations\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -workers N: Set number of parallel workers (0 = auto-detect CPU cores)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -batch N: Set batch size for row generation (higher = faster, more memory)\n\n")
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

	// Initialize logger with verbose mode
	logger.Init(*verbose)

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
		logger.Fatal("Schema file is required")
	}

	logger.Step("1", "Validating schema file")
	
	// Check if schema file exists
	if _, err := os.Stat(*schemaFile); os.IsNotExist(err) {
		logger.Fatal("Schema file '%s' does not exist", *schemaFile)
	}
	
	logger.Debug("Schema file '%s' exists", *schemaFile)

	// Parse schema based on file extension
	var schemaData schema.Schema
	var err error

	logger.Step("2", "Parsing schema file")
	
	if strings.HasSuffix(strings.ToLower(*schemaFile), ".json") {
		logger.Debug("Detected JSON schema format")
		schemaData, err = parser.ParseJSONSchema(*schemaFile)
		if err != nil {
			logger.Fatal("Error parsing JSON schema: %v", err)
		}
	} else {
		logger.Debug("Detected SQL schema format")
		schemaData, err = parser.ParseSQLSchema(*schemaFile)
		if err != nil {
			logger.Fatal("Error parsing SQL schema: %v", err)
		}
	}

	logger.Info("Schema parsed successfully: %d table(s) found", len(schemaData.Tables))
	for _, table := range schemaData.Tables {
		logger.Debug("Table '%s': %d fields", table.Name, len(table.Fields))
	}

	// Generate fake data with optional AI enhancement and performance optimizations
	var generatedFiles []string
	
	logger.Step("3", "Generating fake data")
	
	// Configure performance settings
	var performanceConfig generator.PerformanceConfig
	if *enablePerf {
		performanceConfig = generator.DefaultPerformanceConfig()
		if *workers > 0 {
			performanceConfig.WorkerPoolSize = *workers
		}
		if *batchSize > 0 {
			performanceConfig.BatchSize = *batchSize
		}
		logger.Debug("Performance optimizations enabled: workers=%d, batch=%d, parallel=%v, cache=%v", 
			performanceConfig.WorkerPoolSize, performanceConfig.BatchSize, 
			performanceConfig.EnableParallel, performanceConfig.CacheFieldInference)
	} else {
		// Use non-optimized settings for compatibility
		performanceConfig = generator.PerformanceConfig{
			EnableParallel:      false,
			WorkerPoolSize:      1,
			BatchSize:          100,
			PreallocateMemory:   false,
			CacheFieldInference: false,
		}
	}
	
	if *enableAI {
		aiStatus := getAIStatus()
		logger.Info("AI-enhanced mode enabled (OpenAI API: %s)", aiStatus)
		if aiStatus == "Available" {
			logger.Debug("Using OpenAI API for intelligent field inference")
		}
		
		logger.Time("AI-enhanced data generation", func() {
			generatedFiles, err = generator.GenerateWithAIAndFormat(&schemaData, *numRows, *outputFile, *outputFormat)
		})
	} else {
		logger.Debug("Using standard field inference")
		if *enablePerf {
			logger.Time("Optimized data generation", func() {
				generatedFiles, err = generator.GenerateDataFilesOptimized(schemaData, *numRows, *outputFile, 
					generator.FormatCSV, performanceConfig)
				if *outputFormat == "json" {
					generatedFiles, err = generator.GenerateDataFilesOptimized(schemaData, *numRows, *outputFile, 
						generator.FormatJSON, performanceConfig)
				}
			})
		} else {
			logger.Time("Standard data generation", func() {
				generatedFiles, err = generator.GenerateWithFormat(&schemaData, *numRows, *outputFile, *outputFormat)
			})
		}
	}

	if err != nil {
		logger.Fatal("Error generating data: %v", err)
	}

	logger.Step("4", "Writing output files")

	// Display results
	if len(generatedFiles) == 1 {
		logger.Info("Fake data generated and written to: %s", generatedFiles[0])
		fmt.Printf("Fake data generated and written to: %s\n", generatedFiles[0])
	} else {
		logger.Info("Fake data generated and written to %d files", len(generatedFiles))
		fmt.Printf("Fake data generated and written to %d files:\n", len(generatedFiles))
		for _, file := range generatedFiles {
			logger.Debug("Generated file: %s", file)
			fmt.Printf("  - %s\n", file)
		}
	}
	
	logger.Info("Data generation completed successfully")
}

func getAIStatus() string {
	if os.Getenv("OPENAI_API_KEY") != "" {
		return "Available"
	}
	return "Not configured (set OPENAI_API_KEY)"
}
