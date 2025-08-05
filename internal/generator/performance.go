package generator

import (
	"go-fake/internal/schema"
	"go-fake/pkg/logger"
	"runtime"
	"sync"
)

// PerformanceConfig contains settings for performance optimizations
type PerformanceConfig struct {
	EnableParallel     bool // Enable parallel table generation
	WorkerPoolSize     int  // Number of workers for parallel generation
	BatchSize          int  // Number of rows to generate in each batch
	PreallocateMemory  bool // Pre-allocate slices and maps for better memory usage
	CacheFieldInference bool // Cache field inference results
}

// DefaultPerformanceConfig returns optimized default settings
func DefaultPerformanceConfig() PerformanceConfig {
	numCPU := runtime.NumCPU()
	return PerformanceConfig{
		EnableParallel:      true,
		WorkerPoolSize:      numCPU,
		BatchSize:          1000, // Generate rows in batches of 1000
		PreallocateMemory:   true,
		CacheFieldInference: true,
	}
}

// FieldInferenceCache caches field inference results to avoid repeated computation
type FieldInferenceCache struct {
	mutex sync.RWMutex
	cache map[string]string // field_name -> inferred_type
}

// NewFieldInferenceCache creates a new field inference cache
func NewFieldInferenceCache() *FieldInferenceCache {
	return &FieldInferenceCache{
		cache: make(map[string]string),
	}
}

// Get retrieves cached inference result
func (c *FieldInferenceCache) Get(fieldName string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	result, exists := c.cache[fieldName]
	return result, exists
}

// Set stores inference result in cache
func (c *FieldInferenceCache) Set(fieldName, inferredType string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[fieldName] = inferredType
}

// ParallelTableGenerator handles parallel generation of multiple tables
type ParallelTableGenerator struct {
	config         PerformanceConfig
	inferenceCache *FieldInferenceCache
	fieldInference *FieldTypeInference
}

// NewParallelTableGenerator creates a new parallel table generator
func NewParallelTableGenerator(config PerformanceConfig) *ParallelTableGenerator {
	return &ParallelTableGenerator{
		config:         config,
		inferenceCache: NewFieldInferenceCache(),
		fieldInference: NewFieldTypeInference(),
	}
}

// GenerateTablesParallel generates multiple tables in parallel
func (ptg *ParallelTableGenerator) GenerateTablesParallel(tables []schema.Table, numRows int, relData *RelationshipData, relationships []schema.Relationship) {
	// Separate independent and dependent tables
	var independentTables, dependentTables []schema.Table
	
	for _, table := range tables {
		if !hasReferences(table.Fields) {
			independentTables = append(independentTables, table)
		} else {
			dependentTables = append(dependentTables, table)
		}
	}
	
	// Generate independent tables in parallel
	if len(independentTables) > 0 {
		logger.Debug("Generating %d independent tables in parallel", len(independentTables))
		ptg.generateTablesBatch(independentTables, numRows, relData, relationships)
	}
	
	// Generate dependent tables in parallel (but after independent ones)
	if len(dependentTables) > 0 {
		logger.Debug("Generating %d dependent tables in parallel", len(dependentTables))
		ptg.generateTablesBatch(dependentTables, numRows, relData, relationships)
	}
}

// generateTablesBatch generates a batch of tables in parallel
func (ptg *ParallelTableGenerator) generateTablesBatch(tables []schema.Table, numRows int, relData *RelationshipData, relationships []schema.Relationship) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, ptg.config.WorkerPoolSize)
	
	for _, table := range tables {
		wg.Add(1)
		go func(t schema.Table) {
			defer wg.Done()
			semaphore <- struct{}{} // Acquire worker
			defer func() { <-semaphore }() // Release worker
			
			logger.Debug("Generating data for table: %s", t.Name)
			data := ptg.generateTableDataOptimized(t.Fields, numRows, t.Name, relData, relationships)
			
			// Thread-safe update of relData
			relData.mutex.Lock()
			relData.TableData[t.Name] = data
			populateReferences(t.Name, t.Fields, data, relData)
			relData.mutex.Unlock()
		}(table)
	}
	
	wg.Wait()
}

// generateTableDataOptimized generates table data with performance optimizations
func (ptg *ParallelTableGenerator) generateTableDataOptimized(fields []schema.Field, numRows int, tableName string, relData *RelationshipData, relationships []schema.Relationship) []map[string]interface{} {
	// Pre-allocate slice with exact capacity
	rows := make([]map[string]interface{}, 0, numRows)
	
	// Cache field inference for all fields once
	fieldTypes := make(map[string]string, len(fields))
	if ptg.config.CacheFieldInference {
		for _, field := range fields {
			if cachedType, exists := ptg.inferenceCache.Get(field.Name); exists {
				fieldTypes[field.Name] = cachedType
			} else {
				inferredType := ptg.fieldInference.InferFieldType(field)
				fieldTypes[field.Name] = inferredType
				ptg.inferenceCache.Set(field.Name, inferredType)
			}
		}
	}
	
	// Generate rows in batches for better memory usage
	batchSize := ptg.config.BatchSize
	if numRows < batchSize {
		batchSize = numRows
	}
	
	for batchStart := 0; batchStart < numRows; batchStart += batchSize {
		batchEnd := batchStart + batchSize
		if batchEnd > numRows {
			batchEnd = numRows
		}
		
		// Generate batch of rows
		batch := ptg.generateRowBatch(fields, batchStart, batchEnd, tableName, relData, fieldTypes)
		rows = append(rows, batch...)
	}
	
	return rows
}

// generateRowBatch generates a batch of rows efficiently
func (ptg *ParallelTableGenerator) generateRowBatch(fields []schema.Field, startRow, endRow int, tableName string, relData *RelationshipData, fieldTypes map[string]string) []map[string]interface{} {
	batchSize := endRow - startRow
	batch := make([]map[string]interface{}, 0, batchSize)
	
	for i := startRow; i < endRow; i++ {
		row := make(map[string]interface{}, len(fields))
		
		for _, field := range fields {
			// Use cached inference type if available
			var value interface{}
			if ptg.config.CacheFieldInference && fieldTypes[field.Name] != "" {
				value = ptg.fieldInference.GenerateValueByType(fieldTypes[field.Name], field.Name)
			} else {
				value = generateConstrainedValue(field, relData, tableName, i)
			}
			row[field.Name] = value
		}
		
		batch = append(batch, row)
	}
	
	return batch
}
