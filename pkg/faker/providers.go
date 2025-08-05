package faker

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
)

var firstNames = []string{"John", "Jane", "Alice", "Bob", "Charlie", "Diana", "Emma", "Liam", "Olivia", "Noah", "Ava", "William", "Sophia", "James", "Isabella"}
var lastNames = []string{"Smith", "Doe", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Hernandez"}
var domains = []string{"example.com", "test.com", "demo.com", "sample.com", "fake.org", "placeholder.net"}
var companies = []string{"TechCorp", "DataSoft", "CloudWorks", "InnovateLab", "FutureSystem", "DigitalEdge"}
var streetNames = []string{"Main St", "Oak Ave", "First St", "Second Ave", "Park Rd", "Elm St", "Maple Ave", "Cedar Ln"}
var cities = []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia", "San Antonio", "San Diego", "Dallas", "San Jose"}
var states = []string{"CA", "TX", "FL", "NY", "PA", "IL", "OH", "GA", "NC", "MI"}
var countries = []string{"United States", "Canada", "United Kingdom", "Germany", "France", "Australia", "Japan", "Brazil", "India", "China"}
var jobTitles = []string{"Software Engineer", "Product Manager", "Data Scientist", "Designer", "Marketing Manager", "Sales Representative", "Accountant", "HR Manager", "Operations Manager", "Customer Success Manager"}
var departments = []string{"Engineering", "Marketing", "Sales", "Human Resources", "Finance", "Operations", "Customer Support", "Product", "Design", "Legal"}
var skills = []string{"JavaScript", "Python", "React", "Node.js", "SQL", "Machine Learning", "Project Management", "Communication", "Leadership", "Problem Solving"}
var colors = []string{"Red", "Blue", "Green", "Yellow", "Purple", "Orange", "Pink", "Brown", "Black", "White", "Gray", "Cyan", "Magenta"}
var productNames = []string{"SuperWidget", "ProGadget", "UltraDevice", "SmartTool", "MegaApp", "PowerBox", "FlexiPhone", "QuickPad", "EasyBook", "FastTrack"}
var brandNames = []string{"TechFlow", "DataDrive", "CloudFirst", "SmartEdge", "ProActive", "NextGen", "FastForward", "InnoCore", "FlexTech", "PowerPro"}
var categories = []string{"Technology", "Business", "Education", "Healthcare", "Finance", "Retail", "Entertainment", "Sports", "Travel", "Food"}
var genders = []string{"Male", "Female", "Non-binary", "Prefer not to say"}
var textSamples = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	"The quick brown fox jumps over the lazy dog.",
	"In a hole in the ground there lived a hobbit.",
	"It was the best of times, it was the worst of times.",
	"To be, or not to be, that is the question.",
}

func GenerateName() string {
	firstName := firstNames[rand.IntN(len(firstNames))]
	lastName := lastNames[rand.IntN(len(lastNames))]
	return firstName + " " + lastName
}

func GenerateFirstName() string {
	return firstNames[rand.IntN(len(firstNames))]
}

func GenerateLastName() string {
	return lastNames[rand.IntN(len(lastNames))]
}

func GenerateEmail() string {
	name := GenerateName()
	domain := domains[rand.IntN(len(domains))]
	return formatEmail(name, domain)
}

func GeneratePhone() string {
	return fmt.Sprintf("(%03d) %03d-%04d", rand.IntN(900)+100, rand.IntN(900)+100, rand.IntN(10000))
}

func GenerateCompany() string {
	return companies[rand.IntN(len(companies))]
}

func GenerateAddress() string {
	number := rand.IntN(9999) + 1
	street := streetNames[rand.IntN(len(streetNames))]
	return fmt.Sprintf("%d %s", number, street)
}

func GenerateCity() string {
	return cities[rand.IntN(len(cities))]
}

func GenerateState() string {
	return states[rand.IntN(len(states))]
}

func GenerateZipCode() string {
	return fmt.Sprintf("%05d", rand.IntN(100000))
}

func GenerateDate() string {
	start := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Now()
	delta := end.Unix() - start.Unix()
	randomTime := start.Unix() + rand.Int64N(delta)
	return time.Unix(randomTime, 0).Format("2006-01-02")
}

func GenerateDateTime() string {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Now()
	delta := end.Unix() - start.Unix()
	randomTime := start.Unix() + rand.Int64N(delta)
	return time.Unix(randomTime, 0).Format("2006-01-02 15:04:05")
}

func GenerateBool() string {
	if rand.IntN(2) == 0 {
		return "false"
	}
	return "true"
}

func GenerateFloat() string {
	return fmt.Sprintf("%.2f", rand.Float64()*1000)
}

func GeneratePrice() string {
	return fmt.Sprintf("%.2f", rand.Float64()*500+10)
}

func GenerateUUID() string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		rand.Uint32(),
		rand.Uint32()&0xffff,
		rand.Uint32()&0xffff,
		rand.Uint32()&0xffff,
		rand.Uint64()&0xffffffffffff)
}

func formatEmail(name, domain string) string {
	return strings.ReplaceAll(strings.ToLower(name), " ", ".") + "@" + domain
}

// New faker functions for intelligent field detection

func GenerateCountry() string {
	return countries[rand.IntN(len(countries))]
}

func GenerateText() string {
	return textSamples[rand.IntN(len(textSamples))]
}

func GenerateURL() string {
	domains := []string{"https://example.com", "https://test.org", "https://demo.net", "https://sample.io"}
	paths := []string{"/home", "/about", "/products", "/services", "/contact", "/blog", "/support"}
	domain := domains[rand.IntN(len(domains))]
	path := paths[rand.IntN(len(paths))]
	return domain + path
}

func GenerateImageURL() string {
	sizes := []string{"300x200", "400x300", "500x400", "600x400", "800x600"}
	categories := []string{"nature", "city", "people", "technology", "abstract"}
	size := sizes[rand.IntN(len(sizes))]
	category := categories[rand.IntN(len(categories))]
	return fmt.Sprintf("https://picsum.photos/%s?category=%s", size, category)
}

func GenerateGender() string {
	return genders[rand.IntN(len(genders))]
}

func GenerateCategory() string {
	return categories[rand.IntN(len(categories))]
}

func GenerateJobTitle() string {
	return jobTitles[rand.IntN(len(jobTitles))]
}

func GenerateDepartment() string {
	return departments[rand.IntN(len(departments))]
}

func GenerateSkill() string {
	return skills[rand.IntN(len(skills))]
}

func GenerateColor() string {
	return colors[rand.IntN(len(colors))]
}

func GenerateProductName() string {
	return productNames[rand.IntN(len(productNames))]
}

func GenerateBrandName() string {
	return brandNames[rand.IntN(len(brandNames))]
}

// Username generation
var usernameWords = []string{"cool", "fast", "smart", "super", "mega", "ultra", "pro", "elite", "master", "legend"}
var usernameEndings = []string{"123", "456", "789", "2024", "x", "pro", "dev", "code", "tech"}

func GenerateUsername() string {
	word := usernameWords[rand.IntN(len(usernameWords))]
	name := strings.ToLower(firstNames[rand.IntN(len(firstNames))])
	ending := usernameEndings[rand.IntN(len(usernameEndings))]
	return word + name + ending
}

// Password generation
func GeneratePassword() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	length := 8 + rand.IntN(8) // 8-15 characters
	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.IntN(len(chars))]
	}
	return string(password)
}

// IP Address generation
func GenerateIPAddress() string {
	return fmt.Sprintf("%d.%d.%d.%d", 
		rand.IntN(256), rand.IntN(256), rand.IntN(256), rand.IntN(256))
}

// MAC Address generation
func GenerateMACAddress() string {
	mac := make([]string, 6)
	for i := range mac {
		mac[i] = fmt.Sprintf("%02x", rand.IntN(256))
	}
	return strings.Join(mac, ":")
}

// Credit Card generation (fake numbers)
func GenerateCreditCard() string {
	// Generate a fake 16-digit credit card number
	digits := make([]string, 16)
	digits[0] = "4" // Visa starts with 4
	for i := 1; i < 16; i++ {
		digits[i] = fmt.Sprintf("%d", rand.IntN(10))
	}
	return strings.Join([]string{
		strings.Join(digits[0:4], ""),
		strings.Join(digits[4:8], ""),
		strings.Join(digits[8:12], ""),
		strings.Join(digits[12:16], ""),
	}, "-")
}

// Bank Account generation
func GenerateBankAccount() string {
	return fmt.Sprintf("%09d", rand.IntN(1000000000))
}

// SSN generation (fake format)
func GenerateSSN() string {
	return fmt.Sprintf("%03d-%02d-%04d", 
		rand.IntN(900)+100, rand.IntN(100), rand.IntN(10000))
}

// License number generation
func GenerateLicense() string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	license := ""
	for i := 0; i < 2; i++ {
		license += string(letters[rand.IntN(len(letters))])
	}
	license += fmt.Sprintf("%06d", rand.IntN(1000000))
	return license
}

// Version generation
var versionFormats = []string{"%d.%d.%d", "%d.%d", "%d.%d.%d-beta", "%d.%d.%d-alpha"}

func GenerateVersion() string {
	format := versionFormats[rand.IntN(len(versionFormats))]
	major := rand.IntN(10) + 1
	minor := rand.IntN(20)
	patch := rand.IntN(50)
	return fmt.Sprintf(format, major, minor, patch)
}

// Status generation
var statuses = []string{"Active", "Inactive", "Pending", "Approved", "Rejected", "In Progress", "Completed", "Cancelled", "On Hold", "Review"}

func GenerateStatus() string {
	return statuses[rand.IntN(len(statuses))]
}

// Priority generation
var priorities = []string{"Low", "Medium", "High", "Critical", "Urgent", "Normal", "Minor", "Major"}

func GeneratePriority() string {
	return priorities[rand.IntN(len(priorities))]
}

// Duration generation (in hours)
func GenerateDuration() string {
	hours := rand.IntN(168) + 1 // 1-168 hours (1 week)
	if hours < 24 {
		return fmt.Sprintf("%d hours", hours)
	}
	days := hours / 24
	remainingHours := hours % 24
	if remainingHours == 0 {
		return fmt.Sprintf("%d days", days)
	}
	return fmt.Sprintf("%d days %d hours", days, remainingHours)
}

// Filename generation
var fileExtensions = []string{".txt", ".pdf", ".doc", ".docx", ".jpg", ".png", ".mp4", ".mp3", ".zip", ".csv"}
var fileWords = []string{"document", "report", "image", "photo", "video", "data", "backup", "file", "archive", "presentation"}

func GenerateFilename() string {
	word := fileWords[rand.IntN(len(fileWords))]
	number := rand.IntN(999) + 1
	ext := fileExtensions[rand.IntN(len(fileExtensions))]
	return fmt.Sprintf("%s_%03d%s", word, number, ext)
}

// Hashtag generation
var hashtagWords = []string{"tech", "design", "coding", "startup", "innovation", "digital", "future", "ai", "data", "cloud"}

func GenerateHashtag() string {
	word1 := hashtagWords[rand.IntN(len(hashtagWords))]
	word2 := hashtagWords[rand.IntN(len(hashtagWords))]
	return fmt.Sprintf("#%s%s", word1, word2)
}

// Longitude generation (-180 to 180)
func GenerateLongitude() float64 {
	return (rand.Float64() * 360) - 180
}

// Latitude generation (-90 to 90)
func GenerateLatitude() float64 {
	return (rand.Float64() * 180) - 90
}

// Temperature generation (in Celsius, -50 to 50)
func GenerateTemperature() float64 {
	return (rand.Float64() * 100) - 50
}

// Weight generation (in kg, 0.1 to 200)
func GenerateWeight() float64 {
	return rand.Float64()*199.9 + 0.1
}

// Height generation (in cm, 50 to 250)
func GenerateHeight() float64 {
	return rand.Float64()*200 + 50
}