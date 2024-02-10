package helper

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	firstNames = []string{
		"John", "Emma", "Michael", "Olivia", "William", "Ava", "James", "Sophia", "Benjamin", "Isabella",
		"Mason", "Charlotte", "Jacob", "Amelia", "Liam", "Harper", "Ethan", "Evelyn", "Noah", "Abigail",
		"Alexander", "Emily", "Sebastian", "Elizabeth", "Daniel", "Mia", "Matthew", "Sofia", "Joseph", "Camila",
		"Henry", "Avery", "David", "Scarlett", "Samuel", "Victoria", "Gabriel", "Madison", "Jackson", "Luna",
		"Carter", "Grace", "Owen", "Chloe", "Wyatt", "Penelope", "Jayden", "Layla", "Jack", "Riley",
		"Luke", "Zoey", "Leo", "Nora", "Julian", "Lily", "Oliver", "Ellie", "Christopher", "Hannah",
		"Joshua", "Aria", "Andrew", "Aubrey", "Nathan", "Zoe", "Grayson", "Stella", "Levi", "Natalie",
		"Samuel", "Zara", "Isaac", "Addison", "Lincoln", "Mila", "Anthony", "Eleanor", "Hudson", "Paisley",
		"Elijah", "Brooklyn", "Caleb", "Hazel", "Isaiah", "Audrey", "Hunter", "Aurora", "Christian", "Genesis",
		"Eli", "Bella", "Jonathan", "Savannah", "Landon", "Skylar", "Nicholas", "Victoria", "Aaron", "Claire",
	}

	lastNames = []string{
		"Smith", "Johnson", "Brown", "Taylor", "Miller", "Wilson", "Moore", "Anderson", "Thomas", "Jackson",
		"White", "Harris", "Martin", "Thompson", "Garcia", "Martinez", "Robinson", "Clark", "Rodriguez", "Lewis",
		"Lee", "Walker", "Hall", "Allen", "Young", "Hernandez", "King", "Wright", "Lopez", "Hill",
		"Scott", "Green", "Adams", "Baker", "Gonzalez", "Nelson", "Carter", "Mitchell", "Perez", "Roberts",
		"Turner", "Phillips", "Campbell", "Parker", "Evans", "Edwards", "Collins", "Stewart", "Sanchez", "Morris",
		"Rogers", "Reed", "Cook", "Morgan", "Bell", "Murphy", "Bailey", "Cooper", "Richardson", "Cox",
		"Howard", "Ward", "Torres", "Peterson", "Gray", "Ramirez", "James", "Watson", "Brooks", "Kelly",
		"Sanders", "Price", "Bennett", "Wood", "Barnes", "Ross", "Henderson", "Coleman", "Jenkins", "Perry",
		"Powell", "Long", "Patterson", "Hughes", "Flores", "Washington", "Butler", "Simmons", "Foster", "Gonzales",
	}
)

func GenerateFullName() string {
	rand.Seed(time.Now().UnixNano())

	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	return fmt.Sprintf("%s %s", firstName, lastName)
}
