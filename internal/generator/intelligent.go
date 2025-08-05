package generator

import (
	"go-fake/internal/schema"
	"go-fake/pkg/faker"
	"math/rand/v2"
	"regexp"
	"strings"
)

// FieldTypeInference handles intelligent field type detection and value generation
type FieldTypeInference struct {
	patterns map[string][]string // type -> patterns
	semantic map[string]string   // semantic name -> type
}

// NewFieldTypeInference creates a new intelligent field type detector
func NewFieldTypeInference() *FieldTypeInference {
	return &FieldTypeInference{
		patterns: map[string][]string{
			// Email patterns
			"email": {
				"email", "e_mail", "email_address", "e_mail_address", 
				"contact_email", "user_email", "customer_email", "work_email",
			},
			
			// Name patterns
			"name": {
				"name", "full_name", "fullname", "display_name", "user_name", 
				"customer_name", "client_name", "person_name", "contact_name",
			},
			
			// First name patterns
			"firstname": {
				"first_name", "firstname", "fname", "given_name", "forename",
				"first", "christian_name",
			},
			
			// Last name patterns
			"lastname": {
				"last_name", "lastname", "lname", "surname", "family_name", 
				"last", "sur_name",
			},
			
			// Phone patterns
			"phone": {
				"phone", "phone_number", "phonenumber", "mobile", "mobile_number",
				"cell", "cell_phone", "telephone", "contact_number", "contact_phone",
				"work_phone", "home_phone", "fax", "fax_number",
			},
			
			// Address patterns
			"address": {
				"address", "street_address", "street", "address_line", "addr",
				"home_address", "work_address", "mailing_address", "shipping_address",
				"billing_address", "physical_address",
			},
			
			// City patterns
			"city": {
				"city", "town", "municipality", "locality", "place", "city_name",
				"hometown", "residence_city",
			},
			
			// State patterns
			"state": {
				"state", "province", "region", "territory", "state_code", 
				"province_code", "state_name", "province_name",
			},
			
			// ZIP code patterns
			"zipcode": {
				"zip", "zipcode", "zip_code", "postal_code", "postcode", 
				"postal", "zip_postal", "post_code",
			},
			
			// Country patterns
			"country": {
				"country", "country_name", "country_code", "nation", "nationality",
				"country_iso", "country_alpha", "homeland",
			},
			
			// Company patterns
			"company": {
				"company", "company_name", "organization", "org", "business",
				"corporation", "corp", "firm", "enterprise", "employer",
				"organization_name", "business_name",
			},
			
			// UUID patterns
			"uuid": {
				"uuid", "guid", "id", "identifier", "unique_id", "key", 
				"primary_key", "ref_id", "reference_id", "external_id",
			},
			
			// Date patterns
			"date": {
				"date", "created_date", "updated_date", "birth_date", "birthdate",
				"start_date", "end_date", "due_date", "expiry_date", "expiration_date",
				"registration_date", "join_date", "hired_date",
			},
			
			// DateTime/Timestamp patterns
			"datetime": {
				"datetime", "timestamp", "created_at", "updated_at", "modified_at",
				"last_login", "last_seen", "login_time", "access_time", "event_time",
				"created_on", "updated_on", "processed_at",
			},
			
			// Price/Money patterns
			"price": {
				"price", "cost", "amount", "fee", "charge", "rate", "salary",
				"wage", "payment", "total", "subtotal", "tax", "discount",
				"revenue", "income", "expense", "budget", "balance",
			},
			
			// Boolean patterns
			"boolean": {
				"active", "enabled", "disabled", "verified", "confirmed", "approved",
				"published", "deleted", "archived", "featured", "premium", "paid",
				"completed", "finished", "closed", "open", "available", "visible",
				"public", "private", "is_active", "is_enabled", "is_deleted",
			},
			
			// Description/Text patterns
			"text": {
				"description", "bio", "biography", "summary", "notes", "comments",
				"details", "content", "body", "message", "review", "feedback",
				"about", "info", "information", "remarks", "observations",
			},
			
			// URL patterns
			"url": {
				"url", "website", "link", "homepage", "site", "web_address",
				"web_url", "site_url", "profile_url", "image_url", "avatar_url",
			},
			
			// Image patterns
			"image": {
				"image", "photo", "picture", "avatar", "thumbnail", "logo",
				"banner", "icon", "profile_image", "profile_picture", "img",
			},
			
			// Age patterns
			"age": {
				"age", "years_old", "birth_year", "year_of_birth", "age_years",
			},
			
			// Job title patterns (should come before gender patterns for precedence)
			"jobtitle": {
				"job_title", "job", "position", "role", "occupation", "profession",
				"title", "job_position", "work_title", "career", "designation",
			},
			
			// Gender patterns (removed generic "title" to avoid conflict)
			"gender": {
				"gender", "sex", "mr_mrs", "salutation",
			},
			
			// Department patterns
			"department": {
				"department", "dept", "division", "team", "unit", "section",
				"department_name", "work_department", "business_unit",
			},
			
			// Skill patterns
			"skill": {
				"skill", "skills", "technology", "technologies", "expertise", 
				"competency", "competencies", "capability", "abilities", "talent",
			},
			
			// Color patterns
			"color": {
				"color", "colour", "hue", "shade", "tint", "pigment",
				"primary_color", "background_color", "text_color",
			},
			
			// Product patterns
			"product": {
				"product", "product_name", "item", "item_name", "merchandise",
				"goods", "article", "commodity", "sku", "model",
			},
			
			// Brand patterns
			"brand": {
				"brand", "brand_name", "manufacturer", "make", "label",
				"trademark", "vendor", "supplier", "producer",
			},
			
			// Username patterns
			"username": {
				"username", "user_name", "login", "handle", "nickname", "nick",
				"screen_name", "display_name", "alias", "login_name",
			},
			
			// Password patterns
			"password": {
				"password", "passwd", "pass", "pwd", "secret", "pin",
				"passcode", "access_code", "security_code",
			},
			
			// IP Address patterns
			"ipaddress": {
				"ip", "ip_address", "ipv4", "ipv6", "host", "server_ip",
				"client_ip", "remote_ip", "local_ip", "network_address",
			},
			
			// MAC Address patterns
			"macaddress": {
				"mac", "mac_address", "hardware_address", "physical_address",
				"ethernet_address", "wifi_mac", "device_mac",
			},
			
			// Credit Card patterns
			"creditcard": {
				"credit_card", "creditcard", "card_number", "cc_number",
				"payment_card", "debit_card", "card", "cc",
			},
			
			// Bank Account patterns
			"bankaccount": {
				"bank_account", "routing_number", "account_number",
				"iban", "swift", "bic", "sort_code", "account_no",
			},
			
			// Social Security patterns
			"ssn": {
				"ssn", "social_security", "social_security_number", "tax_id",
				"national_id", "personal_id", "citizen_id",
			},
			
			// License patterns
			"license": {
				"license", "licence", "license_number", "permit", "certificate",
				"registration", "license_plate", "driver_license",
			},
			
			// Version patterns
			"version": {
				"version", "ver", "release", "build", "revision", "v",
				"software_version", "app_version", "api_version",
			},
			
			// Status patterns
			"status": {
				"status", "state", "condition", "stage", "phase", "mode",
				"current_status", "order_status", "payment_status",
			},
			
			// Priority patterns
			"priority": {
				"priority", "importance", "urgency", "level", "rank", "grade",
				"priority_level", "severity", "criticality",
			},
			
			// Duration patterns
			"duration": {
				"duration", "length", "time", "period", "interval", "span",
				"elapsed_time", "runtime", "execution_time",
			},
			
			// File patterns
			"filename": {
				"file", "filename", "file_name", "document", "attachment",
				"upload", "media", "resource", "asset", "path",
			},
			
			// Hashtag patterns
			"hashtag": {
				"hashtag", "tag", "tags", "keyword", "keywords", "label",
				"category_tag", "search_tag", "topic",
			},
			
			// Longitude patterns
			"longitude": {
				"longitude", "lng", "lon", "long", "x_coordinate", "east_west",
			},
			
			// Latitude patterns  
			"latitude": {
				"latitude", "lat", "y_coordinate", "north_south",
			},
			
			// Temperature patterns
			"temperature": {
				"temperature", "temp", "celsius", "fahrenheit", "kelvin",
				"degrees", "thermal", "heat",
			},
			
			// Weight patterns
			"weight": {
				"weight", "mass", "kg", "kilogram", "pound", "lb", "gram",
				"ounce", "ton", "stone",
			},
			
			// Height patterns
			"height": {
				"height", "tall", "stature", "elevation", "altitude", "length",
				"inches", "feet", "cm", "centimeter", "meter",
			},
			
			// Category patterns
			"category": {
				"category", "type", "kind", "classification", "group", "class",
				"tag", "label", "status", "role", "department", "division",
			},
		},
		
		semantic: map[string]string{
			// Common business entities
			"customer":    "name",
			"user":        "name", 
			"client":      "name",
			"employee":    "name",
			"person":      "name",
			"contact":     "name",
			"vendor":      "company",
			"supplier":    "company",
			"manufacturer": "company",
			"brand":       "company",
			
			// Time-related
			"created":     "datetime",
			"updated":     "datetime", 
			"modified":    "datetime",
			"deleted":     "datetime",
			"published":   "datetime",
			"expired":     "datetime",
			"started":     "datetime",
			"finished":    "datetime",
			"birth":       "date",
			"hire":        "date",
			"join":        "date",
			
			// Financial
			"salary":      "price",
			"wage":        "price",
			"cost":        "price",
			"fee":         "price",
			"amount":      "price",
			"total":       "price",
			"balance":     "price",
			"budget":      "price",
			
			// Contact info (removing duplicate "contact" entry)
			"mobile":      "phone",
			"tel":         "phone",
			"telephone":   "phone",
			"cell":        "phone",
		},
	}
}

// InferFieldType intelligently determines the most appropriate field type
func (f *FieldTypeInference) InferFieldType(field schema.Field) string {
	fieldName := strings.ToLower(field.Name)
	fieldType := strings.ToLower(field.Type)
	
	// First, check for direct SQL type mappings
	if sqlType := f.mapSQLType(fieldType); sqlType != "" {
		return sqlType
	}
	
	// Check for exact pattern matches in field name
	for targetType, patterns := range f.patterns {
		for _, pattern := range patterns {
			if fieldName == pattern {
				return targetType
			}
		}
	}
	
	// Check for partial pattern matches (contains)
	for targetType, patterns := range f.patterns {
		for _, pattern := range patterns {
			if strings.Contains(fieldName, pattern) {
				return targetType
			}
		}
	}
	
	// Check semantic understanding (word parts)
	for semantic, targetType := range f.semantic {
		if strings.Contains(fieldName, semantic) {
			return targetType
		}
	}
	
	// Advanced pattern matching with regex
	if inferredType := f.regexPatternMatch(fieldName); inferredType != "" {
		return inferredType
	}
	
	// Context-based inference (suffix/prefix patterns)
	if inferredType := f.contextualInference(fieldName); inferredType != "" {
		return inferredType
	}
	
	// Default fallback based on common naming conventions
	return f.defaultInference(fieldName, fieldType)
}

// mapSQLType handles SQL type mappings
func (f *FieldTypeInference) mapSQLType(sqlType string) string {
	sqlType = strings.ToUpper(sqlType)
	
	switch {
	case strings.Contains(sqlType, "SERIAL"), strings.Contains(sqlType, "AUTO_INCREMENT"):
		return "int"
	case strings.Contains(sqlType, "INTEGER"), strings.Contains(sqlType, "INT"), strings.Contains(sqlType, "BIGINT"):
		return "int"
	// Don't immediately return "string" for generic types - let intelligent inference handle it
	case strings.Contains(sqlType, "VARCHAR"), strings.Contains(sqlType, "TEXT"), strings.Contains(sqlType, "CHAR"):
		return "" // Let intelligent inference determine the actual type
	case sqlType == "STRING": // Only for exact "STRING" match, let intelligent inference handle generic string
		return ""
	case strings.Contains(sqlType, "DECIMAL"), strings.Contains(sqlType, "NUMERIC"), strings.Contains(sqlType, "FLOAT"), strings.Contains(sqlType, "DOUBLE"):
		return "float"
	case strings.Contains(sqlType, "BOOLEAN"), strings.Contains(sqlType, "BOOL"):
		return "boolean"
	case strings.Contains(sqlType, "DATE") && !strings.Contains(sqlType, "TIME"):
		return "date"
	case strings.Contains(sqlType, "TIMESTAMP"), strings.Contains(sqlType, "DATETIME"):
		return "datetime"
	case strings.Contains(sqlType, "UUID"), strings.Contains(sqlType, "GUID"):
		return "uuid"
	case strings.Contains(sqlType, "JSON"):
		return "text"
	case strings.Contains(sqlType, "BINARY"), strings.Contains(sqlType, "BLOB"):
		return "string"
	}
	
	return ""
}

// regexPatternMatch uses regex patterns for advanced matching
func (f *FieldTypeInference) regexPatternMatch(fieldName string) string {
	patterns := map[string]*regexp.Regexp{
		"email":    regexp.MustCompile(`^.*e?mail.*$`),
		"phone":    regexp.MustCompile(`^.*(phone|tel|mobile|cell|contact).*$`),
		"date":     regexp.MustCompile(`^.*(date|day|month|year).*$`),
		"datetime": regexp.MustCompile(`^.*(timestamp|datetime|time|at|on).*$`),
		"price":    regexp.MustCompile(`^.*(price|cost|amount|fee|salary|wage|budget|total).*$`),
		"boolean":  regexp.MustCompile(`^(is_|has_|can_|should_|will_|was_).*$`),
		"url":      regexp.MustCompile(`^.*(url|link|site|web|http).*$`),
		"image":    regexp.MustCompile(`^.*(image|img|photo|picture|avatar|thumbnail).*$`),
		"uuid":     regexp.MustCompile(`^.*(id|key|uuid|guid)$`),
	}
	
	for fieldType, pattern := range patterns {
		if pattern.MatchString(fieldName) {
			return fieldType
		}
	}
	
	return ""
}

// contextualInference uses contextual clues like suffixes and prefixes
func (f *FieldTypeInference) contextualInference(fieldName string) string {
	// Suffix-based inference
	suffixes := map[string]string{
		"_id":          "uuid",
		"_key":         "uuid", 
		"_at":          "datetime",
		"_on":          "datetime",
		"_date":        "date",
		"_time":        "datetime",
		"_url":         "url",
		"_link":        "url",
		"_email":       "email",
		"_phone":       "phone",
		"_address":     "address",
		"_city":        "city",
		"_state":       "state",
		"_zip":         "zipcode",
		"_country":     "country",
		"_price":       "price",
		"_cost":        "price",
		"_amount":      "price",
		"_total":       "price",
		"_count":       "int",
		"_number":      "int",
		"_age":         "age",
		"_year":        "int",
		"_status":      "category",
		"_type":        "category",
		"_category":    "category",
		"_description": "text",
		"_bio":         "text",
		"_notes":       "text",
		"_comments":    "text",
	}
	
	for suffix, fieldType := range suffixes {
		if strings.HasSuffix(fieldName, suffix) {
			return fieldType
		}
	}
	
	// Prefix-based inference
	prefixes := map[string]string{
		"is_":       "boolean",
		"has_":      "boolean",
		"can_":      "boolean", 
		"should_":   "boolean",
		"will_":     "boolean",
		"was_":      "boolean",
		"user_":     "name",
		"customer_": "name",
		"client_":   "name",
		"contact_":  "phone",
		"home_":     "address",
		"work_":     "address",
		"billing_":  "address",
		"shipping_": "address",
	}
	
	for prefix, fieldType := range prefixes {
		if strings.HasPrefix(fieldName, prefix) {
			return fieldType
		}
	}
	
	return ""
}

// defaultInference provides fallback inference based on common patterns
func (f *FieldTypeInference) defaultInference(fieldName, fieldType string) string {
	// Length-based heuristics
	if len(fieldName) <= 3 {
		if strings.Contains(fieldName, "id") {
			return "uuid"
		}
		return "string"
	}
	
	// Common single-word patterns
	singleWords := map[string]string{
		"name":        "name",
		"title":       "string",
		"description": "text",
		"content":     "text",
		"message":     "text",
		"comment":     "text",
		"note":        "text",
		"summary":     "text",
		"bio":         "text",
		"about":       "text",
		"details":     "text",
		"info":        "text",
		"data":        "text",
		"value":       "string",
		"code":        "string",
		"key":         "uuid",
		"token":       "uuid",
		"hash":        "string",
		"slug":        "string",
		"tag":         "string",
		"label":       "string",
		"status":      "category",
		"role":        "category",
		"level":       "int",
		"rank":        "int",
		"order":       "int",
		"position":    "int",
		"index":       "int",
		"count":       "int",
		"quantity":    "int",
		"size":        "int",
		"weight":      "float",
		"height":      "float",
		"width":       "float",
		"length":      "float",
		"score":       "float",
		"rating":      "float",
		"percentage":  "float",
		"ratio":       "float",
	}
	
	if singleType, exists := singleWords[fieldName]; exists {
		return singleType
	}
	
	// If all else fails, use the original type or default to string
	if fieldType != "" && fieldType != "string" {
		return fieldType
	}
	
	return "string"
}

// GenerateIntelligentValue generates a value using the intelligent field type inference
func (f *FieldTypeInference) GenerateIntelligentValue(field schema.Field) interface{} {
	inferredType := f.InferFieldType(field)
	
	// Handle constraints if present
	if field.Constraints != nil {
		return f.generateConstrainedValue(field, inferredType)
	}
	
	return f.generateValueByType(inferredType, field.Name)
}

// generateConstrainedValue generates values respecting field constraints
func (f *FieldTypeInference) generateConstrainedValue(field schema.Field, inferredType string) interface{} {
	constraints := field.Constraints
	
	// Handle min/max for numeric types
	if inferredType == "int" || inferredType == "age" {
		min := 1
		max := 1000
		
		if constraints.MinValue != nil {
			min = *constraints.MinValue
		}
		if constraints.MaxValue != nil {
			max = *constraints.MaxValue
		}
		
		if inferredType == "age" && constraints.MinValue == nil && constraints.MaxValue == nil {
			min = 18
			max = 80
		}
		
		return rand.IntN(max-min+1) + min
	}
	
	if inferredType == "float" || inferredType == "price" {
		// For price fields, use more realistic ranges
		if inferredType == "price" && constraints.MinValue == nil && constraints.MaxValue == nil {
			return float64(rand.IntN(100000)) / 100.0 // $0.00 to $1000.00
		}
		
		min := 0.0
		max := 1000.0
		
		if constraints.MinValue != nil {
			min = float64(*constraints.MinValue)
		}
		if constraints.MaxValue != nil {
			max = float64(*constraints.MaxValue)
		}
		
		return min + (max-min)*rand.Float64()
	}
	
	// For other types, generate normally
	return f.generateValueByType(inferredType, field.Name)
}

// generateValueByType generates values based on the inferred type
func (f *FieldTypeInference) generateValueByType(fieldType, fieldName string) interface{} {
	switch fieldType {
	case "email":
		return faker.GenerateEmail()
	case "name":
		return faker.GenerateName()
	case "firstname":
		return faker.GenerateFirstName()
	case "lastname":
		return faker.GenerateLastName()
	case "phone":
		return faker.GeneratePhone()
	case "address":
		return faker.GenerateAddress()
	case "city":
		return faker.GenerateCity()
	case "state":
		return faker.GenerateState()
	case "zipcode":
		return faker.GenerateZipCode()
	case "country":
		return faker.GenerateCountry()
	case "company":
		return faker.GenerateCompany()
	case "uuid":
		return faker.GenerateUUID()
	case "date":
		return faker.GenerateDate()
	case "datetime":
		return faker.GenerateDateTime()
	case "price":
		return faker.GeneratePrice()
	case "boolean":
		return rand.IntN(2) == 1
	case "text":
		return faker.GenerateText()
	case "url":
		return faker.GenerateURL()
	case "image":
		return faker.GenerateImageURL()
	case "jobtitle":
		return faker.GenerateJobTitle()
	case "department":
		return faker.GenerateDepartment()
	case "skill":
		return faker.GenerateSkill()
	case "color":
		return faker.GenerateColor()
	case "product":
		return faker.GenerateProductName()
	case "brand":
		return faker.GenerateBrandName()
	case "username":
		return faker.GenerateUsername()
	case "password":
		return faker.GeneratePassword()
	case "ipaddress":
		return faker.GenerateIPAddress()
	case "macaddress":
		return faker.GenerateMACAddress()
	case "creditcard":
		return faker.GenerateCreditCard()
	case "bankaccount":
		return faker.GenerateBankAccount()
	case "ssn":
		return faker.GenerateSSN()
	case "license":
		return faker.GenerateLicense()
	case "version":
		return faker.GenerateVersion()
	case "status":
		return faker.GenerateStatus()
	case "priority":
		return faker.GeneratePriority()
	case "duration":
		return faker.GenerateDuration()
	case "filename":
		return faker.GenerateFilename()
	case "hashtag":
		return faker.GenerateHashtag()
	case "longitude":
		return faker.GenerateLongitude()
	case "latitude":
		return faker.GenerateLatitude()
	case "temperature":
		return faker.GenerateTemperature()
	case "weight":
		return faker.GenerateWeight()
	case "height":
		return faker.GenerateHeight()
	case "age":
		return rand.IntN(63) + 18 // 18-80 years old
	case "gender":
		return faker.GenerateGender()
	case "category":
		return faker.GenerateCategory()
	case "int", "integer":
		return rand.IntN(1000) + 1
	case "float":
		return rand.Float64() * 1000
	case "string":
		// Try to make a contextual guess based on field name
		return f.generateContextualString(fieldName)
	default:
		return faker.GenerateName()
	}
}

// generateContextualString generates contextual strings based on field name
func (f *FieldTypeInference) generateContextualString(fieldName string) string {
	fieldName = strings.ToLower(fieldName)
	
	// Context-based string generation
	if strings.Contains(fieldName, "title") {
		return faker.GenerateJobTitle()
	}
	if strings.Contains(fieldName, "department") {
		return faker.GenerateDepartment()
	}
	if strings.Contains(fieldName, "skill") || strings.Contains(fieldName, "technology") {
		return faker.GenerateSkill()
	}
	if strings.Contains(fieldName, "color") {
		return faker.GenerateColor()
	}
	if strings.Contains(fieldName, "product") {
		return faker.GenerateProductName()
	}
	if strings.Contains(fieldName, "brand") {
		return faker.GenerateBrandName()
	}
	
	// Default to name if no context match
	return faker.GenerateName()
}
