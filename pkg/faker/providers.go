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