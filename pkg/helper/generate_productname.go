package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateProductName() string {
	adjectives := []string{"Awesome", "Fantastic", "Superb", "Incredible", "Amazing"}
	nouns := []string{"Widget", "Gadget", "Device", "Tool", "Apparatus"}

	rand.Seed(time.Now().UnixNano())

	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]

	productName := fmt.Sprintf("%s %s", adjective, noun)

	return productName
}
