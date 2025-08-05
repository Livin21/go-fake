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