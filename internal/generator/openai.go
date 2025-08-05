package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// OpenAIConfig holds configuration for OpenAI API integration
type OpenAIConfig struct {
	APIKey      string
	Model       string
	MaxTokens   int
	Temperature float64
	Enabled     bool
}

// OpenAIFieldInference handles AI-powered field type inference
type OpenAIFieldInference struct {
	config OpenAIConfig
	client *http.Client
}

// NewOpenAIFieldInference creates a new AI-powered field inference client
func NewOpenAIFieldInference() *OpenAIFieldInference {
	apiKey := os.Getenv("OPENAI_API_KEY")
	enabled := apiKey != ""
	
	return &OpenAIFieldInference{
		config: OpenAIConfig{
			APIKey:      apiKey,
			Model:       "gpt-3.5-turbo",
			MaxTokens:   100,
			Temperature: 0.1, // Low temperature for consistent results
			Enabled:     enabled,
		},
		client: &http.Client{},
	}
}

// OpenAIRequest represents the request structure for OpenAI API
type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents the response from OpenAI API
type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
	Error   *APIError `json:"error,omitempty"`
}

// Choice represents a choice in the OpenAI response
type Choice struct {
	Message Message `json:"message"`
}

// APIError represents an error from the OpenAI API
type APIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}

// InferFieldTypeWithAI uses OpenAI to intelligently infer field types
func (ai *OpenAIFieldInference) InferFieldTypeWithAI(fieldName, fieldType, tableName string, sampleData []string) (string, float64, error) {
	if !ai.config.Enabled {
		return "", 0.0, fmt.Errorf("OpenAI API not configured")
	}

	prompt := ai.buildInferencePrompt(fieldName, fieldType, tableName, sampleData)
	
	response, err := ai.callOpenAI(prompt)
	if err != nil {
		return "", 0.0, err
	}
	
	return ai.parseInferenceResponse(response)
}

// buildInferencePrompt creates a prompt for field type inference
func (ai *OpenAIFieldInference) buildInferencePrompt(fieldName, fieldType, tableName string, sampleData []string) string {
	prompt := fmt.Sprintf(`You are an expert database schema analyzer. Given the following field information, determine the most appropriate data type for generating realistic fake data.

Field Name: %s
Declared Type: %s
Table Name: %s`, fieldName, fieldType, tableName)

	if len(sampleData) > 0 {
		prompt += fmt.Sprintf("\nSample Data: %s", strings.Join(sampleData, ", "))
	}

	prompt += `

Available data types:
- email: Email addresses
- name: Full names
- firstname: First names only
- lastname: Last names only  
- phone: Phone numbers
- address: Street addresses
- city: City names
- state: State/province codes
- zipcode: ZIP/postal codes
- country: Country names
- company: Company names
- uuid: Unique identifiers
- date: Dates (YYYY-MM-DD)
- datetime: Date with time
- price: Monetary amounts
- boolean: True/false values
- text: Long text content
- url: Web URLs
- image: Image URLs
- jobtitle: Job titles
- department: Department names
- skill: Skills/technologies
- color: Color names
- product: Product names
- brand: Brand names
- username: Usernames
- password: Passwords
- ipaddress: IP addresses
- macaddress: MAC addresses
- creditcard: Credit card numbers
- bankaccount: Bank account numbers
- ssn: Social security numbers
- license: License numbers
- version: Software versions
- status: Status values
- priority: Priority levels
- duration: Time durations
- filename: File names
- hashtag: Social media hashtags
- longitude: Geographic longitude
- latitude: Geographic latitude
- temperature: Temperature values
- weight: Weight measurements
- height: Height measurements
- age: Age values
- gender: Gender values
- category: Categories/classifications
- int: Integer numbers
- float: Decimal numbers
- string: Generic text

Respond with ONLY the data type name and a confidence score (0.0-1.0) in this format:
datatype:confidence

Example: email:0.95`

	return prompt
}

// callOpenAI makes a request to the OpenAI API
func (ai *OpenAIFieldInference) callOpenAI(prompt string) (string, error) {
	request := OpenAIRequest{
		Model:       ai.config.Model,
		MaxTokens:   ai.config.MaxTokens,
		Temperature: ai.config.Temperature,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ai.config.APIKey)

	resp, err := ai.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response OpenAIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if response.Error != nil {
		return "", fmt.Errorf("OpenAI API error: %s", response.Error.Message)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenAI")
	}

	return response.Choices[0].Message.Content, nil
}

// parseInferenceResponse parses the AI response to extract data type and confidence
func (ai *OpenAIFieldInference) parseInferenceResponse(response string) (string, float64, error) {
	response = strings.TrimSpace(response)
	parts := strings.Split(response, ":")
	
	if len(parts) != 2 {
		return "", 0.0, fmt.Errorf("invalid response format: %s", response)
	}
	
	dataType := strings.TrimSpace(parts[0])
	confidenceStr := strings.TrimSpace(parts[1])
	
	var confidence float64
	if _, err := fmt.Sscanf(confidenceStr, "%f", &confidence); err != nil {
		return "", 0.0, fmt.Errorf("invalid confidence score: %s", confidenceStr)
	}
	
	return dataType, confidence, nil
}

// GenerateSchemaDescription uses AI to generate human-readable schema descriptions
func (ai *OpenAIFieldInference) GenerateSchemaDescription(tableName string, fields []string) (string, error) {
	if !ai.config.Enabled {
		return "", fmt.Errorf("OpenAI API not configured")
	}

	prompt := fmt.Sprintf(`Analyze this database table schema and provide a brief, professional description of what this table represents and its purpose.

Table Name: %s
Fields: %s

Provide a 1-2 sentence description focusing on the business purpose and data it contains.`, tableName, strings.Join(fields, ", "))

	response, err := ai.callOpenAI(prompt)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(response), nil
}

// SuggestRelatedFields uses AI to suggest fields that might be missing from a table
func (ai *OpenAIFieldInference) SuggestRelatedFields(tableName string, existingFields []string) ([]string, error) {
	if !ai.config.Enabled {
		return nil, fmt.Errorf("OpenAI API not configured")
	}

	prompt := fmt.Sprintf(`Given this database table, suggest 3-5 additional fields that would commonly be found in this type of table.

Table Name: %s
Existing Fields: %s

Respond with ONLY field names, one per line, without explanations.`, tableName, strings.Join(existingFields, ", "))

	response, err := ai.callOpenAI(prompt)
	if err != nil {
		return nil, err
	}

	suggestions := strings.Split(strings.TrimSpace(response), "\n")
	var cleanSuggestions []string
	for _, suggestion := range suggestions {
		if trimmed := strings.TrimSpace(suggestion); trimmed != "" {
			cleanSuggestions = append(cleanSuggestions, trimmed)
		}
	}

	return cleanSuggestions, nil
}
